#!/bin/bash

# 编译项目
echo "编译 ReportCraft..."
go build -o reportcraft ./cmd/reportcraft

# 确保输出目录存在
mkdir -p output

# 测试所有格式
echo "测试 PDF 格式输出..."
cd examples && ../reportcraft --config=spectrum-waterfall-enhanced-pdf.json
cd ..

echo "测试 Word 格式输出..."
cd examples && ../reportcraft --config=spectrum-waterfall-enhanced-word.json
cd ..

echo "测试 Excel 格式输出..."
cd examples && ../reportcraft --config=spectrum-waterfall-enhanced-excel.json
cd ..

echo "测试 JSON 格式输出..."
cd examples && ../reportcraft --config=spectrum-waterfall-enhanced-json.json
cd ..

echo "测试 HTML 格式输出..."
cd examples && ../reportcraft --config=spectrum-waterfall-enhanced.json
cd ..

echo "所有格式测试完成，输出文件位于 output/ 目录"
ls -la output/
