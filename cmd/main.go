package main

import (
	"fmt"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/providers"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

func main() {
	pass, _ := utils.HashPassword("123123123")
	isOk, _ := utils.ComparePassword(pass, "123123123")
	fmt.Println(isOk)
	if err := providers.InitServer(); err != nil {
		fmt.Println(err)
	}
}
