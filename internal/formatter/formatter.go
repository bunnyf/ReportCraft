package formatter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

// OutputFormatter formats data into a specific output format
type OutputFormatter interface {
	Format(data interface{}, outputPath string) error
}

// DocxFormatter formats data into a Word document
type DocxFormatter struct{}

// Format formats the data into a Word document
func (f *DocxFormatter) Format(data interface{}, outputPath string) error {
	fmt.Println("Generating Word document...")
	
	// 获取数据内容
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data is not a map")
	}
	
	// 创建一个文本内容
	var docContent strings.Builder
	
	// 添加标题
	title := "Report"
	if t, ok := dataMap["title"].(string); ok {
		title = t
	}
	
	docContent.WriteString("# " + title + "\n\n")
	
	// 如果有副标题，添加副标题
	if params, ok := dataMap["parameters"].(map[string]interface{}); ok {
		if subtitle, ok := params["subtitle"].(string); ok {
			docContent.WriteString("## " + subtitle + "\n\n")
		}
	}
	
	// 添加设备信息
	docContent.WriteString("## 设备信息：\n\n")
	if params, ok := dataMap["parameters"].(map[string]interface{}); ok {
		if equipInfo, ok := params["equipmentInfo"].(map[string]interface{}); ok {
			for key, value := range equipInfo {
				docContent.WriteString(fmt.Sprintf("- %s: %v\n", key, value))
			}
		}
	}
	
	// 添加结论
	docContent.WriteString("\n## 分析结论：\n\n")
	if params, ok := dataMap["parameters"].(map[string]interface{}); ok {
		if conclusions, ok := params["conclusions"].([]interface{}); ok {
			for _, conclusion := range conclusions {
				docContent.WriteString(fmt.Sprintf("- %v\n", conclusion))
			}
		}
	}
	
	// 创建输出目录（如果不存在）
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	
	// 保存文档
	if err := os.WriteFile(outputPath, []byte(docContent.String()), 0644); err != nil {
		return fmt.Errorf("failed to save Word document: %w", err)
	}
	
	fmt.Printf("Generated Word document at %s\n", outputPath)
	return nil
}

// ExcelFormatter formats data into an Excel spreadsheet
type ExcelFormatter struct{}

