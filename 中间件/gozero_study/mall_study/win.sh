# 导入 Windows 环境变量到 WSL
while IFS='=' read -r var val; do
    if [[ -n $var && -n $val ]]; then
        # 移除 Windows 换行符
        val="${val//$'\r'/}"
        export "$var=$val"
    fi
done < <(cmd.exe /C set)