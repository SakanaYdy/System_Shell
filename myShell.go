package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"os/exec"
	"strings"
	"system/api"
)

func execInput(input string) error {

	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)

	}

	// Pass the program and the arguments separately.
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func testApi(input string) {

	//fmt.Print(">")

	apiKey := "sk-5806dc5ecf1b4438bc35629dacbf0553"

	client := api.NewTongYiClient(apiKey)

	// 设置 prompt 和历史对话
	prompt := input + " 仅仅需要给出正确的Linux指令，其他的都不要，字符串格式即可"
	// 示例：设置正确的历史对话格式
	//history := []map[string]string{
	//	{
	//		"user": "什么是Go语言？",                    // 用户提问
	//		"bot":  "Go语言是一种开源编程语言，设计上以简洁、高效为目标。", // 机器人回应
	//	},
	//}

	// 调用 GenerateText 函数
	resp, err := client.GenerateText(context.Background(), prompt)
	if err != nil {
		log.Fatalf("调用 GenerateText 失败: %v", err)
	}

	ans := resp.Output.Text
	err = execInput(ans[1 : len(ans)-1])
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 清空当前行并打印新的命令
func clearAndPrintCommand(commamd, dir string) {
	// 将光标移动到行首并清空当前命令
	fmt.Print("\r")                 // 回车符：移动到行首
	fmt.Print("                ")   // 打印空格覆盖当前行，确保清除原来的命令
	fmt.Print("\r")                 // 再次回到行首
	fmt.Print(dir + "> " + commamd) // 打印新的命令
}

func main() {

	//Queue := api.Queue{}
	//recordQueue := Queue.Init()
	//
	//reader := bufio.NewReader(os.Stdin)

	keysEvents, err := keyboard.GetKeys(10)
	var commandSlice []string
	commandIndex := 0

	if err != nil {
		fmt.Println("keyboard error")
	}

	defer func() {
		_ = keyboard.Close()
	}()

	var commamd string
	dir, err2 := os.Getwd()
	if err2 != nil {
		fmt.Println("文件目录获取失败")
	}
	fmt.Print(dir + "> ")
	for {
		// 获取当前工作目录
		dir, err2 := os.Getwd()
		if err2 != nil {
			fmt.Println("文件目录获取失败")
			return
		}

		// 获取键盘事件
		event := <-keysEvents

		// 处理按键事件
		switch {
		case event.Key == keyboard.KeyArrowUp:
			// 上箭头，获取上一个命令
			commamd = api.GetCommand(&commandIndex, commandSlice, true)
			clearAndPrintCommand(commamd, dir)

		case event.Key == keyboard.KeyArrowDown:
			// 下箭头，获取下一个命令
			commamd = api.GetCommand(&commandIndex, commandSlice, false)
			clearAndPrintCommand(commamd, dir)

		case "A" <= string(event.Rune) && string(event.Rune) <= "z":
			// 拼接字母或数字到命令
			commamd = commamd + string(event.Rune)
			clearAndPrintCommand(commamd, dir) // 清空并重新打印当前命令

		case event.Key == keyboard.KeyEnter:
			// 回车键，执行命令
			fmt.Println()
			//fmt.Println(commamd)
			if commamd != "" {
				// 保存命令到历史
				api.SaveCommand(commamd, &commandSlice)
				commandIndex = len(commandSlice) // 更新命令索引
			}
			if err := execInput(commamd); err != nil {
				testApi(commamd)
			}
			commamd = "" // 清空当前命令
			fmt.Print(dir + "> ")

			//	input, err := reader.ReadString('\n')
			//	if err != nil {
			//		fmt.Fprintln(os.Stderr, err)
			//	}
			//
			//	api.SaveCommand(input, &commandSlice)
			//
			//	if recordQueue.GetSize() == 3 {
			//		recordQueue.Show()
			//	}
			//
			//	//  如果用户指令无法识别，调用AI接口修正
			//	if err = execInput(input); err != nil {
			//		//testApi()
			//		testApi(input)
			//		//fmt.Fprintln(os.Stderr, err)
			//	}
			//}
		}
	}
}
