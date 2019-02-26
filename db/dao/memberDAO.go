package dao

import (
	"test_echo_crud/db/model"
	"test_echo_crud/db/util"
)

func MigrateMembers(){
	db := util.GormConnect()
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&model.Member{})

}

func GetMemberById(id int) model.Member{
	db := util.GormConnect()
	defer db.Close()
	member := model.Member{}
	db.First(&member,id)
	return member
}

func GetMembers() []model.Member{
	db := util.GormConnect()
	defer db.Close()
	members := make([]model.Member,0)
	db.Find(&members)
	return members
}

func CreateMember() {
	db := util.GormConnect()
	defer db.Close()
	//INSERT文
	//// 構造体のインスタンス化
	MemberEx := model.Member{}
	//// 挿入したい情報を構造体に与える
	MemberEx.Name = "Takashi"
	//INSERTを実行
	db.Create(&MemberEx)
}

//func UpdateMember() {
//
//}
//
//func DeleteMember()  {
//
//}
//
//func PutMembers(da *gorm.DB) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		//// INSERT文
//		//// 構造体のインスタンス化
//		//MemberEx := db.Member{}
//		//// 挿入したい情報を構造体に与える
//		//MemberEx.Name = "Takashi"
//		//// INSERTを実行
//		//da.Create(&MemberEx)
//
//		members := []model.Member{}
//		data := da.Find(&members)
//		return c.JSON(http.StatusOK, data)
//	}
//}
//
//func DeleteMembers(da *gorm.DB) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		//// INSERT文
//		//// 構造体のインスタンス化
//		//MemberEx := db.Member{}
//		//// 挿入したい情報を構造体に与える
//		//MemberEx.Name = "Takashi"
//		//// INSERTを実行
//		//da.Create(&MemberEx)
//
//		members := []model.Member{}
//		data := da.Find(&members)
//		return c.JSON(http.StatusOK, data)
//	}
//}


