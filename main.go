/*
  Copyright (C) 2021 Sinuhé Téllez Rivera

  sparkplugclidecoder is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  sparkplugclidecoder is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with sparkplugclidecoder.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dubyte/sparkplugclidecoder/internal/sparkplug"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	sep           = kingpin.Flag("sep", "adds a separator between messages: <newline>sep<newline>").Default("").Short('s').String()
	format        = kingpin.Flag("format", "Output format: prototext, protojson").Default("prototext").Short('f').String()
	maxBufferSize = kingpin.Flag("maxBufferSize", "size of the buffer").Default(strconv.Itoa(bufio.MaxScanTokenSize)).Short('t').Int()
)

const startBufSize = 4096

//go:generate protoc --go_out=. --proto_path=./internal/sparkplug sparkplug.proto
func main() {
	kingpin.Parse()
	newBuf := make([]byte, startBufSize)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(newBuf, *maxBufferSize)

	defer func() {
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()

	for scanner.Scan() {
		var payload sparkplug.Payload
		hexstr := scanner.Bytes()

		if len(hexstr) == 0 {
			return
		}

		payloadBin := make([]byte, hex.DecodedLen(len(hexstr)))
		n, err := hex.Decode(payloadBin, hexstr)
		if err != nil {
			log.Fatalf("err while decoding hex: %s", err)
		}

		err = proto.Unmarshal(payloadBin[:n], &payload)
		if err != nil {
			log.Fatalf("err while unmarshalling: %s", err)
		}
		var out string
		switch *format {
		case "protojson":
			out = protojson.Format(&payload)
		case "prototext":

			out = prototext.Format(&payload)
		default:
			log.Fatalf("unknown format: %s", err)
		}

		fmt.Printf("%s", out)

		if *sep != "" {
			fmt.Printf("\n%s\n", *sep)
		}
	}

}
