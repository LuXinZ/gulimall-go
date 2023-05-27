package start

import (
	"github.com/gin-gonic/gin"
	"gulimall-product/internal/gulimall_product/controller"
)

// installRouters 安装 接口路由.
func InitRouters(g *gin.Engine) error {

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	categoryController := category.New()
	{
		// 创建 users 路由分组
		product := v1.Group("/product")
		{
			categoryRouter := product.Group("/category")
			{
				categoryRouter.GET("/list/tree", categoryController.List)
			}
		}
	}

	return nil
}