// Format formats the data into an Excel spreadsheet
func (f *ExcelFormatter) Format(data interface{}, outputPath string) error {
	fmt.Println("Generating Excel document...")
	
	// 获取数据内容
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data is not a map")
	}
	
	// 创建新的Excel文件
	excel := excelize.NewFile()
	
	// 获取默认Sheet
	sheetName := "Report Data"
	excel.SetSheetName("Sheet1", sheetName)
	
	// 设置标题样式
	titleStyle, err := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   16,
			Family: "Arial",
			Color:  "#333333",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E6F0FF"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create title style: %w", err)
	}
	
	// 设置表头样式
	headerStyle, err := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   11,
			Family: "Arial",
			Color:  "#FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4472C4"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create header style: %w", err)
	}
	
	// 设置数据单元格样式
	dataStyle, err := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size:   10,
			Family: "Arial",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create data style: %w", err)
	}
	
	// 设置交替行样式用于提高可读性
	alternateRowStyle, err := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size:   10,
			Family: "Arial",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#F2F2F2"},
			Pattern: 1,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
			{Type: "right", Color: "#000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create alternate row style: %w", err)
	}
	
	// 添加标题
	title := "设备健康趋势分析报告"
	if t, ok := dataMap["title"].(string); ok {
		title = t
	}
	excel.SetCellValue(sheetName, "A1", title)
	excel.MergeCell(sheetName, "A1", "H1")
	excel.SetCellStyle(sheetName, "A1", "H1", titleStyle)
	excel.SetRowHeight(sheetName, 1, 30)
	
	// 添加报告日期
	reportDate := fmt.Sprintf("报告日期: %s", time.Now().Format("2006-01-02"))
	excel.SetCellValue(sheetName, "A2", reportDate)
	excel.MergeCell(sheetName, "A2", "H2")
	excel.SetCellStyle(sheetName, "A2", "H2", dataStyle)
	
	// 处理数据源
	currentRow := 4
	
	// 添加调试信息
	excel.SetCellValue(sheetName, "A3", fmt.Sprintf("数据源路径: %s", dataSourcePath))
	excel.MergeCell(sheetName, "A3", "H3")
	excel.SetCellStyle(sheetName, "A3", "H3", dataStyle)
	currentRow++
	
	// 获取数据源配置
	var dataSourcePath string
	var dataSourceType string
	if ds, ok := dataMap["dataSource"].(map[string]interface{}); ok {
		if path, ok := ds["path"].(string); ok {
			// 将相对路径转换为绝对路径，以确保正确读取
			if strings.HasPrefix(path, "./") {
				// 移除开头的 ./
				cleanPath := path[2:]
				// 获取工作目录
				workDir, err := os.Getwd()
				if err == nil {
					dataSourcePath = filepath.Join(workDir, cleanPath)
				} else {
					dataSourcePath = path // 如果获取工作目录失败，使用原始路径
				}
			} else {
				dataSourcePath = path
			}
			
			// 确保使用斜杠（而不是反斜杠）以获得更好的跨平台兼容性
			dataSourcePath = strings.ReplaceAll(dataSourcePath, "\\", "/")
			
			fmt.Printf("使用数据源: %s\n", dataSourcePath)
		}
		
		// 获取数据源类型
		if dsType, ok := ds["type"].(string); ok {
			dataSourceType = dsType
			fmt.Printf("数据源类型: %s\n", dataSourceType)
		}
	}
	
	// 如果数据源路径有效，尝试处理数据
	if dataSourcePath != "" {
		// 检查文件是否存在
		if _, err := os.Stat(dataSourcePath); os.IsNotExist(err) {
			fmt.Printf("警告：数据源文件不存在：%s\n", dataSourcePath)
			
			// 尝试其他可能的路径组合
			altPath := filepath.Join(".", dataSourcePath)
			if _, err := os.Stat(altPath); err == nil {
				dataSourcePath = altPath
				fmt.Printf("找到替代路径: %s\n", dataSourcePath)
			} else {
				return fmt.Errorf("data source file does not exist: %s", dataSourcePath)
			}
		}
		
		// 检查文件扩展名
		fileExt := strings.ToLower(filepath.Ext(dataSourcePath))
		
		// 如果数据源类型已指定，优先使用该类型
		if dataSourceType != "" {
			switch dataSourceType {
			case "csv":
				fileExt = ".csv"
			case "json":
				fileExt = ".json"
			}
		}
		
		switch fileExt {
		case ".csv":
			// 处理 CSV 数据
			file, err := os.Open(dataSourcePath)
			if err != nil {
				fmt.Printf("无法打开CSV文件: %v\n", err)
				return fmt.Errorf("failed to open CSV file: %w", err)
			}
			defer file.Close()
			
			reader := csv.NewReader(file)
			reader.TrimLeadingSpace = true // 去除前导空格
			
			// 读取标题行
			headers, err := reader.Read()
			if err != nil {
				fmt.Printf("无法读取CSV标题: %v\n", err)
				return fmt.Errorf("failed to read CSV headers: %w", err)
			}
			
			// 写入表头
			for i, header := range headers {
				col, _ := excelize.ColumnNumberToName(i + 1)
				cell := col + fmt.Sprintf("%d", currentRow)
				excel.SetCellValue(sheetName, cell, header)
				excel.SetCellStyle(sheetName, cell, cell, headerStyle)
			}
			excel.SetRowHeight(sheetName, currentRow, 20)
			currentRow++
			
			// 读取并写入数据行
			rowNumber := 0
			for {
				record, err := reader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Printf("读取CSV行时出错: %v\n", err)
					continue // 跳过此行，继续处理
				}
				
				// 确定行样式 (交替行样式)
				rowStyle := dataStyle
				if rowNumber%2 == 1 {
					rowStyle = alternateRowStyle
				}
				
				// 写入行数据
				for i, value := range record {
					col, _ := excelize.ColumnNumberToName(i + 1)
					cell := col + fmt.Sprintf("%d", currentRow)
					excel.SetCellValue(sheetName, cell, value)
					excel.SetCellStyle(sheetName, cell, cell, rowStyle)
				}
				currentRow++
				rowNumber++
			}
			
			fmt.Printf("成功处理CSV数据，共%d行\n", rowNumber)
			
		case ".json":
			// 处理 JSON 数据
			fileData, err := os.ReadFile(dataSourcePath)
			if err != nil {
				fmt.Printf("无法读取JSON文件: %v\n", err)
				return fmt.Errorf("failed to read JSON file: %w", err)
			}
			
			var jsonData interface{}
			err = json.Unmarshal(fileData, &jsonData)
			if err != nil {
				fmt.Printf("无法解析JSON数据: %v\n", err)
				return fmt.Errorf("failed to parse JSON data: %w", err)
			}
			
			// 处理 JSON 数据 (假设是一个对象数组)
			if jsonArray, ok := jsonData.([]interface{}); ok && len(jsonArray) > 0 {
				// 从第一个对象提取标题
				if firstObj, ok := jsonArray[0].(map[string]interface{}); ok {
					// 提取并排序标题键
					headers := make([]string, 0, len(firstObj))
					for key := range firstObj {
						headers = append(headers, key)
					}
					sort.Strings(headers)
					
					// 写入表头
					for i, header := range headers {
						col, _ := excelize.ColumnNumberToName(i + 1)
						cell := col + fmt.Sprintf("%d", currentRow)
						excel.SetCellValue(sheetName, cell, header)
						excel.SetCellStyle(sheetName, cell, cell, headerStyle)
					}
					excel.SetRowHeight(sheetName, currentRow, 20)
					currentRow++
					
					// 写入数据行
					rowNumber := 0
					for _, item := range jsonArray {
						if obj, ok := item.(map[string]interface{}); ok {
							// 确定行样式 (交替行样式)
							rowStyle := dataStyle
							if rowNumber%2 == 1 {
								rowStyle = alternateRowStyle
							}
							
							// 遍历所有标题键
							for i, header := range headers {
								col, _ := excelize.ColumnNumberToName(i + 1)
								cell := col + fmt.Sprintf("%d", currentRow)
								
								// 安全地设置值
								if value, exists := obj[header]; exists {
									excel.SetCellValue(sheetName, cell, value)
								} else {
									excel.SetCellValue(sheetName, cell, "")
								}
								excel.SetCellStyle(sheetName, cell, cell, rowStyle)
							}
							currentRow++
							rowNumber++
						}
					}
					
					fmt.Printf("成功处理JSON数据，共%d行\n", rowNumber)
				}
			}
		default:
			fmt.Printf("不支持的文件类型: %s\n", fileExt)
		}
	} else {
		fmt.Println("警告：未找到有效的数据源路径")
	}
	
	// 添加设备信息
	if params, ok := dataMap["parameters"].(map[string]interface{}); ok {
		if equipInfo, ok := params["equipmentInfo"].(map[string]interface{}); ok && len(equipInfo) > 0 {
			currentRow++
			excel.SetCellValue(sheetName, "A"+fmt.Sprintf("%d", currentRow), "设备信息")
			excel.MergeCell(sheetName, "A"+fmt.Sprintf("%d", currentRow), "C"+fmt.Sprintf("%d", currentRow))
			excel.SetCellStyle(sheetName, "A"+fmt.Sprintf("%d", currentRow), "C"+fmt.Sprintf("%d", currentRow), headerStyle)
			currentRow++
			
			// 设备信息表头
			excel.SetCellValue(sheetName, "A"+fmt.Sprintf("%d", currentRow), "项目")
			excel.SetCellValue(sheetName, "B"+fmt.Sprintf("%d", currentRow), "值")
			excel.SetCellStyle(sheetName, "A"+fmt.Sprintf("%d", currentRow), "B"+fmt.Sprintf("%d", currentRow), headerStyle)
			currentRow++
			
			// 设备信息数据
			rowNumber := 0
			for key, value := range equipInfo {
				// 确定行样式 (交替行样式)
				rowStyle := dataStyle
				if rowNumber%2 == 1 {
					rowStyle = alternateRowStyle
				}
				
				excel.SetCellValue(sheetName, "A"+fmt.Sprintf("%d", currentRow), key)
				excel.SetCellValue(sheetName, "B"+fmt.Sprintf("%d", currentRow), value)
				excel.SetCellStyle(sheetName, "A"+fmt.Sprintf("%d", currentRow), "B"+fmt.Sprintf("%d", currentRow), rowStyle)
				currentRow++
				rowNumber++
			}
		}
	}
	
	// 添加结论
	if params, ok := dataMap["parameters"].(map[string]interface{}); ok {
		if conclusions, ok := params["conclusions"].([]interface{}); ok && len(conclusions) > 0 {
			currentRow++
			excel.SetCellValue(sheetName, "A"+fmt.Sprintf("%d", currentRow), "分析结论")
			excel.MergeCell(sheetName, "A"+fmt.Sprintf("%d", currentRow), "C"+fmt.Sprintf("%d", currentRow))
			excel.SetCellStyle(sheetName, "A"+fmt.Sprintf("%d", currentRow), "C"+fmt.Sprintf("%d", currentRow), headerStyle)
			currentRow++
			
			// 结论数据
			for i, conclusion := range conclusions {
				rowStyle := dataStyle
				if i%2 == 1 {
					rowStyle = alternateRowStyle
				}
				
				cell := "A" + fmt.Sprintf("%d", currentRow+i)
				excel.SetCellValue(sheetName, cell, conclusion)
				excel.MergeCell(sheetName, cell, "C"+fmt.Sprintf("%d", currentRow+i))
				excel.SetCellStyle(sheetName, cell, "C"+fmt.Sprintf("%d", currentRow+i), rowStyle)
			}
		}
	}
	
	// 自动调整列宽以适应内容
	for i := 1; i <= 10; i++ {
		col, _ := excelize.ColumnNumberToName(i)
		excel.SetColWidth(sheetName, col, col, 15)
	}
	
	// 创建输出目录（如果不存在）
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	
	// 保存Excel文件
	if err := excel.SaveAs(outputPath); err != nil {
		return fmt.Errorf("failed to save Excel document: %w", err)
	}
	
	fmt.Printf("Generated Excel document at %s\n", outputPath)
	return nil
}

