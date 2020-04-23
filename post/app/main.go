package main

import(
	"gin-blog-post/router"
    orm "gin-blog-post/database"
)

func main() {
    defer orm.Db.Close()
    router := router.InitRouter()
    router.Run(":8000")
}