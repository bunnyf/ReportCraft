@echo off
setlocal enabledelayedexpansion

:: ReportCraft批量报告生成脚本
:: 此脚本自动生成所有示例报告

echo ReportCraft批量报告生成工具
echo =============================
echo.

:: 设置参数
set "REPORTCRAFT_PATH=..\..\reportcraft.exe"
set "EXAMPLES_DIR=.."
set "OUTPUT_DIR=..\output\batch"
set "LOG_FILE=%OUTPUT_DIR%\generation_log.txt"

:: 创建输出目录
if not exist "%OUTPUT_DIR%" (
    echo 创建输出目录: %OUTPUT_DIR%
    mkdir "%OUTPUT_DIR%"
)

:: 清空日志文件
echo ReportCraft批量生成日志 > "%LOG_FILE%"
echo 生成时间: %date% %time% >> "%LOG_FILE%"
echo. >> "%LOG_FILE%"

:: 检查ReportCraft是否存在
if not exist "%REPORTCRAFT_PATH%" (
    echo 错误: ReportCraft可执行文件未找到: %REPORTCRAFT_PATH%
    echo 请确保可执行文件位于正确位置，或者修改脚本中的REPORTCRAFT_PATH变量
    goto :error
)

:: 查找所有JSON配置文件
echo 查找所有JSON配置文件...
set "CONFIG_COUNT=0"
for %%f in ("%EXAMPLES_DIR%\*.json") do (
    set /a CONFIG_COUNT+=1
)

if %CONFIG_COUNT% equ 0 (
    echo 错误: 未找到任何JSON配置文件
    goto :error
)

echo 找到 %CONFIG_COUNT% 个配置文件
echo.

:: 生成所有报告
echo 开始生成报告...
echo.

set "SUCCESS_COUNT=0"
set "FAILED_COUNT=0"

for %%f in ("%EXAMPLES_DIR%\*.json") do (
    set "CONFIG_FILE=%%f"
    set "REPORT_NAME=%%~nf"
    
    echo 处理: !REPORT_NAME!
    
    :: 确定输出文件
    for /f "tokens=1,* delims=[]" %%a in ('type "!CONFIG_FILE!" ^| findstr "outputPath"') do (
        set "OUTPUT_LINE=%%b"
    )
    
    :: 提取文件扩展名
    set "OUTPUT_EXT=html"
    echo !OUTPUT_LINE! | findstr /C:".docx" > nul && set "OUTPUT_EXT=docx"
    echo !OUTPUT_LINE! | findstr /C:".pdf" > nul && set "OUTPUT_EXT=pdf"
    echo !OUTPUT_LINE! | findstr /C:".xlsx" > nul && set "OUTPUT_EXT=xlsx"
    
    :: 设置新的输出路径
    set "OUTPUT_FILE=%OUTPUT_DIR%\!REPORT_NAME!.!OUTPUT_EXT!"
    
    echo   源配置: !CONFIG_FILE!
    echo   输出文件: !OUTPUT_FILE!
    
    :: 调用ReportCraft
    echo   正在生成...
    "%REPORTCRAFT_PATH%" -config="!CONFIG_FILE!" -output="!OUTPUT_FILE!" > "%TEMP%\rc_output.txt" 2>&1
    
    if %ERRORLEVEL% equ 0 (
        echo   [成功] 报告生成完成
        set /a SUCCESS_COUNT+=1
        echo [成功] !REPORT_NAME! >> "%LOG_FILE%"
        echo   配置: !CONFIG_FILE! >> "%LOG_FILE%"
        echo   输出: !OUTPUT_FILE! >> "%LOG_FILE%"
        type "%TEMP%\rc_output.txt" >> "%LOG_FILE%"
    ) else (
        echo   [失败] 报告生成失败
        set /a FAILED_COUNT+=1
        echo [失败] !REPORT_NAME! >> "%LOG_FILE%"
        echo   配置: !CONFIG_FILE! >> "%LOG_FILE%"
        echo   错误: >> "%LOG_FILE%"
        type "%TEMP%\rc_output.txt" >> "%LOG_FILE%"
    )
    
    echo. >> "%LOG_FILE%"
    echo.
)

:: 生成摘要
echo 生成摘要:
echo 总共: %CONFIG_COUNT%, 成功: %SUCCESS_COUNT%, 失败: %FAILED_COUNT%
echo.
echo 详细日志保存在: %LOG_FILE%

echo. >> "%LOG_FILE%"
echo 摘要: >> "%LOG_FILE%"
echo 总共: %CONFIG_COUNT%, 成功: %SUCCESS_COUNT%, 失败: %FAILED_COUNT% >> "%LOG_FILE%"

:: 自动打开输出目录
if %SUCCESS_COUNT% gtr 0 (
    echo 是否打开输出目录? [Y/N]
    set /p OPEN_DIR=
    if /i "!OPEN_DIR!"=="Y" (
        start "" "%OUTPUT_DIR%"
    )
)

goto :end

:error
echo.
echo 脚本执行失败，请检查上述错误。
exit /b 1

:end
echo.
echo 批处理执行完成。
endlocal
exit /b 0
