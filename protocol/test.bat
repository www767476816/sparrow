@echo off


if not exist error_msg (
   echo "文件夹已存在"
) else (
md error_msg
)


pause