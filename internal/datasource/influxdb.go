package datasource

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// InfluxDBDataSource is a data source that connects to an InfluxDB 2.0 instance
type InfluxDBDataSource struct {
	client   influxdb2.Client
	queryAPI api.QueryAPI
	org      string
}

// Initialize sets up the InfluxDB data source with the provided configuration
func (ds *InfluxDBDataSource) Initialize(config map[string]interface{}) error {
	// Extract connection parameters from config
	url, ok := config["url"].(string)
	if !ok {
		return fmt.Errorf("url is required")
	}

	token, ok := config["token"].(string)
	if !ok {
		return fmt.Errorf("token is required")
	}

	org, ok := config["org"].(string)
	if !ok {
		return fmt.Errorf("org is required")
	}

	ds.org = org

	// Create a new InfluxDB client
	ds.client = influxdb2.NewClient(url, token)

	// Get the query API
	ds.queryAPI = ds.client.QueryAPI(org)

	return nil
}

// FetchData retrieves data from InfluxDB based on the provided Flux query
func (ds *InfluxDBDataSource) FetchData(ctx context.Context, query interface{}) (interface{}, error) {
	// Check if query is a string
	fluxQuery, ok := query.(string)
	if !ok {
		return nil, fmt.Errorf("query must be a Flux query string")
	}

	// Run the query
	result, err := ds.queryAPI.Query(ctx, fluxQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer result.Close()

	// Process the results
	var data []map[string]interface{}
	for result.Next() {
		record := make(map[string]interface{})

		// Get the record values
		for k, v := range result.Record().Values() {
			record[k] = v
		}

		// Add time field if not present
		if _, ok := record["_time"]; !ok {
			recordTime := result.Record().Time()
			if !recordTime.IsZero() {
				record["_time"] = recordTime.Format(time.RFC3339)
			}
		}

		data = append(data, record)
	}

	// Check for error in the result iteration
	if result.Err() != nil {
		return nil, fmt.Errorf("error parsing query result: %w", result.Err())
	}

	return data, nil
}

// Close cleans up resources used by the InfluxDB data source
func (ds *InfluxDBDataSource) Close() error {
	if ds.client != nil {
		ds.client.Close()
	}
	return nil
}

// NewInfluxDBDataSource creates a new InfluxDB data source
func NewInfluxDBDataSource() *InfluxDBDataSource {
	return &InfluxDBDataSource{}
}
