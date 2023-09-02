package main

import (
	"flag"
	"fmt"
	"os"
)

func Usage() {
	fmt.Printf(`nvdcve2json -- parser of XML to JSON for NVD's CVE source files

Usage of %s:
  nvdcve2json [--filter "cpe:/o:apple:iphone_os" --filter "cpe:/o:google:android"...]
              [--input FILE | -i FILE]
  nvdcve2json -h | --help
  nvdcve2json --version

Options:
  --filter <string>	Filters by the given string on the CPE fields.
  -i --input FILE	Input XML file for the NVD's CVEs, defaults to stdin.
  -h --help     	Show this screen.
  --version     	Show version.

Examples:
  gunzip file.xml.gz | nvdcve2json \
    --filter "cpe:/o:microsoft:windows" \
	--filter "cpe:/o:apple:mac_os_x" > output.json

  nvdcve2json < file.xml > output.json

Released under the MIT license. Source code available at
https://github.com/ghostbar/nvdcve2json.git.

Copyright (c) 2016-2023 Jose-Luis Rivas. All Rights Reserved.
`, os.Args[0])
	Version()
}

// The version value must be changed with -ldflags during build, like this:
//
//	go build -ldflags="-X 'main.version=$(git describe --tags --always)'"
var version = "development"

func Version() {
	fmt.Printf(`Version: %s.`, version)
}

func main() {
	var (
		filters   StringFlags
		inputFlag string
	)

	printVersion := flag.Bool("version", false, "Show version")
	flag.Var(&filters, "filter", "Filters by the given string on the cpe fields (can take multiple)")
	flag.StringVar(&inputFlag, "input", "", "Input XML file for the NVD's CVEs. Uses stdin if not set")
	flag.StringVar(&inputFlag, "i", "", "Input XML file for the NVD's CVEs. Uses stdin if not set (shorthand)")

	flag.Usage = Usage
	flag.Parse()

	in := os.Stdin

	switch {
	case *printVersion:
		Version()
	case inputFlag != "":
		var err error
		in, err = os.Open(inputFlag)
		if err != nil {
			panic(err)
		}
		defer in.Close()
		fallthrough
	default:
		decodeXML(filters.Slice(), in)
	}
}
