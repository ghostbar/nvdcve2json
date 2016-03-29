package main

import (
	"github.com/docopt/docopt-go"

	"encoding/json"
	"encoding/xml"
	"os"
	"regexp"
)

const version = "nvdcve2json 1.0.1"
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

type FactRef struct {
	Name string `xml:"name,attr"`
}

type LogicalTest struct {
	Operator    string        `xml:"operator,attr"`
	Negate      string        `xml:"negate,attr"`
	FactRef     []FactRef     `xml:"fact-ref"`
	LogicalTest []LogicalTest `xml:"logical-test"`
}

type VulnerableConfiguration struct {
	LogicalTest []LogicalTest `xml:"logical-test"`
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
	os.Stdout.Write(entry)
}

func matchesFilter(filter *regexp.Regexp, logicTests []LogicalTest) (matches bool) {
	for _, logicTest := range logicTests {
		if logicTest.Negate == "false" {
			if logicTest.FactRef != nil {
				for _, product := range logicTest.FactRef {
					if filter.MatchString(product.Name) {
						return true
					}
				}
			} else {
				return matchesFilter(filter, logicTest.LogicalTest)
			}
		}
	}

	return false
}

func filterVulnConfs(filter string, vulnConfs []VulnerableConfiguration) (matches bool) {
	match := regexp.MustCompile(filter)
	for _, vulnConf := range vulnConfs {
		if matchesFilter(match, vulnConf.LogicalTest) {
			return true
		}
	}

	return false // if none matched
}

func writeComma(initial bool) (alwaysFalse bool) {
	if !initial {
		os.Stdout.WriteString(",")
	}
	return false
}

func decodeXML(args map[string]interface{}, input *os.File) {
	decoder := xml.NewDecoder(input)
	var initial bool = true
	filters := args["--filter"].([]string)

	os.Stdout.WriteString("[")

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "entry" {
				var entry Entry
				decoder.DecodeElement(&entry, &se)

				if len(filters) != 0 {
					for _, filter := range filters {
						if filterVulnConfs(filter, entry.VulnerableConfiguration) {
							initial = writeComma(initial)
							writeDecoded(args, entry)
							break
						}
					}
				} else {
					initial = writeComma(initial)
					writeDecoded(args, entry)
				}
			}
		}
	}

	os.Stdout.WriteString("]")
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
