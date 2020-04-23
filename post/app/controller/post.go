package controller

import (
    "github.com/gin-gonic/gin"
    model "gin-blog-post/models"
    "net/http"
    "strconv"
    
    ord "gin-blog-post/client/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

const ADDRESS = "localhost:50051"

func getCategorys() map[string]interface{} {
	reg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"consul:8500",
		}
	})
	service := micro.NewService(micro.Registry(reg),micro.Name("order.client"))
	service.Init()
	helloService := ord.NewCategoryService("getcategorys", service.Client())
	orderInfo, err := helloService.GetCategorys(context.TODO(), &ord.CategoryRequest{})
	if err != nil {
		fmt.Println(err)
	}
	
	catMap := make(map[string]interface{})
	for _, v := range orderInfo.Ele {
		catMap[v.Name] = v
	}
	
	return catMap
}

//添加数据
func Create(c *gin.Context) {
    var post model.Post
    post.PostTitle = c.Request.FormValue("title")
    post.PostExcerpt = c.Request.FormValue("desc")
    post.PostContent = c.Request.FormValue("content")
    post.Tags = c.Request.FormValue("tags")
    post.Categorys = c.Request.FormValue("categorys")
    is_published,_ := strconv.ParseInt(c.Request.FormValue("is_published"),0,0)
    post.IsPublished = is_published
    post.Poster = c.Request.FormValue("poster")
    
    //仅用于测试微服务，此处做分类传递判断
    catMap := getCategorys()
    fmt.Println(catMap)
    if catMap[post.Categorys] == nil {
    	c.JSON(http.StatusOK, gin.H{
            "code":    -1,
            "message": "分类不存在",
        })
        return
    }
    
    id, err := post.Insert()

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
func Posts(c *gin.Context) {
    var post model.Post
    //user.Username = c.Request.FormValue("username")
    //user.Password = c.Request.FormValue("password")
    page,_ := strconv.Atoi(c.Query("page"))
    result, err := post.Posts(int64(page))

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

