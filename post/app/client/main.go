package main
 
import (
	ord "gin-blog-post/client/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	
)
 
const ADDRESS = "localhost:50051"
 
func main() {
	reg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{
			"127.0.0.1:8500",
		}
	})
	service := micro.NewService(micro.Registry(reg),micro.Name("order.client"))
	service.Init()
	helloService := ord.NewCategoryService("getcategorys", service.Client())
	orderInfo, err := helloService.GetCategorys(context.TODO(), &ord.CategoryRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orderInfo.Ele)
}