package api

// 实现指令的记录
// 暂时使用一个队列来模拟

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
