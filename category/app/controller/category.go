package controller

import (
    "github.com/gin-gonic/gin"
    model "gin-blog-category/models"
    "net/http"
    "strconv"
)

//添加数据
func Create(c *gin.Context) {
    var category model.Category
    category.Name = c.Request.FormValue("name")
    category.Slug = c.Request.FormValue("slug")
    order_num,_ := strconv.ParseInt(c.Request.FormValue("order_num"),0,0)
    category.OrderNum = order_num
    
    id, err := category.Insert()

    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "code":    -1,
            "message": err,
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code":  1,
        "message": "添加成功",
        "data":    id,
    })
}

//列表数据
func Categorys(c *gin.Context) {
    var category model.Category
    //user.Username = c.Request.FormValue("username")
    //user.Password = c.Request.FormValue("password")
    result, err := category.Categorys()

    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "code":    -1,
            "message": "抱歉未找到相关信息",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "data":   result,
    })
}


