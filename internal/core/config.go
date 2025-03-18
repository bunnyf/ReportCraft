package core

import (
	"encoding/json"
	"os"
)

// 确保此处的 Config 结构体与 interfaces.go 中的定义一致
type Config struct {
	OutputPath   string                 `json:"outputPath"`
	Template     string                 `json:"template"`
	ReportType   string                 `json:"reportType"`
	OutputFormat string                 `json:"outputFormat"`
	Parameters   map[string]interface{} `json:"parameters"`
	DataSources  []DataSourceConfig     `json:"dataSources"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
