# GenRep - General Reporting Tool

GenRep is a powerful, extensible report generation tool designed to create various types of reports from multiple data sources.

## Overview

GenRep is a command-line tool written in Go that generates reports based on a JSON configuration file. It is designed to be:

1. **Extensible**: New report types can be easily added through a plugin architecture
2. **Versatile**: Supports multiple data sources and report output formats
3. **Programmatic**: Can be called by other programs as a utility
4. **Self-contained**: All configuration is contained in a single JSON file

## Usage

1. First, make sure you have Go 1.20 or higher installed

2. Clone the repository:

```bash
git clone https://github.com/genrep/ReportCraft.git
```

3. Navigate to the project directory and build the project:

```bash
cd ReportCraft
go build -o reportcraft ./cmd/reportcraft/
```

4. Run the program:

```bash
./reportcraft -config=path/to/config.json
```

5. The generated report will be saved at the output path specified in the config file

## Configuration Format

The configuration file is a JSON file with the following structure:

```json
{
  "reportType": "string",           // Type of report to generate
  "outputPath": "string",           // Where to save the generated report
  "outputFormat": "string",         // Output format (docx, xlsx, etc.)
  "parameters": {                   // Report-specific parameters
    // Varies based on report type
  },
  "dataSources": [                  // Array of data sources
    {
      "id": "string",               // Unique identifier for this data source
      "type": "string",             // Type of data source (db, api, file, etc.)
      "config": {                   // Data source-specific configuration
        // Varies based on data source type
      }
    }
  ]
}
```

## Table Report Feature

ReportCraft now supports generating data tables from different sources:

1. External JSON file: Data can be retrieved from a JSON file with a specific data path
2. External CSV file: Data can be imported directly from a CSV file
3. Embedded data: Data can be included directly in the configuration file

### Example Configurations

Example configuration files are available in the `examples` directory:

- `table-report-json.json`: Shows how to use data from an external JSON file
- `table-report-csv.json`: Shows how to use data from an external CSV file
- `table-report-embedded.json`: Shows how to embed data directly in the configuration

### Table Configuration Options

The table report allows for the following configuration options:

```json
{
  "tableConfig": {
    "title": "Table Title",
    "columns": [
      {"field": "columnName", "header": "Column Header", "width": 15},
      {"field": "anotherColumn", "header": "Another Header", "width": 20, "format": "date"}
    ],
    "headerStyle": {
      "bold": true,
      "background": "#DDEBF7",
      "color": "#000000"
    },
    "alternateRowStyle": true
  }
}
```

## Architecture

GenRep is designed with a plugin architecture that allows for easy extension with new report types and data sources:

1. **Core Engine**: Handles configuration parsing, plugin management, and orchestration
2. **Report Plugins**: Implement specific report generation logic
3. **Data Source Adapters**: Connect to and retrieve data from various sources
4. **Output Formatters**: Convert report data to the desired output format

## Adding New Report Types

New report types can be added by implementing the `ReportGenerator` interface and registering the implementation with the plugin system.

## Supported Data Sources

- Databases (SQL, NoSQL)
- APIs (REST, GraphQL)
- Files (CSV, JSON, Excel)
- Custom data sources through the plugin system

## Supported Output Formats

- Microsoft Word (DOCX)
- Microsoft Excel (XLSX)
- PDF
- HTML
- Custom formats through the plugin system
