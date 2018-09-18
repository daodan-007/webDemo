package controllers
//123
import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"github.com/astaxie/beego/orm"
	"classOne/models"
	"strconv"
	"math"
)

type ArticleController struct {
	beego.Controller
}
//处理下拉框改变法的请求(post请求)
func(this*ArticleController)HandleSelect(){
	//1.接受数据
	typeName:=this.GetString("select")
	//beego.Info(typeName)
	//2.处理数据
	if typeName == ""{
		beego.Info("下拉框传递数据失败")
		return
	}
	//3.查询数据
	o := orm.NewOrm()
	var articles[]models.Article
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articles)
	//beego.Info(articles)
}

//文章列表页
func (this*ArticleController)ShowArticleList(){
	//1.查询

		//1.有一个orm对象
		o := orm.NewOrm()
		qs :=o.QueryTable("Article")
		var articles[]models.Article
		//qs.All(&articles)  //select * from Article
		//pageIndex1 := 1
		pageIndex := this.GetString("pageIndex")
		pageIndex1,err:=strconv.Atoi(pageIndex)
		if err !=nil{
			pageIndex1 = 1//自己
		}



		//qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articles)//1.pagesize  一页显示多少   2.start起始位置



		if err != nil{
			beego.Info("查询错误")
			return
		}

		FirstPage := false//标识是否是首页

		EndPage := false //标识是否是末页
		//首页末页数据处理
		if pageIndex1 == 1{
			FirstPage = true
		}
		if pageIndex1 == int(pageCount){
			EndPage = true
		}
		//获取类型数据


		var types []models.ArticleType
		o.QueryTable("ArticleType").All(&types)
		this.Data["types"] = types

		//根据类型获取数据

	//1.接受数据
	typeName:=this.GetString("typeName")
	//beego.Info(typeName)
	//2.处理数据
	pageSize := 2
	var count int64
	var pageCount1 int64
	var articleswithtype []models.Article
	if typeName == ""{
		count ,_= qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).Count()//返回数据条目数   加过滤器
		pageCount := float64(count)/float64(pageSize)  //求总页数
		pageCount1=math.Ceil(pageCount)  //天花板函数获取正确总页码
		//获取总页数
		pageSize := 1//设置每页显示数

		start:=pageSize*(pageIndex1 -1)//位置
		beego.Info("下拉框传递数据失败")
		qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articleswithtype)
	}else {
		qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articleswithtype)
	}
	//3.查询数据

		userName:=this.GetSession("userName")
		this.Data["userName"] = userName


		this.Data["typeName"] = typeName
		beego.Info("count=",count)
		this.Data["EndPage"] = EndPage
		this.Data["FirstPage"] = FirstPage
		this.Data["count"] = count
		this.Data["pageCount"] = pageCount1
		this.Data["pageIndex"] = pageIndex1

		this.Data["articles"] = articleswithtype


	//2.把数据传递给视图显示
	this.Layout = "layout.html"
	this.TplName = "index.html"
}

//显示文章界面
func(this*ArticleController)ShowAddArticle(){
	//查询类型数据，传递到视图中
	o:=orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	this.Data["types"] = types

	this.TplName = "add.html"
}
//处理增加文章业务
/*
1.那数据
2.判断数据
3.插入数据
4.返回试图
 */
func (this*ArticleController)HandleAddArtcile(){
	//1.那数据
	//那标题
	artiName:= this.GetString("articleName")
	artiContent := this.GetString("content")
	f,h,err:=this.GetFile("uploadname")

	defer f.Close()
	//上传文件处理
	//1.判断文件格式
	ext := path.Ext(h.Filename)
	if ext != ".jpg" && ext != ".png"&&ext != ".jpeg"{
		beego.Info("上传文件格式不正确")
		return
	}

	//2.文件大小
	if h.Size>5000000{
		beego.Info("文件太大，不允许上传")
		return
	}

	//3.不能重名
	fileName := time.Now().Format("2006-01-02 15:04:05")


	err2:=this.SaveToFile("uploadname","./static/img/"+fileName+ext)
	if err != nil{
		beego.Info("上传文件失败")
		return
	}

	if err != nil{
		beego.Info("上传文件失败",err2)
		return
	}




	//3.插入数据
		//1.获取orm对象
		o := orm.NewOrm()
		//2.创建一个插入对象
		article := models.Article{}
		//3.赋值
		article.Title = artiName
		article.Content = artiContent
		article.Img = "/static/img/"+fileName+ext


		//4.返回试图


		//给article对象复制
		//获取到下拉框传递过来的类型数据
		typeName:=this.GetString("select")
		//类型判断
		if typeName == ""{
			beego.Info("下拉匡数据错误")
			return
		}
		//获取type对象
		var artiType models.ArticleType
		artiType.TypeName = typeName
		err=o.Read(&artiType,"TypeName")
		if err != nil{
			beego.Info("获取类型错误")
			return
		}
		article.ArticleType = &artiType


	//4.插入
	_,err = o.Insert(&article)
	if err != nil{
		beego.Info("插入数据失败")
		return
	}
		this.Redirect("/Article/ShowArticle",302)

}



