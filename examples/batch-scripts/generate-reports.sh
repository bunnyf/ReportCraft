#!/bin/bash

# ReportCraft批量报告生成脚本
# 此脚本自动生成所有示例报告

echo "ReportCraft批量报告生成工具"
echo "============================="
echo

# 设置参数
REPORTCRAFT_PATH="../../reportcraft"
EXAMPLES_DIR=".."
OUTPUT_DIR="../output/batch"
LOG_FILE="${OUTPUT_DIR}/generation_log.txt"

# 创建输出目录
if [ ! -d "$OUTPUT_DIR" ]; then
    echo "创建输出目录: $OUTPUT_DIR"
    mkdir -p "$OUTPUT_DIR"
fi

# 清空日志文件
echo "ReportCraft批量生成日志" > "$LOG_FILE"
echo "生成时间: $(date)" >> "$LOG_FILE"
echo >> "$LOG_FILE"

# 检查ReportCraft是否存在
if [ ! -f "$REPORTCRAFT_PATH" ]; then
    echo "错误: ReportCraft可执行文件未找到: $REPORTCRAFT_PATH"
    echo "请确保可执行文件位于正确位置，或者修改脚本中的REPORTCRAFT_PATH变量"
    exit 1
fi

# 确保可执行
if [ ! -x "$REPORTCRAFT_PATH" ]; then
    echo "添加执行权限到 $REPORTCRAFT_PATH"
    chmod +x "$REPORTCRAFT_PATH"
fi

# 查找所有JSON配置文件
echo "查找所有JSON配置文件..."
CONFIG_FILES=("$EXAMPLES_DIR"/*.json)
CONFIG_COUNT=${#CONFIG_FILES[@]}

if [ $CONFIG_COUNT -eq 0 ]; then
    echo "错误: 未找到任何JSON配置文件"
    exit 1
fi

echo "找到 $CONFIG_COUNT 个配置文件"
echo

# 生成所有报告
echo "开始生成报告..."
echo

SUCCESS_COUNT=0
FAILED_COUNT=0

for CONFIG_FILE in "${CONFIG_FILES[@]}"; do
    REPORT_NAME=$(basename "$CONFIG_FILE" .json)
    
    echo "处理: $REPORT_NAME"
    
    # 确定输出文件
    # 提取文件扩展名，默认为html
    OUTPUT_EXT="html"
    if grep -q "\"outputFormat\"[[:space:]]*:[[:space:]]*\"docx\"" "$CONFIG_FILE"; then
        OUTPUT_EXT="docx"
    elif grep -q "\"outputFormat\"[[:space:]]*:[[:space:]]*\"pdf\"" "$CONFIG_FILE"; then
        OUTPUT_EXT="pdf"
    elif grep -q "\"outputFormat\"[[:space:]]*:[[:space:]]*\"xlsx\"" "$CONFIG_FILE"; then
        OUTPUT_EXT="xlsx"
    fi
    
    # 设置新的输出路径
    OUTPUT_FILE="${OUTPUT_DIR}/${REPORT_NAME}.${OUTPUT_EXT}"
    
    echo "  源配置: $CONFIG_FILE"
    echo "  输出文件: $OUTPUT_FILE"
    
    # 调用ReportCraft
    echo "  正在生成..."
    OUTPUT=$("$REPORTCRAFT_PATH" -config="$CONFIG_FILE" -output="$OUTPUT_FILE" 2>&1)
    RC_RESULT=$?
    
    if [ $RC_RESULT -eq 0 ]; then
        echo "  [成功] 报告生成完成"
        ((SUCCESS_COUNT++))
        echo "[成功] $REPORT_NAME" >> "$LOG_FILE"
        echo "  配置: $CONFIG_FILE" >> "$LOG_FILE"
        echo "  输出: $OUTPUT_FILE" >> "$LOG_FILE"
        echo "$OUTPUT" >> "$LOG_FILE"
    else
        echo "  [失败] 报告生成失败"
        ((FAILED_COUNT++))
        echo "[失败] $REPORT_NAME" >> "$LOG_FILE"
        echo "  配置: $CONFIG_FILE" >> "$LOG_FILE"
        echo "  错误:" >> "$LOG_FILE"
        echo "$OUTPUT" >> "$LOG_FILE"
    fi
    
    echo >> "$LOG_FILE"
    echo
done

# 生成摘要
echo "生成摘要:"
echo "总共: $CONFIG_COUNT, 成功: $SUCCESS_COUNT, 失败: $FAILED_COUNT"
echo
echo "详细日志保存在: $LOG_FILE"

echo >> "$LOG_FILE"
echo "摘要:" >> "$LOG_FILE"
echo "总共: $CONFIG_COUNT, 成功: $SUCCESS_COUNT, 失败: $FAILED_COUNT" >> "$LOG_FILE"

# 自动打开输出目录
if [ $SUCCESS_COUNT -gt 0 ]; then
    echo "是否打开输出目录? [Y/N]"
    read -r OPEN_DIR
    if [[ "$OPEN_DIR" == [Yy]* ]]; then
        if [ "$(uname)" == "Darwin" ]; then
            # macOS
            open "$OUTPUT_DIR"
        elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
            # Linux
            if command -v xdg-open > /dev/null; then
                xdg-open "$OUTPUT_DIR"
            else
                echo "无法自动打开目录，请手动浏览: $OUTPUT_DIR"
            fi
        fi
    fi
fi

echo
echo "批处理执行完成。"
exit 0
