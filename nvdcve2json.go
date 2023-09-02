package main

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"regexp"
)

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

func writeDecoded(decoded Entry) {
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

func decodeXML(filters []string, input *os.File) {
	decoder := xml.NewDecoder(input)
	var initial bool = true

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
							writeDecoded(entry)
							break
						}
					}
				} else {
					initial = writeComma(initial)
					writeDecoded(entry)
				}
			}
		}
	}

	os.Stdout.WriteString("]")
}
