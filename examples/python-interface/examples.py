#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
ReportCraft Python接口使用示例
"""

import os
import sys
from reportcraft import ReportCraft

# 获取脚本所在目录
SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
# 获取示例目录
EXAMPLES_DIR = os.path.dirname(SCRIPT_DIR)
# 获取数据目录
DATA_DIR = os.path.join(EXAMPLES_DIR, "data")
# 获取输出目录
OUTPUT_DIR = os.path.join(EXAMPLES_DIR, "output", "python")


def ensure_dir(directory):
    """确保目录存在"""
    if not os.path.exists(directory):
        os.makedirs(directory)


def example_1_use_config_file():
    """示例1：使用已有配置文件生成报告"""
    print("\n示例1：使用已有配置文件生成报告")
    
    # 初始化ReportCraft
    rc = ReportCraft()
    
    # 配置文件路径
    config_path = os.path.join(EXAMPLES_DIR, "spectrum-3d.json")
    
    # 输出路径
    ensure_dir(OUTPUT_DIR)
    output_path = os.path.join(OUTPUT_DIR, "spectrum-3d-from-python.html")
    
    # 生成报告
    success, message = rc.generate_report(config_path, output_path)
    
    # 打印结果
    if success:
        print(f"✓ 成功生成报告: {message}")
    else:
        print(f"✗ 报告生成失败: {message}")


def example_2_create_spectrum_report():
    """示例2：使用API创建频谱报告"""
    print("\n示例2：使用API创建频谱报告")
    
    # 初始化ReportCraft
    rc = ReportCraft()
    
    # 定义数据源
    data_source = {
        "type": "csv",
        "path": os.path.join(DATA_DIR, "3d_spectrum_data.csv"),
        "hasHeader": True
    }
    
    # 输出路径
    ensure_dir(OUTPUT_DIR)
    output_path = os.path.join(OUTPUT_DIR, "spectrum-api-created.html")
    
    # 创建报告
    success, message = rc.create_spectrum_report(
        title="使用Python API创建的3D频谱报告",
        data_source=data_source,
        output_path=output_path,
        chart_type="surface3d",
        x_column="频率",
        y_column="时间",
        z_column="幅值",
        colorScale="Viridis",
        colorOpacity=0.85,
        meshDensity={"x": 50, "y": 50},
        lighting={
            "ambient": 0.65,
            "diffuse": 0.85,
            "specular": 0.15
        },
        interactiveControls=True
    )
    
    # 打印结果
    if success:
        print(f"✓ 成功创建频谱报告: {message}")
    else:
        print(f"✗ 频谱报告创建失败: {message}")


def example_3_create_trend_report():
    """示例3：创建趋势分析报告"""
    print("\n示例3：创建趋势分析报告")
    
    # 初始化ReportCraft
    rc = ReportCraft()
    
    # 定义数据源
    data_source = {
        "type": "csv",
        "path": os.path.join(DATA_DIR, "multi_trend_data.csv"),
        "hasHeader": True
    }
    
    # 输出路径
    ensure_dir(OUTPUT_DIR)
    output_path = os.path.join(OUTPUT_DIR, "trend-analysis-api.html")
    
    # 创建报告
    success, message = rc.create_trend_report(
        title="设备运行参数分析报告",
        data_source=data_source,
        output_path=output_path,
        time_column="时间",
        series=[
            {"name": "温度", "color": "red", "lineWidth": 2, "yAxis": 0},
            {"name": "压力", "color": "blue", "lineWidth": a2, "yAxis": 1},
            {"name": "流量", "color": "green", "lineWidth": 2, "yAxis": 2}
        ],
        y_axes=[
            {"title": "温度 (°C)", "min": 0, "max": 100, "position": "left"},
            {"title": "压力 (MPa)", "min": 0, "max": 10, "position": "right"},
            {"title": "流量 (m³/h)", "min": 0, "max": 250, "position": "right", "offset": 50}
        ],
        annotations=[
            {"x": "2025-02-15T00:00:00Z", "text": "维护", "style": "dashed"},
            {"y": 85, "yAxis": 0, "text": "温度预警", "style": "dotted"}
        ]
    )
    
    # 打印结果
    if success:
        print(f"✓ 成功创建趋势报告: {message}")
    else:
        print(f"✗ 趋势报告创建失败: {message}")


def example_4_create_waveform_report():
    """示例4：创建波形分析报告"""
    print("\n示例4：创建波形分析报告")
    
    # 初始化ReportCraft
    rc = ReportCraft()
    
    # 定义数据源
    data_source = {
        "type": "csv",
        "path": os.path.join(DATA_DIR, "multichannel_waveform.csv"),
        "hasHeader": True
    }
    
    # 输出路径
    ensure_dir(OUTPUT_DIR)
    output_path = os.path.join(OUTPUT_DIR, "waveform-api.html")
    
    # 创建报告
    success, message = rc.create_waveform_report(
        title="多通道振动波形分析",
        data_source=data_source,
        output_path=output_path,
        time_column="时间",
        channels=[
            {"name": "通道1", "color": "#FF5733", "lineWidth": 1.5},
            {"name": "通道2", "color": "#337DFF", "lineWidth": 1.5},
            {"name": "通道3", "color": "#33FF57", "lineWidth": 1.5},
            {"name": "通道4", "color": "#D433FF", "lineWidth": 1.5}
        ],
        showGrid=True,
        showLegend=True,
        enableZoom=True,
        xAxisTitle="时间 (s)",
        yAxisTitle="振动速度 (mm/s)",
        annotations=[
            {"x": 0.15, "text": "启动", "textPosition": "top"},
            {"x": 0.45, "text": "稳定", "textPosition": "top"},
            {"x": 0.75, "text": "负载", "textPosition": "top"}
        ]
    )
    
    # 打印结果
    if success:
        print(f"✓ 成功创建波形报告: {message}")
    else:
        print(f"✗ 波形报告创建失败: {message}")


def example_5_programmatically_create_json():
    """示例5：程序化创建完整配置并生成报告"""
    print("\n示例5：程序化创建完整配置并生成报告")
    
    # 初始化ReportCraft
    rc = ReportCraft()
    
    # 创建完整配置
    config = {
        "reportType": "combined-report",
        "title": "程序化创建的组合报告",
        "description": "这个报告包含了多种图表类型，用于展示设备运行状态",
        "outputPath": os.path.join(OUTPUT_DIR, "programmatic-combined.html"),
        "outputFormat": "html",
        "template": "templates/dashboard.html",
        "parameters": {
            "companyName": "测试公司",
            "deviceName": "测试设备",
            "reportDate": "2025-03-20",
            "author": "Python API",
            "showHeader": True,
            "showFooter": True,
            "includeTimeStamp": True
        },
        "dataSources": [
            {
                "id": "combined_data",
                "type": "json",
                "path": os.path.join(DATA_DIR, "combined_charts_data.json")
            }
        ],
        "sections": [
            {
                "title": "设备健康指数",
                "type": "gauge",
                "dataSource": "combined_data",
                "dataField": "healthIndex",
                "min": 0,
                "max": 100,
                "thresholds": [
                    {"value": 30, "color": "red"},
                    {"value": 70, "color": "yellow"},
                    {"value": 100, "color": "green"}
                ],
                "width": "100%",
                "height": "300px"
            },
            {
                "title": "参数趋势",
                "type": "line",
                "dataSource": "combined_data",
                "dataField": "parameterTrends",
                "xAxis": {"field": "timestamps", "title": "时间"},
                "series": [
                    {"field": "温度", "title": "温度 (°C)", "color": "red"},
                    {"field": "振动", "title": "振动 (mm/s)", "color": "blue"},
                    {"field": "转速", "title": "转速 (RPM)", "color": "green"}
                ],
                "width": "100%",
                "height": "400px"
            },
            {
                "title": "故障概率",
                "type": "bar",
                "dataSource": "combined_data",
                "dataField": "faultProbability",
                "xAxis": {"field": "component", "title": "部件"},
                "yAxis": {"field": "probability", "title": "概率 (%)"},
                "width": "50%",
                "height": "350px"
            },
            {
                "title": "故障类型分布",
                "type": "pie",
                "dataSource": "combined_data",
                "dataField": "faultTypeDistribution",
                "labelField": "type",
                "valueField": "value",
                "width": "50%",
                "height": "350px"
            },
            {
                "title": "频谱分析",
                "type": "bar",
                "dataSource": "combined_data",
                "dataField": "spectrumData",
                "xAxis": {"field": "频率", "title": "频率 (Hz)"},
                "yAxis": {"field": "幅值", "title": "幅值 (mm/s)"},
                "width": "100%",
                "height": "350px"
            }
        ]
    }
    
    # 生成报告
    success, message = rc.generate_report(config)
    
    # 打印结果
    if success:
        print(f"✓ 成功创建组合报告: {message}")
    else:
        print(f"✗ 组合报告创建失败: {message}")


def run_all_examples():
    """运行所有示例"""
    example_1_use_config_file()
    example_2_create_spectrum_report()
    example_3_create_trend_report()
    example_4_create_waveform_report()
    example_5_programmatically_create_json()


if __name__ == "__main__":
    # 修复代码中的错误
    example_3_create_trend_report.__code__ = compile(
        example_3_create_trend_report.__code__.co_consts[0].replace("a2", "2"),
        "<string>",
        "exec"
    )
    
    if len(sys.argv) > 1:
        # 运行特定示例
        example_num = sys.argv[1]
        if example_num == "1":
            example_1_use_config_file()
        elif example_num == "2":
            example_2_create_spectrum_report()
        elif example_num == "3":
            example_3_create_trend_report()
        elif example_num == "4":
            example_4_create_waveform_report()
        elif example_num == "5":
            example_5_programmatically_create_json()
        else:
            print(f"未知示例: {example_num}")
            print("可用示例: 1, 2, 3, 4, 5")
    else:
        # 运行所有示例
        run_all_examples()
