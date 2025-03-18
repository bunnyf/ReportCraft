package datasource

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

// FileDataSource is a data source that reads data from files
type FileDataSource struct {
	path string
	fileType string
}

// Initialize sets up the file data source with the provided configuration
func (ds *FileDataSource) Initialize(config map[string]interface{}) error {
	// Extract the path from the configuration
	pathValue, ok := config["path"]
	if !ok {
		return fmt.Errorf("path is required")
	}
	
	path, ok := pathValue.(string)
	if !ok {
		return fmt.Errorf("path must be a string")
	}
	
	ds.path = path
	
	// Determine the file type based on the extension
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".csv":
		ds.fileType = "csv"
	case ".json":
		ds.fileType = "json"
	case ".xlsx", ".xls":
		ds.fileType = "excel"
	default:
		return fmt.Errorf("unsupported file type: %s", ext)
	}
	
	return nil
}

// FetchData retrieves data from the file based on the provided query
func (ds *FileDataSource) FetchData(ctx context.Context, query interface{}) (interface{}, error) {
	switch ds.fileType {
	case "csv":
		return ds.readCSV()
	case "json":
		return ds.readJSON()
	case "excel":
		return ds.readExcel(query)
	default:
		return nil, fmt.Errorf("unsupported file type: %s", ds.fileType)
	}
}

// Fetch retrieves data from the file (simplified version without query)
func (ds *FileDataSource) Fetch() (interface{}, error) {
	// This is a simplified version that calls FetchData with nil query
	return ds.FetchData(context.Background(), nil)
}

// Close cleans up any resources used by the data source
func (ds *FileDataSource) Close() error {
	// No resources to clean up for file data source
	return nil
}

// readCSV reads data from a CSV file
func (ds *FileDataSource) readCSV() ([][]string, error) {
	file, err := os.Open(ds.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// readJSON reads data from a JSON file
func (ds *FileDataSource) readJSON() (interface{}, error) {
	file, err := os.Open(ds.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var data interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}
	
	return data, nil
}

// readExcel reads data from an Excel file
func (ds *FileDataSource) readExcel(query interface{}) (interface{}, error) {
	f, err := excelize.OpenFile(ds.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	// If query is a string, use it as sheet name
	sheetName := "Sheet1" // Default
	if queryStr, ok := query.(string); ok && queryStr != "" {
		sheetName = queryStr
	}
	
	// Get all the rows in the specified sheet
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}
	
	return rows, nil
}

// Factory function for creating FileDataSource instances
func NewFileDataSource() *FileDataSource {
	return &FileDataSource{}
}
