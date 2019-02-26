package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"test_echo_crud/db/dao"
	"test_echo_crud/db/model"
)


// HandleIndexGet は Index のGet時のHTMLデータ生成処理を行います。
func HandleGetPosts(c echo.Context) error {
	//member := dao.GetMember()
	data := dao.GetPostsAll()
	//fmt.Println(dao.GetMember())
	// rows := dao.GetMembers()
	//dao.MigratePosts()
	//dao.CreatePost()
	//var v = model.Member{Name:"taro"}
	//rows = append(rows, v)
	//for _,ve := range rows{
	//	fmt.Println(ve.Name)
	//}

	//got := []model.Member{}
	//json.Unmarshal(data.,got)
	return c.Render(http.StatusOK, "posts_all", data)
}


// HandleIndexGet は Index のGet時のHTMLデータ生成処理を行います。
func HandlePostPost(c echo.Context) error {

	memberId,_ := strconv.Atoi(c.FormValue("memberId"))
	member := dao.GetMemberById(memberId)
	post := model.Post{
		Member:member,
		MemberID:memberId,
		Name:c.FormValue("name"),
		Content:c.FormValue("content"),
	}
	dao.CreatePost(post)
	data := dao.GetPostsAll()
	return c.Render(http.StatusOK, "posts_all", data)
}

// HandleHelloGet は /hello のGet時のHTMLデータ生成処理を行います。
func HandleEditPost(c echo.Context) error {
	postId,_ := strconv.Atoi(c.FormValue("postId"))
	post := dao.GetPostById(postId)
	return c.Render(http.StatusOK, "update_post_form", post)

}

// HandleHelloGet は /hello のGet時のHTMLデータ生成処理を行います。
func HandleUpdatePost(c echo.Context) error {
	postId,_ := strconv.Atoi(c.FormValue("postId"))
	post := dao.GetPostById(postId)
	post.Name = c.FormValue("name")
	post.Content = c.FormValue("content")
	dao.UpdatePost(post)
	data := dao.GetPostsAll()
	return c.Render(http.StatusOK, "posts_all", data)
}


// HandleHelloPost は /hello のPost時のHTMLデータ生成処理を行います。
func HandleDeletePost(c echo.Context) error {
	postId,_ := strconv.Atoi(c.FormValue("id"))
	post := dao.GetPostById(postId)
	dao.DeletePost(post)
	data := dao.GetPostsAll()
	return c.Render(http.StatusOK, "posts_all", data)
}
