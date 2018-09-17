package main

import (
	"gin-blog/models"
	"fmt"
	"gin-blog/pkg/setting"
)

func main()  {
	setting.Setup()
	models.Setup()

	fmt.Println("ok")
}

