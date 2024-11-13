# System_Shell
自定义实现基于Linux的Shell页面

## 环境需要

```shell
// 切换CN镜像源
go env -w GOPROXY=https://goproxy.cn,direct

go get github.com/eiannone/keyboard
```



## 功能实现

### 1. 查看历史记录指令

基于队列实现，可以设置能够保存的最大容量。

```go
// GetCommand 提取输入命令
func GetCommand(commandIndex *int, commandSlice []string, true bool) (commamd string) {
	if true {
		if *commandIndex > 0 {
			*commandIndex--
		}
	} else {
		if *commandIndex < len(commandSlice) {
			*commandIndex++
		}
	}
	if *commandIndex < len(commandSlice) {
		return commandSlice[*commandIndex]
	} else {
		return ""
	}
}

// SaveCommand 保存输入命令
func SaveCommand(str string, commandSlice *[]string) {
	*commandSlice = append(*commandSlice, str)
	if len(*commandSlice) > 10 {
		*commandSlice = (*commandSlice)[1:]
	}
}
```

### 2 . 获取当前工作目录并显示

```go
dir, err2 := os.Getwd()
```

### 3. 接入AI接口，实现指令提示以及询问

```go
//  如果用户指令无法识别，调用AI接口修正
if err = execInput(input); err != nil {
	testApi(input)
}
```

