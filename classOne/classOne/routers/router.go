package routers

import (
	"classOne/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.InsertFilter("/Article/*",beego.BeforeRouter,FilterFunc)
    beego.Router("/register",&controllers.RegController{},"get:ShowReg;post:HandleReg")
	beego.Router("/",&controllers.LoginController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/Article/ShowArticle",&controllers.ArticleController{},"get:ShowArticleList;post:HandleSelect")
	beego.Router("/Article/AddArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArtcile")
	beego.Router("/Article/ArticleContent",&controllers.ArticleController{},"get:ShowContent")
	beego.Router("/Article/DeleteArticle",&controllers.ArticleController{},"get:HandleDelete")
	//更新功能
	beego.Router("/Article/UpdateArticle",&controllers.ArticleController{},"get:ShowUpdate;post:HandleUpdate")
	//添加类型
	beego.Router("/Article/AddArticleType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")
	//推出登陆
	beego.Router("/Article/Logout",&controllers.ArticleController{},"get:Logout")
	//删除类型
	beego.Router("/Article/deleteType",&controllers.ArticleController{},"get:DeleteType")
}

var FilterFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil{
		ctx.Redirect(302,"/")//如果有输出不再往下执行
	}
}