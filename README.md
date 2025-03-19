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

1. Clone the repository:

```bash
git clone https://github.com/genrep/ReportCraft.git
```

1. Navigate to the project directory and build the project:

```bash
# macOS/Linux
cd ReportCraft
go build -o reportcraft ./cmd/reportcraft/

# Windows
cd ReportCraft
go build -o reportcraft.exe ./cmd/reportcraft/
```

1. Run the program:

```bash
# macOS/Linux
./reportcraft -config=path/to/config.json

# Windows
reportcraft.exe -config=path\to\config.json
```

1. The generated report will be saved at the output path specified in the config file

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

## Vibration Analysis Features

ReportCraft supports generating device vibration waveform and spectrum analysis reports with the following features:

1. Device Information Display: Device name, ID, measurement location, etc.
2. Waveform Data Charts: Display time-domain vibration waveforms
3. Spectrum Data Charts: Display frequency-domain analysis results
4. Feature Parameter Tables: Show characteristic parameters of waveform and spectrum

### Chart Visualization Types

ReportCraft supports multiple visualization types for different analysis needs:

#### Basic Chart Types

- Line Chart: For time domain signals and trends
- Bar Chart: For frequency spectrum display
- Scatter Plot: For correlation analysis

#### Advanced Visualization Types

- **Heat Map**: For displaying intensity variations across two dimensions
- **3D Surface Plot**: For three-dimensional data visualization
- **Waterfall Plot**: For analyzing spectrum changes over time or operational conditions

```json
"chartType": "heatmap",
"chartType": "3dsurface",
"chartType": "waterfall"
```

### Chart Style Configuration

Vibration analysis reports offer rich chart style configuration options:

#### Basic Style Configuration

```json
"chartStyle": {
  "lineColor": "#1E90FF",
  "lineWidth": 1.5,
  "gridLines": true
}
```

#### Enhanced Style Configuration

```json
"chartStyle": {
  // Basic style
  "lineColor": "#1E90FF",
  "lineWidth": 1.5,
  "gridLines": true,
  
  // Coordinate axis range
  "axis": {
    "xMin": 0,
    "xMax": 100,
    "yMin": -0.5,
    "yMax": 0.5,
    "zMin": 0,  // For 3D charts
    "zMax": 1.0 // For 3D charts
  },
  
  // Data point markers
  "markers": {
    "show": true,
    "size": 4,
    "color": "#FF4500",
    "shape": "circle"  // Options: circle, square, triangle, diamond
  },
  
  // Grid lines configuration
  "grid": {
    "show": true,
    "color": "#CCCCCC",
    "lineStyle": "solid", // Options: solid, dashed, dotted
    "lineWidth": 0.5,
    "minorGrid": {
      "show": true,
      "color": "#EEEEEE",
      "lineStyle": "dotted",
      "lineWidth": 0.25
    }
  },
  
  // Highlight regions
  "highlight": {
    "regions": [
      {
        "xStart": 40,
        "xEnd": 60,
        "color": "rgba(255, 100, 100, 0.2)",
        "label": "Feature Region"
      }
    ]
  },
  
  // Title and label styles
  "labels": {
    "title": {
      "fontSize": 14,
      "fontWeight": "bold",
      "color": "#333333"
    },
    "axis": {
      "fontSize": 12,
      "color": "#666666"
    }
  },
  
  // Legend configuration
  "legend": {
    "position": "bottom", // Options: top, bottom, left, right
    "fontSize": 11,
    "color": "#333333"
  },
  
  // Specific for Heat Map
  "heatmap": {
    "colorScale": "viridis", // Options: viridis, jet, plasma, inferno
    "showColorBar": true,
    "interpolate": true
  },
  
  // Specific for 3D Surface
  "3dsurface": {
    "wireframe": true,
    "colorScale": "jet",
    "rotation": {
      "x": 45,
      "y": 30,
      "z": 0
    }
  },
  
  // Specific for Waterfall Plot
  "waterfall": {
    "baseColor": "#1E90FF",
    "colorGradient": true,
    "spacing": 0.1,
    "perspective": 30
  }
}
```

### Example Configurations

`examples` directory contains example configuration files:

- `spectrum-report-embedded.json`: Shows how to create a vibration analysis report using embedded data
- `spectrum-report-csv.json`: Shows how to create a vibration analysis report using data from an external CSV file

### Configuration Options

Vibration analysis reports support the following configuration options:

```json
{
  "deviceInfo": {
    "deviceName": "Device Name",
    "deviceId": "Device ID",
    "location": "Location Information",
    "measurementPoint": "Measurement Point Information"
  },
  "waveformConfig": {
    "title": "Waveform Chart Title",
    "xLabel": "X-axis Label",
    "yLabel": "Y-axis Label"
  },
  "spectrumConfig": {
    "title": "Spectrum Chart Title",
    "xLabel": "X-axis Label",
    "yLabel": "Y-axis Label"
  }
}
```

## Windows环境注意事项

在Windows环境中使用ReportCraft时，需要注意以下几点：

1. 文件路径分隔符：Windows使用反斜杠(`\`)作为路径分隔符，而配置文件中建议使用正斜杠(`/`)，两者都可以正常工作。

2. 输出文件命名：建议使用英文文件名，避免使用中文或特殊字符作为文件名，以防止编码问题。

3. 命令行运行：在Windows环境中，使用`reportcraft.exe`而不是`./reportcraft`来运行程序。

4. 配置文件示例：所有示例配置文件中的路径格式都兼容Windows环境。

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