// PDFFormatter formats data into a PDF document
type PDFFormatter struct{}

// Format formats the data into a PDF document
func (f *PDFFormatter) Format(data interface{}, outputPath string) error {
	fmt.Println("Generating PDF document...")
	
	// 获取数据内容
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data is not a map")
	}
	
	// 创建一个新的PDF文档
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	// 设置中文字体
	pdf.SetFont("Arial", "B", 16)
	
	// 添加标题
	title := "Report"
	if t, ok := dataMap["title"].(string); ok {
		title = t
	}
	
	pdf.Cell(190, 10, title)
	pdf.Ln(15)
	
	// 如果有副标题，添加副标题
	if params, ok := dataMap["parameters"].(map[string]interface{}); ok {
		if subtitle, ok := params["subtitle"].(string); ok {
			pdf.SetFont("Arial", "I", 12)
			pdf.Cell(190, 10, subtitle)
			pdf.Ln(15)
		}
	}
	
	// 添加设备信息
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 10, "设备信息：")
	pdf.Ln(10)
	
	pdf.SetFont("Arial", "", 10)
	if params, ok := dataMap["parameters"].(map[string]interface{}); ok {
		if equipInfo, ok := params["equipmentInfo"].(map[string]interface{}); ok {
			for key, value := range equipInfo {
				pdf.Cell(190, 8, fmt.Sprintf("%s: %v", key, value))
				pdf.Ln(8)
			}
		}
	}
	
	// 添加结论
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 10, "分析结论：")
	pdf.Ln(10)
	
	pdf.SetFont("Arial", "", 10)
	if params, ok := dataMap["parameters"].(map[string]interface{}); ok {
		if conclusions, ok := params["conclusions"].([]interface{}); ok {
			for _, conclusion := range conclusions {
				pdf.Cell(190, 8, fmt.Sprintf("- %v", conclusion))
				pdf.Ln(8)
			}
		}
	}
	
	// 创建输出目录（如果不存在）
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	
	// 保存PDF文档
	if err := pdf.OutputFileAndClose(outputPath); err != nil {
		return fmt.Errorf("failed to save PDF file: %w", err)
	}
	
	fmt.Printf("Generated PDF at %s\n", outputPath)
	return nil
}

