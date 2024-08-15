package scan

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func (sc *UdpProbeScanner) SaveJson(output *os.File) {
	if data, err := json.Marshal(&sc.results); err == nil {
		output.Write(data) // TODO: make sure this adds line break
	}
}

func (sc *UdpProbeScanner) SaveTable(format string, output *os.File) {

	resultsTable := table.NewWriter()
	//resultsTable.AppendHeader(table.Row{"Host", "Port", "Service", "Description"})
	resultsTable.AppendHeader(table.Row{"Host", "Port", "Service"})

	for _, result := range sc.results {
		resultsTable.AppendRow(table.Row{
			result.Target,
			fmt.Sprintf("%s/%d", strings.ToUpper(result.Transport), result.Port),
			result.Service.Name,
			//result.Service.Description,
		})
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
