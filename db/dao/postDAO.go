package dao

import (
	"test_echo_crud/db/model"
	"test_echo_crud/db/util"
)


type PostMember struct{
	model.Post
	MemberName string
}



func MigratePosts(){
	db := util.GormConnect()
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&model.Post{})

}

func GetPostById(id int) model.Post{
	db := util.GormConnect()
	defer db.Close()
	post := model.Post{}
	db.First(&post,id)
	return post
}

func GetPostsAll() []PostMember{
	db := util.GormConnect()
	defer db.Close()
	posts := make([]model.Post,0)
	db.Find(&posts)
	rows := []PostMember{}
	for _,v := range posts{

		row := PostMember{
			Post:       v,
			MemberName: GetMemberById(v.MemberID).Name,
		}
		rows = append(rows,row)
	}
	//db.Find(&posts)
	return rows
}

func GetPostsByMemberId(member model.Member) []model.Post{
	db := util.GormConnect()
	defer db.Close()
	posts := make([]model.Post,0)
	db.Where("member_id = ?", member.ID).Find(&posts)
	//db.Find(&posts)
	return posts
}

func CreatePost(post model.Post) {
	db := util.GormConnect()
	defer db.Close()
	//INSERT文
	//// 構造体のインスタンス化
	PostEx := model.Post{}
	//// 挿入したい情報を構造体に与える
	PostEx.Member = GetMemberById(post.MemberID)
	PostEx.Name = post.Name
	PostEx.Content = post.Content
	//INSERTを実行
	db.Create(&PostEx)
}

func UpdatePost(post model.Post) {
	db := util.GormConnect()
	defer db.Close()
	db.Save(&post)
}

func DeletePost(post model.Post)  {
	db := util.GormConnect()
	defer db.Close()
	db.Delete(&post)
	}