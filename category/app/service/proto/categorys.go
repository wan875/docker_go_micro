package proto
 
import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	//"time"
	"errors"
	model "gin-blog-category/models"
	"strconv"
)
 
type CategorysServer struct {
}
 
func (s *CategorysServer) GetCategorys(ctx context.Context, request *CategoryRequest, resp *CategoryList) error {
	var category model.Category
	result, err := category.Categorys()
	if err != nil {
		return errors.New("error")
	}
	
	array := []*Category{}
	for i:=0;i<len(result);i++ {
		array = append(array,&Category{Id:strconv.Itoa(int(result[i].Id)),Name:result[i].Name,Slug:result[i].Slug,OrderNum:result[i].OrderNum})
	}
	
	resp.Ele = array
	return nil
}
 
func RunService() {
	reg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"consul:8500",}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("getcategorys"), //服务名称
	)
	service.Init()
	RegisterCategoryServiceHandler(service.Server(), new(CategorysServer))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
 
}
