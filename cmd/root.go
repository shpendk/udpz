package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"udpz/pkg/scan"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/spf13/cobra"
)

var (
	// Scan options
	concurrency     uint = 60
	timeoutMs       uint = 2000
	retransmissions uint = 2

	// Misc options
	noLog bool

	// Output options
	outputPath   string
	outputFormat string = "text"
	outputAppend bool   = true

	// Constraints
	supportedFormats = map[string]bool{
		"text":   true,
		"txt":    true,
		"csv":    true,
		"tsv":    true,
		"pretty": true,
		"json":   true,
	}
)

func init() {
	rootCmd.InitDefaultVersionFlag()
	rootCmd.PersistentFlags().BoolVarP(&outputAppend, "append", "a", outputAppend, "Append results to output file")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "f", outputFormat, "Output format [text, csv, tsv]")
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", outputPath, "Save results to file")
	rootCmd.PersistentFlags().UintVarP(&concurrency, "threads", "c", concurrency, "Number of goroutines to use for scanning")
	rootCmd.PersistentFlags().UintVarP(&retransmissions, "retries", "r", retransmissions, "Number of probe retransmissions")
	rootCmd.PersistentFlags().UintVarP(&timeoutMs, "timeout", "t", timeoutMs, "UDP Probe timeout in milliseconds")
	rootCmd.PersistentFlags().BoolVarP(&noLog, "no-log", "n", noLog, "Disable logging and progress output")
}

var rootCmd = &cobra.Command{
	Use:     "udpz [flags] [targets ...]",
	Short:   "Speedy probe-oriented UDP port scanner",
	Version: "0.0.1-beta",
	Long: `
  ┳┳  ┳┓  ┏┓  ┏┓
  ┃┃━━┃┃━━┃┃━━┏┛
  ┗┛  ┻┛  ┣┛  ┗┛

  Author: Bryan McNulty (@bryanmcnulty)
  Source: https://github.com/bryanmcnulty/udpz`,

	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, targets []string) (err error) {

		var outputFile *os.File
		var outputFlags int = os.O_WRONLY | os.O_CREATE
		outputFormat = strings.ToLower(outputFormat)

		if _, ok := supportedFormats[outputFormat]; !ok {
			return fmt.Errorf("invalid output format: %s", outputFormat)
		}
		if concurrency < 1 {
			return errors.New("concurrency value must be > 0")
		}
		if timeoutMs < 1 {
			return errors.New("timeout value must be > 0")
		}
		if outputAppend {
			outputFlags |= os.O_APPEND
		}

		pw := progress.NewWriter()
		pw.SetTrackerPosition(progress.PositionRight)
		pw.SetMessageLength(15)
		pw.SetOutputWriter(os.Stderr)
		pw.SetNumTrackersExpected(1)
		pw.SetUpdateFrequency(10 * time.Millisecond)
		pw.SetAutoStop(true)
		go pw.Render()

		scanner := scan.NewUdpProbeScanner(concurrency, retransmissions, time.Duration(timeoutMs)*time.Millisecond, pw)

		scanner.Scan(targets)
		for pw.IsRenderInProgress() {
			time.Sleep(5 * time.Millisecond)
		}
		if scanner.Length() > 0 {
			if outputPath == "" {
				outputFile = os.Stdout
			} else if outputFile, err = os.OpenFile(outputPath, outputFlags, 0o644); err == nil {
				defer outputFile.Close()
			} else {
				fmt.Fprintf(os.Stderr, "Could not open output file for writing: %s", outputPath)
				outputFile = os.Stdout
			}
			if outputFormat == "json" {
				scanner.SaveJson(outputFile)
			} else if outputFormat == "yml" {
				// TODO
			} else {
				scanner.SaveTable(outputFormat, outputFile)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Could not identify any UDP ports/services")
		}

		return
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
