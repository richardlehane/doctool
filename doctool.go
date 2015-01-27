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
	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() { // iterate through entries of OLE document
		if entry.Name == "WordDocument" {
			fib := make([]byte, 314)
			i, _ := entry.Read(fib)
			if i < 314 {
				return nil
			}
			// we only care about first 314 bytes of the WordDocument stream.
			// The File Information Block (FIB) starts at 0 offset. The FibRgFcLcb97 section of the FIB starts 154 bytes in.
			// All the items in the FibRgFcLcb97 are listed in the fib_bits.txt doc in this repo. They are each 4 bytes long.
			// You can calculate the relevant offsets by looking at the place of these items in the fib_bits.txt list.
			mn := binary.LittleEndian.Uint32(fib[286:290]) // Interpret the bytes as an unsigned 32-bit integer in little endian order
			hdr := binary.LittleEndian.Uint32(fib[294:298])
			ftn := binary.LittleEndian.Uint32(fib[302:306])
			com := binary.LittleEndian.Uint32(fib[310:314])
			fmt.Printf("Field size of main doc: %d\n"+
				"Field size of header/footer section: %d\n"+
				"Field size of footnote section: %d\n"+
				"Field size of comment section: %d\n",
				mn, hdr, ftn, com)
			return nil // found what we are looking for, can finish here
		}
	}
	return nil // didn't find a WordDocument stream
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
