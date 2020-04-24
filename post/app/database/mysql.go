package database

import(
	_ "github.com/go-sql-driver/mysql" //加载mysql
    "github.com/jinzhu/gorm"
    "fmt"
)

var Db *gorm.DB

func init(){
	var err error
	Db, err = gorm.Open("mysql", "root:root@tcp(mysql:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }

    if Db.Error != nil {
        fmt.Printf("database error %v", Db.Error)
    }
}
