package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
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

	fmt.Print(">")

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
	// 输出生成的文本内容
	//fmt.Println("生成的文本:")
	//fmt.Println(ans) // 假设 TongYiRsp 结构体已经处理了响应数据的提取
	err = execInput(ans[1 : len(ans)-1])
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		//  如果用户指令无法识别，调用AI接口修正
		if err = execInput(input); err != nil {
			//testApi()
			testApi(input)
			//fmt.Fprintln(os.Stderr, err)
		}
	}
}
