package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mohan-zeyu/golang-login-zju/util"
	"github.com/mohan-zeyu/golang-login-zju/zju"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入学号: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Println("请输入密码: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	z := zju.NewZJUAM(username, password, zju.WithRedirectsDisabled())
	ctx, cancel	:= util.CreateCtx()
	defer cancel()
	z.Judge(ctx)
}
