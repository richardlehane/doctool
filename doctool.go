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
	"io"
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
		return err
	}
	for {
		entry, err := doc.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if entry.Name == "WordDocument" {
			fib := make([]byte, 320)
			i, _ := entry.Read(fib)
			if i < 320 {
				return nil
			}
			var mn, hdr, ftn, com uint32
			mn = binary.LittleEndian.Uint32(fib[286:290])
			hdr = binary.LittleEndian.Uint32(fib[294:298])
			ftn = binary.LittleEndian.Uint32(fib[302:306])
			com = binary.LittleEndian.Uint32(fib[310:314])
			fmt.Printf("Field size of main doc: %d\nField size of header section: %d\nField size of footnote section: %d\nField size of comment section: %d\n", mn, hdr, ftn, com)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	ins := flag.Args()
	if len(ins) < 1 {
		log.Fatalln("Missing required argument: path_to_compound_object")
	}
	for _, in := range ins {
		err := process(in)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
