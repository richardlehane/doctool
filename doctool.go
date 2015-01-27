// Copyright 2015 Richard Lehane. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Check file information block in word docs for presence for fields (gives raw byte size of field information)
// Examples:
//    ./doctool test.doc
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/richardlehane/mscfb"
)

const (
	UNSET int = iota
	TAB0
	TAB1
)

// start of a func, just a stub
func processField(b []byte) string {
	numDataElements := (len(b) - 4) / 6
	return fmt.Sprintf("%d", numDataElements)
}

func process(in string) error {
	file, err := os.Open(in)
	if err != nil {
		return err
	}
	defer file.Close()
	doc, err := mscfb.New(file)
	if err != nil {
		return err // not an OLE file?
	}
	var table, table1, table0, wordDoc *mscfb.File
	whichTable := UNSET
	var fib []byte
	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() { // iterate through entries of OLE document
		switch entry.Name {
		default:
			continue
		case "0Table":
			table0 = entry
			if whichTable == TAB0 {
				break
			}
		case "1Table":
			table1 = entry
			if whichTable == TAB1 {
				break
			}
		case "WordDocument":
			wordDoc = entry
			fib = make([]byte, 634)
			i, _ := wordDoc.Read(fib)
			if i < 634 {
				return nil // fib is not long enough
			}
			byt := fib[11]
			whichTable = int(byt>>1&1) + 1 // set which table (0Table or 1Table) is the table stream. Do this because a doc can have both but only one will be referenced. It marked by a single bit within the llth byte of the header.
			if (whichTable == TAB0 && table0 != nil) || (whichTable == TAB1 && table1 != nil) {
				break
			}
		}
	}
	// set the table to either 0Table or 1Table stream
	switch whichTable {
	case UNSET:
		return nil
	case TAB0:
		if table0 == nil {
			return nil
		}
		table = table0
	case TAB1:
		if table1 == nil {
			return nil
		}
		table = table1
	}
	// Get offsets (in table stream) and sizes of field data from the FibRgFcLcb97 section of the FIB (which starts 154 bytes in).
	// All the items in the FibRgFcLcb97 are listed in the fib_bits.txt doc in this repo. They are each 4 bytes long.
	// You can calculate the relevant offsets by looking at the place of these items in the fib_bits.txt list.
	mo, mn := binary.LittleEndian.Uint32(fib[282:286]), binary.LittleEndian.Uint32(fib[286:290]) // Interpret the bytes as an unsigned 32-bit integer in little endian order
	hdro, hdr := binary.LittleEndian.Uint32(fib[290:294]), binary.LittleEndian.Uint32(fib[294:298])
	ftno, ftn := binary.LittleEndian.Uint32(fib[298:302]), binary.LittleEndian.Uint32(fib[302:306])
	como, com := binary.LittleEndian.Uint32(fib[306:310]), binary.LittleEndian.Uint32(fib[310:314])
	endo, end := binary.LittleEndian.Uint32(fib[538:542]), binary.LittleEndian.Uint32(fib[542:546])
	txbo, txb := binary.LittleEndian.Uint32(fib[618:622]), binary.LittleEndian.Uint32(fib[622:626])
	htxo, htx := binary.LittleEndian.Uint32(fib[626:630]), binary.LittleEndian.Uint32(fib[630:634])
	if mn+hdr+ftn+com+end+txb+htx == 0 {
		return nil // no fields
	}
	tableBuf := make([]byte, int(table.Size)) // read all the Table stream into a byte buffer
	table.Read(tableBuf)
	// now for each offset and length pair, print the relevant bytes from the table stream (after checking that don't overflow bounds of that slice)
	if mn > 0 {
		if int(mo+mn) > len(tableBuf) {
			return nil
		}
		fmt.Printf("Main field data: %v\n", tableBuf[int(mo):int(mo+mn)])
	}
	if hdr > 0 {
		if int(hdro+hdr) > len(tableBuf) {
			return nil
		}
		fmt.Printf("Header/footer field data: %v\n", tableBuf[int(hdro):int(hdro+hdr)])
	}
	if ftn > 0 {
		if int(ftno+ftn) > len(tableBuf) {
			return nil
		}
		fmt.Printf("Footnote field data: %v\n", tableBuf[int(ftno):int(ftno+ftn)])
	}
	if com > 0 {
		if int(como+com) > len(tableBuf) {
			return nil
		}
		fmt.Printf("Comment field data: %v\n", tableBuf[int(como):int(como+com)])
	}
	if end > 0 {
		if int(endo+end) > len(tableBuf) {
			return nil
		}
		fmt.Printf("Endnote field data: %v\n", tableBuf[int(endo):int(endo+end)])
	}
	if txb > 0 {
		if int(txbo+txb) > len(tableBuf) {
			return nil
		}
		fmt.Printf("Textbox field data: %v\n", tableBuf[int(txbo):int(txbo+txb)])
	}
	if htx > 0 {
		if int(htxo+htx) > len(tableBuf) {
			return nil
		}
		fmt.Printf("Header/footer textbox field data: %v\n", tableBuf[int(htxo):int(htxo+htx)])
	}
	return nil
}

func main() {
	flag.Parse() // we don't actually use any flags in this script, I just copied this from another file (comdump)! Left in in case add flags in future
	ins := flag.Args()
	if len(ins) < 1 {
		log.Fatalln("Missing required argument: path to a word document")
	}
	for _, in := range ins { // you can process a bunch of files at once by using: ./doctool doc1.doc doc2.doc doc3.doc etc.
		err := process(in)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
