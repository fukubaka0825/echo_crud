package util
import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "sT%818129"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME   := "users"
	Parse    := "?parseTime=true"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+Parse
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	// ログ出力を有効にする
	db.LogMode(true)

	return db
}
