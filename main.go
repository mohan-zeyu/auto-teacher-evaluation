package main

import (
	"fmt"

	"github.com/mohan-zeyu/golang-login-zju/util"
	"github.com/mohan-zeyu/golang-login-zju/zju"
)

func main() {
	var username, password string
	fmt.Println("请输入学号: ")
	fmt.Scanf("%s", &username)
	fmt.Println("请输入密码: ")
	fmt.Scanf("%s", &password)

	z := zju.NewZJUAM(username, password, zju.WithRedirectsDisabled())
	ctx, cancel	:= util.CreateCtx()
	defer cancel()
	z.Judge(ctx)
}
