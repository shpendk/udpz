package scan

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v3"
)

func (sc *UdpProbeScanner) SaveJson(output *os.File) error {
	if data, err := json.Marshal(&sc.results); err == nil {
		output.Write(data)
		return nil
	} else {
		return err
	}
}

func (sc *UdpProbeScanner) SaveYAML(output *os.File) error {
	if data, err := yaml.Marshal(&sc.results); err == nil {
		output.Write(data)
		return nil
	} else {
		return err
	}
}

func (sc *UdpProbeScanner) SaveTable(format string, output *os.File) {

	resultsTable := table.NewWriter()
	resultsTable.AppendHeader(table.Row{"Host", "Port", "State", "Service", "Probes"})

	for host, ports := range sc.resultsMap {
		for port, results := range ports {
			resultMap := make(map[string][]PortResult)

			for _, result := range results {
				if _, ok := resultMap[result.Service.NameShort]; !ok {
					resultMap[result.Service.NameShort] = []PortResult{}
				}
				resultMap[result.Service.NameShort] = append(resultMap[result.Service.NameShort], result)
			}
			for service, results := range resultMap {
				probeNamesMap := make(map[string]bool)
				probeNames := []string{}

				for _, result := range results {
					if stat, ok := probeNamesMap[result.Probe.Name]; !(ok && stat) {
						probeNames = append(probeNames, result.Probe.Name)
						probeNamesMap[result.Probe.Name] = true
					}
				}
				resultsTable.AppendRow(table.Row{
					host,
					fmt.Sprintf("%d/UDP", port),
					"OPEN",
					service,
					strings.Join(probeNames, ",\n"),
				})
			}
		}
		resultsTable.AppendSeparator()
	}
	resultsTable.SetOutputMirror(output)
	if format == "text" || format == "txt" {
		resultsTable.Render()
	} else if format == "tsv" {
		resultsTable.RenderTSV()
	} else if format == "csv" {
		resultsTable.RenderCSV()
	} else if format == "pretty" {
		resultsTable.SetStyle(table.StyleRounded)
		resultsTable.Render()
	}
}