//显示文章详情
func (this*ArticleController)ShowContent(){
	//1.获取Id
	id:=this.GetString("id")
	beego.Info(id)
	//2.查询数据
		//1.获取orm对象
		o:=orm.NewOrm()
		//2.获取查询对象
		id2,_:=strconv.Atoi(id)
		article := models.Article{Id2:id2}
		//3.查询
		err := o.Read(&article)
		if err != nil{
			beego.Info("查询数据为空")
			return
		}
		article.Count+=1
		//多对多插入读者

		//1.获取操作对象
		//artile := models.Article{Id2:id2}
		//2.获取多对多操作对象
		m2m := o.QueryM2M(&article,"Users")
		//3.获取插入对象
		userName:=this.GetSession("userName")
		user := models.User{}
		user.UserName = userName.(string)
		o.Read(&user,"UserName")
		//4.多对多插入
		_,err=m2m.Add(&user)
		if err != nil{
			beego.Info("插入失败")
			return
		}
		o.Update(&article) //没有指定更新哪一列，他会自己查

		//o.LoadRelated(&article,"Users")
		//多对多查询的时候会出错

		//o.QueryTable("Article").Filter("Id2",id2).Filter("Users__User__Id",user.Id).Distinct().One(&article)
		var users[]models.User
		beego.Info(article)
		o.QueryTable("User").Filter("Articles__Article__Id2",id2).Distinct().All(&users)
	//3.传递数据给视图
		beego.Info(article)
		this.Data["users"] = users
		this.Data["article"] = article
		this.Layout = "layout.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["contentHead"] = "head.html"
		this.TplName = "content.html"
}


//1.URLchuanzhi
//2.执行delete操作


//删除文章
func (this*ArticleController)HandleDelete(){
	id,_:=this.GetInt("id")
	//1.orm对象
	o := orm.NewOrm()

	//要有删除对象
	article := models.Article{Id2:id}

	//3.删除
	o.Delete(&article)

	this.Redirect("/Article/ShowArticle",302)
}

//显示更新页面
func (this*ArticleController)ShowUpdate(){
	//获取数据
	id := this.GetString("id")
	//判断
	if id == ""{
		beego.Info("连接错误")
		return
	}
	//查询操作
	o := orm.NewOrm()
	article := models.Article{}
	//类型转换
	id2,_:=strconv.Atoi(id)
	article.Id2 = id2

	err:=o.Read(&article)
	if err !=nil{
		beego.Info("查询错误")
		return
	}

	//把数据传递给视图
	this.Data["article"] = article
	this.TplName = "update.html"


}

//处理更新数据
func(this*ArticleController)HandleUpdate(){
	//1.拿数据
	name:=this.GetString("articleName")
	content := this.GetString("content")
	id,_:=this.GetInt("id")

	//问题一 id是不是没有传过来
	//2.判断数据
	if name == "" || content == ""{
		beego.Info("更新数据失败")
		return
	}
	f,h,err:=this.GetFile("uploadname")
	if err != nil{
		beego.Info("上传文件失败")
		return
	}
	defer f.Close()
	//1.判断大小
	if h.Size > 500000{
		beego.Info("图片太大")
		return
	}
	//2.判断类型
	ext:=path.Ext(h.Filename)
	if ext != ".jpg"&&ext!=".png"&&ext!=".jpeg"{
		beego.Info("上传文件类型错误")
		return
	}
	//3.防止文件名重复
	filename:=time.Now().Format("2006-01-02-15:04:05")
	this.SaveToFile("uploadname","./static/img/"+filename+ext)


	//更新操作
	o:=orm.NewOrm()
	article := models.Article{Id2:id}
	//读取操作
	err = o.Read(&article)
	if err != nil{
		beego.Info("要更新的文章不存在")
		return
	}
	//更新
	article.Title = name
	article.Content = content
	article.Img = "./static/img/"+filename+ext
	_,err=o.Update(&article)
	if err != nil{
		beego.Info("更新失败")
		return
	}

	//跳转
	this.Redirect("/Article/ShowArticle",302)

}


func (this*ArticleController)ShowAddType(){
	//1.读取类型表，显示数据
	o := orm.NewOrm()
	var artiTypes[]models.ArticleType
	//查询
	_,err:=o.QueryTable("ArticleType").All(&artiTypes)
	if err != nil{
		beego.Info("查询类型错误")
	}
	this.Data["title"] = "<title>添加类型</title>"
	this.Data["types"] = artiTypes
	this.Layout = "layout.html"
	this.TplName = "addType.html"
}
//处理添加类型业务
func (this*ArticleController)HandleAddType(){
	//1.获取数据
	typename:=this.GetString("typeName")
	//2.判断数据
	if typename == ""{
		beego.Info("添加类型数据为空")
		return
	}
	//3.执行插入操作
	o := orm.NewOrm()
	var artiType models.ArticleType
	artiType.TypeName = typename
	_,err:=o.Insert(&artiType)
	if err != nil{
		beego.Info("插入失败")
		return
	}
	//4.展示视图？
	this.Redirect("/Article/AddArticleType",302)
}

//退出登陆
func (this*ArticleController)Logout(){
	//1.删除登陆状态
	this.DelSession("userName")
	//2.跳转登陆页面
	this.Redirect("/",302)
}
//删除文章类型
func (this*ArticleController)DeleteType(){
	//1.获取类型Id
	id:=this.GetString("id")
	id2 ,_:= strconv.Atoi(id)
	//2.都要进行数据判断
	if id2 == 0{
		beego.Info("获取id错误")
		return
	}



	//3.删除操作
	o := orm.NewOrm()
	artiType := models.ArticleType{Id:id2}
	o.Delete(&artiType)

	//4.返回视图
	this.Redirect("/Article/AddArticleType",302)
}















































