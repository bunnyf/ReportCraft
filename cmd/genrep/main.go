package main

import (
	"fmt"
	"os"

	"github.com/genrep/internal/cmd"
	"github.com/genrep/internal/core"
	"github.com/genrep/internal/datasource"
	"github.com/genrep/internal/logger"
	"github.com/genrep/internal/report"
)

func main() {
	// Set log level to INFO
	logger.SetDefaultLevel(logger.INFO)

	// Initialize report engine
	engine := core.NewReportEngine()

	// Register data sources
	engine.RegisterDataSource("file", func() core.DataSource {
		return datasource.NewFileDataSource()
	})
	engine.RegisterDataSource("influxdb", func() core.DataSource {
		return datasource.NewInfluxDBDataSource()
	})
	engine.RegisterDataSource("minio", func() core.DataSource {
		return datasource.NewMinioDataSource()
	})

	// Register report generators
	engine.RegisterReportGenerator("hello-word", report.NewHelloWordReport())
	engine.RegisterReportGenerator("hello-excel", report.NewHelloExcelReport())
	engine.RegisterReportGenerator("waveform", report.NewWaveformReport())
	engine.RegisterReportGenerator("template", report.NewTemplateReport())

	// Set up the command line interface
	cmd.SetEngine(engine)

	// Execute the root command
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
