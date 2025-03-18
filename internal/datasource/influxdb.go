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
	
	// Get query API for the org
	ds.queryAPI = ds.client.QueryAPI(org)
	
	return nil
}

// FetchData retrieves data from InfluxDB based on the provided Flux query
func (ds *InfluxDBDataSource) FetchData(ctx context.Context, query interface{}) (interface{}, error) {
	// The query should be provided as a string containing a Flux query
	fluxQuery, ok := query.(string)
	if !ok {
		return nil, fmt.Errorf("query must be a Flux query string")
	}
	
	// Run query
	result, err := ds.queryAPI.Query(ctx, fluxQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying InfluxDB: %w", err)
	}
	
	// Process the results
	var rows []map[string]interface{}
	for result.Next() {
		// Convert the record to a map
		record := make(map[string]interface{})
		
		// Get all the values
		for k, v := range result.Record().Values() {
			record[k] = v
		}
		
		// Add special handling for time
		t := result.Record().Time()
		if !t.IsZero() {
			record["time"] = t.Format(time.RFC3339)
		}
		
		rows = append(rows, record)
	}
	
	// Check for errors after iterating through the result
	if result.Err() != nil {
		return nil, fmt.Errorf("error processing query results: %w", result.Err())
	}
	
	return rows, nil
}

// Fetch retrieves default data from InfluxDB 
func (ds *InfluxDBDataSource) Fetch() (interface{}, error) {
	// Provide a default query that returns some basic data
	defaultQuery := `from(bucket:"_monitoring") 
                   |> range(start: -1h) 
                   |> filter(fn: (r) => r._measurement == "system" and r._field == "uptime") 
                   |> limit(n: 10)`
	
	return ds.FetchData(context.Background(), defaultQuery)
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
