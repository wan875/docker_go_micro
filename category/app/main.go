package main

import(
	"gin-blog-category/router"
    orm "gin-blog-category/database"
)

func main() {
    defer orm.Db.Close()
    router := router.InitRouter()
    router.Run(":8000")
}