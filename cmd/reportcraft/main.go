package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/genrep/internal/core"
	"github.com/genrep/internal/report"
)

func main() {
	configPath := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	if *configPath == "" {
		fmt.Println("Error: config flag is required")
		os.Exit(1)
	}

	cfg, err := core.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Register report generators based on the report type
	var reportGenerator core.ReportGenerator
	switch cfg.ReportType {
	case "hello-word":
		reportGenerator = report.NewHelloWordReport()
	case "hello-excel":
		reportGenerator = report.NewHelloExcelReport()
	case "waveform-report":
		reportGenerator = report.NewWaveformReport()
	case "spectrum-report":
		reportGenerator = report.NewSpectrumReport()
	default:
		fmt.Printf("Error: Unsupported report type: %s\n", cfg.ReportType)
		os.Exit(1)
	}

	err = reportGenerator.Initialize(cfg)
	if err != nil {
		fmt.Printf("Error initializing report generator: %v\n", err)
		os.Exit(1)
	}

	err = reportGenerator.Generate(context.Background(), make(map[string]core.DataSource))
	if err != nil {
		fmt.Printf("Error generating report: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Report generated successfully")
}
