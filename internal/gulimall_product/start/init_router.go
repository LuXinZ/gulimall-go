package start

import (
	"github.com/gin-gonic/gin"
	"gulimall-product/internal/gulimall_product/controller"
)

// installRouters 安装 接口路由.
func InitRouters(g *gin.Engine) error {

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	c := category.New()
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/list")
		{
			userv1.GET("", c.List)
		}
	}

	return nil
}
