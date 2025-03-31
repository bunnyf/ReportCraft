#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
ReportCraft Python接口

这个模块提供了从Python中调用ReportCraft的接口，方便在Python应用中集成报表生成功能。
"""

import json
import os
import platform
import subprocess
import tempfile
from typing import Dict, List, Union, Optional, Any, Tuple


class ReportCraft:
    """ReportCraft Python接口类"""

    def __init__(self, executable_path: str = None):
        """
        初始化ReportCraft接口

        Args:
            executable_path: ReportCraft可执行文件的路径。如果为None，将在默认位置查找。
        """
        self.executable_path = executable_path or self._find_executable()
        if not os.path.exists(self.executable_path):
            raise FileNotFoundError(f"ReportCraft可执行文件未找到: {self.executable_path}")

    def _find_executable(self) -> str:
        """
        在默认位置查找ReportCraft可执行文件

        Returns:
            找到的可执行文件路径
        """
        # 获取平台信息
        system = platform.system()
        
        # 在脚本所在目录的父目录中查找可执行文件
        script_dir = os.path.dirname(os.path.abspath(__file__))
        parent_dir = os.path.dirname(os.path.dirname(script_dir))
        
        # 检查平台并构造可执行文件名
        if system == "Windows":
            executable = "reportcraft.exe"
        else:
            executable = "reportcraft"
            
        # 可能的路径列表
        paths = [
            os.path.join(parent_dir, executable),                # 根目录
            os.path.join(parent_dir, "bin", executable),         # bin目录
            os.path.join(parent_dir, "build", executable),       # build目录
            os.path.join(parent_dir, "cmd", "reportcraft", executable)  # cmd/reportcraft目录
        ]
        
        # 尝试找到可执行文件
        for path in paths:
            if os.path.exists(path):
                return path
                
        # 如果未找到，返回默认路径，之后会抛出更明确的错误
        return os.path.join(parent_dir, executable)

    def generate_report(self, config: Union[Dict, str], output_path: str = None) -> Tuple[bool, str]:
        """
        生成报告

        Args:
            config: 报告配置，可以是配置字典或配置文件路径
            output_path: 输出文件路径，如果为None则使用配置中的路径

        Returns:
            (成功标志, 输出消息)
        """
        # 处理配置
        if isinstance(config, dict):
            # 如果提供了输出路径，更新配置
            if output_path:
                config["outputPath"] = output_path
                
            # 创建临时配置文件
            fd, temp_config = tempfile.mkstemp(suffix='.json')
            try:
                with os.fdopen(fd, 'w') as f:
                    json.dump(config, f, indent=2)
                config_path = temp_config
            except Exception as e:
                os.unlink(temp_config)
                return False, f"创建临时配置文件失败: {str(e)}"
        else:
            # 使用提供的配置文件路径
            config_path = config
            
            # 如果提供了输出路径，尝试修改配置文件
            if output_path:
                try:
                    with open(config_path, 'r') as f:
                        config_data = json.load(f)
                    
                    config_data["outputPath"] = output_path
                    
                    fd, temp_config = tempfile.mkstemp(suffix='.json')
                    with os.fdopen(fd, 'w') as f:
                        json.dump(config_data, f, indent=2)
                    config_path = temp_config
                except Exception as e:
                    return False, f"修改配置文件失败: {str(e)}"

        # 调用ReportCraft
        try:
            cmd = [self.executable_path, f"-config={config_path}"]
            result = subprocess.run(cmd, capture_output=True, text=True)
            success = result.returncode == 0
            message = result.stdout if success else result.stderr
            
            # 提取生成的文件路径
            if success:
                for line in result.stdout.splitlines():
                    if "Generated" in line and "at" in line:
                        parts = line.split("at")
                        if len(parts) > 1:
                            generated_path = parts[1].strip()
                            message += f"\n生成的文件: {generated_path}"
                            break
            
            return success, message
        except Exception as e:
            return False, f"执行ReportCraft失败: {str(e)}"
        finally:
            # 清理临时文件
            if isinstance(config, dict) or output_path:
                try:
                    os.unlink(config_path)
                except:
                    pass

    def create_spectrum_report(self, 
                               title: str, 
                               data_source: Dict[str, Any], 
                               output_path: str, 
                               output_format: str = "html",
                               template: str = "templates/spectrum.html",
                               **kwargs) -> Tuple[bool, str]:
        """
        创建频谱分析报告
        
        Args:
            title: 报告标题
            data_source: 数据源配置
            output_path: 输出文件路径
            output_format: 输出格式 (html, docx, pdf)
            template: 模板路径
            **kwargs: 其他参数
            
        Returns:
            (成功标志, 输出消息)
        """
        # 构建基本配置
        config = {
            "reportType": "spectrum-report",
            "outputPath": output_path,
            "template": template,
            "outputFormat": output_format,
            "parameters": {
                "title": title,
                **kwargs
            },
            "dataSources": [data_source]
        }
        
        return self.generate_report(config)
        
    def create_waveform_report(self, 
                               title: str, 
                               data_source: Dict[str, Any], 
                               output_path: str, 
                               output_format: str = "html",
                               template: str = "templates/waveform.html",
                               **kwargs) -> Tuple[bool, str]:
        """
        创建波形分析报告
        
        Args:
            title: 报告标题
            data_source: 数据源配置
            output_path: 输出文件路径
            output_format: 输出格式 (html, docx, pdf)
            template: 模板路径
            **kwargs: 其他参数
            
        Returns:
            (成功标志, 输出消息)
        """
        # 构建基本配置
        config = {
            "reportType": "waveform-report",
            "outputPath": output_path,
            "template": template,
            "outputFormat": output_format,
            "parameters": {
                "title": title,
                **kwargs
            },
            "dataSources": [data_source]
        }
        
        return self.generate_report(config)
    
    def create_trend_report(self, 
                            title: str, 
                            data_source: Dict[str, Any], 
                            output_path: str, 
                            output_format: str = "html",
                            template: str = "templates/trend.html",
                            **kwargs) -> Tuple[bool, str]:
        """
        创建趋势分析报告
        
        Args:
            title: 报告标题
            data_source: 数据源配置
            output_path: 输出文件路径
            output_format: 输出格式 (html, docx, pdf)
            template: 模板路径
            **kwargs: 其他参数
            
        Returns:
            (成功标志, 输出消息)
        """
        # 构建基本配置
        config = {
            "reportType": "trend-report",
            "outputPath": output_path,
            "template": template,
            "outputFormat": output_format,
            "parameters": {
                "title": title,
                **kwargs
            },
            "dataSources": [data_source]
        }
        
        return self.generate_report(config)


# 命令行接口
if __name__ == "__main__":
    import argparse
    
    parser = argparse.ArgumentParser(description="ReportCraft Python接口")
    parser.add_argument("-c", "--config", help="配置文件路径")
    parser.add_argument("-o", "--output", help="输出文件路径")
    parser.add_argument("-e", "--executable", help="ReportCraft可执行文件路径")
    
    args = parser.parse_args()
    
    if not args.config:
        print("错误: 必须提供配置文件路径")
        parser.print_help()
        exit(1)
        
    try:
        rc = ReportCraft(args.executable)
        success, message = rc.generate_report(args.config, args.output)
        
        if success:
            print(f"成功: {message}")
            exit(0)
        else:
            print(f"失败: {message}")
            exit(1)
    except Exception as e:
        print(f"错误: {str(e)}")
        exit(1)