// HTMLFormatter formats data into an HTML document
type HTMLFormatter struct{}

// Format formats the data into an HTML document
func (f *HTMLFormatter) Format(data interface{}, outputPath string) error {
	// Check if template path is provided in the data
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("data is not a map")
	}

	templatePath, ok := dataMap["_templatePath"].(string)
	if !ok {
		return fmt.Errorf("template path not found in data")
	}

	// Read the template file
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	// Parse the template
	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create output directory if it doesn't exist
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Apply the template with the provided data
	if err := tmpl.Execute(outputFile, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	fmt.Printf("Generated HTML document at %s\n", outputPath)
	return nil
}

// JSONFormatter formats data into a JSON file
type JSONFormatter struct{}

// Format formats the data into a JSON file
func (f *JSONFormatter) Format(data interface{}, outputPath string) error {
	fmt.Println("Generating JSON document...")
	
	// 将数据序列化为格式化的JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}
	
	// 创建输出目录（如果不存在）
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}
	
	// 写入文件
	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}
	
	fmt.Printf("Generated JSON file at %s\n", outputPath)
	return nil
}

// CreateFormatter creates a formatter based on the output format
func CreateFormatter(format string) (OutputFormatter, error) {
	switch strings.ToLower(format) {
	case "docx", ".docx":
		return &DocxFormatter{}, nil
	case "xlsx", ".xlsx":
		return &ExcelFormatter{}, nil
	case "pdf", ".pdf":
		return &PDFFormatter{}, nil
	case "html", ".html":
		return &HTMLFormatter{}, nil
	case "json", ".json":
		return &JSONFormatter{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

// GetFormatFromPath extracts the format from the file path
func GetFormatFromPath(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return ""
	}
	return strings.ToLower(ext[1:])
}
