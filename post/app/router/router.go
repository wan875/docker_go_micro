package router

import (
    "github.com/gin-gonic/gin"
    . "gin-blog-post/controller"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/posts", Posts)

    router.POST("/post", Create)

    //router.PUT("/post/:id", Update)

    //router.DELETE("/post/:id", Destroy)

    return router
}