package cmd

import (
	"fmt"

	"github.com/genrep/internal/core"
	"github.com/spf13/cobra"
)

var configFile string
var engine *core.ReportEngine

var rootCmd = &cobra.Command{
	Use:   "genrep",
	Short: "GenRep - General Reporting Tool",
	Long: `GenRep is a powerful, extensible report generation tool 
designed to create various types of reports from multiple data sources.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if configFile == "" {
			return fmt.Errorf("config file is required")
		}

		if engine == nil {
			return fmt.Errorf("report engine not initialized")
		}

		return engine.GenerateReport(configFile)
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// SetEngine sets the report engine to be used by the commands
func SetEngine(e *core.ReportEngine) {
	engine = e
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to the configuration file (required)")
	rootCmd.MarkFlagRequired("config")
}
