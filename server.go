package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"test_echo_crud/const"
	"test_echo_crud/handler"
)

// レイアウト適用済のテンプレートを保存するmap
var templates map[string]*template.Template

// Template はHTMLテンプレートを利用するためのRenderer Interfaceです。
type Template struct {
	templates *template.Template
}

// Render はHTMLテンプレートにデータを埋め込んだ結果をWriterに書き込みます。
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return getTemplate(name).ExecuteTemplate(w, "layout.html", data)
}

// 初期化を行います。
func init() {
	//loadTemplates()
	//dao.MigrateMembers()
	//dao.CreateMember()
	//dao.MigratePosts()
}
//
//// 各HTMLテンプレートに共通レイアウトを適用した結果を保存します（初期化時に実行）。
//func loadTemplates() {
//	templates = make(map[string]*template.Template)
//	templates["update_post_form"] = template.Must(
//		template.ParseFiles(_const.BaseTemplate, _const.UpdataPost))
//	templates["posts_all"] = template.Must(
//		template.ParseFiles(_const.BaseTemplate, _const.PostsAll))
//}

func getTemplate(name string)(*template.Template){
	switch name {
	case "posts_all":
		return template.Must(template.ParseFiles(_const.BaseTemplate, _const.PostsAll))
	case "update_post_form":
		return template.Must(template.ParseFiles(_const.BaseTemplate, _const.UpdataPost))
	default:
		return nil
	}
}

func main() {	e := echo.New()
	//e.Pre(MethodOverride)

	// テンプレートを利用するためのRendererの設定
	t := &Template{
		templates: template.Must(template.ParseGlob(_const.BaseTemplate)),
	}
	e.Renderer = t


	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.HandleGetPosts)
	e.POST("/posts", handler.HandlePostPost)
	e.POST("/post/edit", handler.HandleEditPost)
	e.POST("/post/update", handler.HandleUpdatePost)
	e.POST("/post/delete", handler.HandleDeletePost)
	//e.PUT("/post/:id", handler.HandleUpdatePost)
	//e.DELETE("/post/:id", handler.HandleDeletePost)
	e.Logger.Fatal(e.Start(":3000"))

}






//{
//		// member
//		api.GET("/members", dao.GetMembers(da))
//		api.POST("/members", handler.PostMembers())
//		api.PUT("/members", handler.PutMembers())
//		api.DELETE("/members", handler.DeleteMembers())
//}


//e := controller.Init(db)
//e.Logger.Fatal(e.Start(":1323"))
//http.HandleFunc("/members", handler.HtmlHandler0)
// サーバーを起動
//http.ListenAndServe(":8989", nil)

//api.GET("/members", func(c echo.Context) error {
//	return c.String(http.StatusOK, "members GET")
//})
////func main() {
////	e := controller.Init()
////	e.Logger.Fatal(e.Start(":1323"))
////}
//
////　hello,world!
//
////func main() {
////	e := echo.New()
////	e.GET("/", func(c echo.Context) error {
////		return c.String(http.StatusOK, "Hello, World!")
////	})
////	e.Logger.Fatal(e.Start(":6060"))
////}
//
////  get post put delete区別
//
////func main() {
////	e := echo.New()
////
//// /api/members でアクセス
//api := e.Group("/api")
//{
//	api.GET("/members", func(c echo.Context) error {
//		return c.String(http.StatusOK, "members GET")
//	})
//	api.POST("/members", func(c echo.Context) error {
//		return c.String(http.StatusOK, "members POST")
//	})
//	api.PUT("/members", func(c echo.Context) error {
//		return c.String(http.StatusOK, "members PUT")
//	})
//	api.DELETE("/members", func(c echo.Context) error {
//		return c.String(http.StatusOK, "members POST")
//	})
//}
//e.Logger.Fatal(e.Start(":1323"))
