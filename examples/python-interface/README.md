# ReportCraft Python接口

这个Python模块提供了访问ReportCraft报表生成功能的接口，使其能够轻松集成到Python应用中。

## 特点

- 简单易用的API，隐藏底层复杂性
- 支持所有ReportCraft报表类型和配置
- 提供便捷的报表类型专用方法
- 自动查找ReportCraft可执行文件
- 支持命令行和程序化使用

## 安装

将`reportcraft.py`模块复制到您的项目目录，或者添加到Python路径中即可使用。

## 基本用法

```python
from reportcraft import ReportCraft

# 初始化ReportCraft接口
rc = ReportCraft()  # 自动查找可执行文件
# 或者指定可执行文件位置
# rc = ReportCraft("/path/to/reportcraft")

# 方法1：使用配置文件生成报告
success, message = rc.generate_report("path/to/config.json", "output.html")

# 方法2：使用字典配置生成报告
config = {
    "reportType": "trend-report",
    "outputPath": "trend-report.html",
    "template": "templates/trend.html",
    "outputFormat": "html",
    "parameters": {
        "title": "趋势分析报告"
    },
    "dataSources": [
        {
            "type": "csv",
            "path": "data.csv",
            "hasHeader": True
        }
    ]
}
success, message = rc.generate_report(config)

# 检查结果
if success:
    print(f"报告生成成功: {message}")
else:
    print(f"报告生成失败: {message}")
```

## 专用报表类型API

### 频谱报告

```python
rc.create_spectrum_report(
    title="3D频谱分析",
    data_source={
        "type": "csv",
        "path": "spectrum_data.csv",
        "hasHeader": True
    },
    output_path="spectrum-report.html",
    chart_type="surface3d",
    x_column="频率",
    y_column="时间",
    z_column="幅值"
)
```

### 波形报告

```python
rc.create_waveform_report(
    title="振动波形分析",
    data_source={
        "type": "csv",
        "path": "waveform_data.csv",
        "hasHeader": True
    },
    output_path="waveform-report.html",
    time_column="时间",
    channels=[
        {"name": "通道1", "color": "red"},
        {"name": "通道2", "color": "blue"}
    ]
)
```

### 趋势报告

```python
rc.create_trend_report(
    title="设备运行参数趋势",
    data_source={
        "type": "csv",
        "path": "trend_data.csv",
        "hasHeader": True
    },
    output_path="trend-report.html",
    time_column="时间",
    series=[
        {"name": "温度", "color": "red"},
        {"name": "压力", "color": "blue"}
    ]
)
```

## 命令行使用

```bash
# 基本用法
python reportcraft.py -c config.json -o output.html

# 指定ReportCraft可执行文件
python reportcraft.py -c config.json -o output.html -e /path/to/reportcraft
```

## 示例

详见`examples.py`文件中的完整示例，包含多种报表类型和用法。

运行所有示例：
```bash
python examples.py
```

运行特定示例：
```bash
python examples.py 1  # 运行示例1
```

## 集成到项目

### Web应用集成

```python
from flask import Flask, request, jsonify
from reportcraft import ReportCraft

app = Flask(__name__)
rc = ReportCraft()

@app.route('/generate-report', methods=['POST'])
def generate_report():
    config = request.json
    success, message = rc.generate_report(config)
    return jsonify({
        "success": success,
        "message": message
    })

if __name__ == '__main__':
    app.run(debug=True)
```

### 数据分析流程集成

```python
import pandas as pd
from reportcraft import ReportCraft

# 数据处理
df = pd.read_csv('raw_data.csv')
processed_df = process_data(df)  # 自定义处理函数
processed_df.to_csv('processed_data.csv', index=False)

# 生成报告
rc = ReportCraft()
rc.create_trend_report(
    title="数据分析结果",
    data_source={
        "type": "csv",
        "path": "processed_data.csv",
        "hasHeader": True
    },
    output_path="analysis-report.html"
)
```
