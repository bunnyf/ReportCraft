# 添加高级图表可视化功能

添加了三种新的图表可视化类型，以增强振动分析报告的数据可视化能力：热图、三维表面图和瀑布图。

## 主要更改

1. **新增图表类型**：
   
   - 热图（Heatmap）：用于显示二维空间内的强度变化
   - 三维表面图（3D Surface Plot）：提供完整的三维数据可视化
   - 瀑布图（Waterfall Plot）：用于分析随时间或条件变化的频谱

2. **修改的文件**：
   
   - `spectrum_report.go`：添加新图表类型的处理逻辑
   - `README.md` 和 `README_zh.md`：更新文档，添加新图表类型说明
   - `user-guide.md`：添加新图表类型的详细用法说明
   - `spectrum-report-embedded.json` 和 `spectrum-report-csv.json`：更新示例配置

3. **新增数据文件**：
   
   - 添加了示例CSV数据文件用于测试新图表类型：
     - `waveform.csv`：时域波形数据
     - `spectrum.csv`：频谱数据
     - `heatmap.csv`：热图数据
     - `surface.csv`：3D表面图数据
     - `waterfall.csv`：瀑布图数据

## 技术细节

- 热图配置支持多种色彩映射方案（viridis、jet、plasma等）
- 3D表面图支持线框模式和旋转控制
- 瀑布图支持渐变色和透视角度控制
- 所有新图表类型都支持完整的轴配置和样式自定义

## 备注

此次更新大大增强了振动分析报告的数据可视化能力，特别是对于需要分析振动数据随时间变化或多维关系的场景非常有用。
