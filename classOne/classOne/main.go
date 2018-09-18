package main

import (
	_ "classOne/routers"//_ 作用就是调用后面包里面 的init函数
	"github.com/astaxie/beego"
	//_ "classOne/models"
	"strconv"
)

func main() {
	//beego.SetStaticPath()
	beego.AddFuncMap("ShowPrePage",HandlePrePage)
	beego.AddFuncMap("ShowNextPage",HandleNextPage)
	beego.Run()
}


func HandlePrePage(data int)(string){

	pageIndex := data - 1

	pageIndex1 := strconv.Itoa(pageIndex)
	return pageIndex1
}

func HandleNextPage(data int)(int){
	pageIndex := data + 1
	return pageIndex
}
