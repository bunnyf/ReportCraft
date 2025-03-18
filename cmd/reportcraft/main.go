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

	reportGenerator := report.NewHelloWordReport()
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
