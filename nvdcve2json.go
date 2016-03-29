package main

import (
	"github.com/docopt/docopt-go"

	"encoding/json"
	"encoding/xml"
	"os"
)

const version = "nvdcve2json 0.2.0"
const usage = `nvdcve2json.

Usage:
  nvdcve2json [--filter "cpe:/o:apple:iphone_os" --filter "cpe:/o:google:android"]
              [--input FILE | -i FILE]
  nvdcve2json -h | --help
  nvdcve2json --version

Options:
  --filter <string>	Filters by the given string on the cpe fields (can take multiple).
  -i --input FILE	Input XML file for the nvd's CVEs, defaults to stdin.
  -h --help     	Show this screen.
  --version     	Show version.
`

type VulnerableConfiguration struct {
	Id string `xml:"id"`
}

type VulnerableSoftwareList struct {
	Product []string `xml:"product"`
}

type CVSS struct {
	Score               float32 `xml:"score"`
	AccessVector        string  `xml:"access-vector"`
	AccessComplexity    string  `xml:"access-complexity"`
	Authentication      string  `xml:"authentication"`
	ConfidentialyImpact string  `xml:"confidentiality-impact"`
	IntegrityImpact     string  `xml:"integrity-impact"`
	AvailabilityImpact  string  `xml:"availability-impact"`
	Source              string  `xml:"source"`
	Generated           string  `xml:"generated-on-datetime"`
}

type Entry struct {
	Id                      string                    `xml:"cve-id"`
	Published               string                    `xml:"published-datetime"`
	Modified                string                    `xml:"last-modified-datetime"`
	VulnerableConfiguration []VulnerableConfiguration `xml:"vulnerable-configuration"`
	CVSS                    CVSS                      `xml:"cvss>base_metrics"`
	Summary                 string                    `xml:"summary"`
	VulnerableSoftwareList  VulnerableSoftwareList    `xml:"vulnerable-software-list"`
}

func writeDecoded(args map[string]interface{}, decoded Entry) {
	entry, _ := json.Marshal(decoded)
	if args["--output"] == nil {
		os.Stdout.Write(entry)
	}
}

func decodeXML(args map[string]interface{}, input *os.File) {
	decoder := xml.NewDecoder(input)
	var inElement string

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			inElement = se.Name.Local
			if inElement == "entry" {
				var entry Entry
				decoder.DecodeElement(&entry, &se)
				writeDecoded(args, entry)
			}
		}
	}
}

func main() {
	args, _ := docopt.Parse(usage, nil, true, version, false)

	if args["--input"] != nil {
		inputArg, _ := args["--input"].(string)

		input, err := os.Open(inputArg)
		if err != nil {
			panic(err)
		}
		defer input.Close()
		decodeXML(args, input)
	} else {
		input := os.Stdin
		decodeXML(args, input)
	}
}
