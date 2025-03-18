package datasource

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioDataSource is a data source that connects to a MinIO instance
type MinioDataSource struct {
	client     *minio.Client
	bucketName string
}

// MinioQuery represents a query for the MinIO data source
type MinioQuery struct {
	ObjectName string `json:"objectName"`
	Format     string `json:"format,omitempty"` // json, csv, text, binary
}

// Initialize sets up the MinIO data source with the provided configuration
func (ds *MinioDataSource) Initialize(config map[string]interface{}) error {
	// Extract connection parameters from config
	endpoint, ok := config["endpoint"].(string)
	if !ok {
		return fmt.Errorf("endpoint is required")
	}

	accessKey, ok := config["accessKey"].(string)
	if !ok {
		return fmt.Errorf("accessKey is required")
	}

	secretKey, ok := config["secretKey"].(string)
	if !ok {
		return fmt.Errorf("secretKey is required")
	}

	bucketName, ok := config["bucket"].(string)
	if !ok {
		return fmt.Errorf("bucket is required")
	}

	ds.bucketName = bucketName

	// Check for SSL configuration
	useSSL := true
	if ssl, ok := config["useSSL"].(bool); ok {
		useSSL = ssl
	}

	// Create a new MinIO client
	var err error
	ds.client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Verify bucket exists
	exists, err := ds.client.BucketExists(context.Background(), bucketName)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("bucket %s does not exist", bucketName)
	}

	return nil
}

// FetchData retrieves data from a MinIO object
func (ds *MinioDataSource) FetchData(ctx context.Context, query interface{}) (interface{}, error) {
	var minioQuery MinioQuery

	// Handle different query formats
	switch q := query.(type) {
	case string:
		// If query is a string, assume it's the object name
		minioQuery.ObjectName = q
		// Determine format from file extension
		if strings.HasSuffix(q, ".json") {
			minioQuery.Format = "json"
		} else if strings.HasSuffix(q, ".csv") {
			minioQuery.Format = "csv"
		} else {
			minioQuery.Format = "text"
		}
	case map[string]interface{}:
		// Convert map to JSON and then to MinioQuery
		jsonBytes, err := json.Marshal(q)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal query: %w", err)
		}
		if err := json.Unmarshal(jsonBytes, &minioQuery); err != nil {
			return nil, fmt.Errorf("failed to unmarshal query: %w", err)
		}
	case MinioQuery:
		minioQuery = q
	default:
		return nil, fmt.Errorf("query must be a string, map, or MinioQuery")
	}

	// Get the object from MinIO
	obj, err := ds.client.GetObject(ctx, ds.bucketName, minioQuery.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}
	defer obj.Close()

	// Read the object data
	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to read object: %w", err)
	}

	// Process data based on format
	switch minioQuery.Format {
	case "json":
		var result interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
		return result, nil
	case "csv":
		// Return raw data for CSV (can be further processed by the report generator)
		return string(data), nil
	case "binary":
		return data, nil
	default:
		// Default to text format
		return string(data), nil
	}
}

// Close cleans up resources used by the MinIO data source
func (ds *MinioDataSource) Close() error {
	// No specific cleanup needed for MinIO client
	return nil
}

// NewMinioDataSource creates a new MinIO data source
func NewMinioDataSource() *MinioDataSource {
	return &MinioDataSource{}
}
