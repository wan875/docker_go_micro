package router

import (
    "github.com/gin-gonic/gin"
    . "gin-blog-category/controller"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/categorys", Categorys)

    router.POST("/category", Create)

    //router.PUT("/post/:id", Update)

    //router.DELETE("/post/:id", Destroy)

    return router
}