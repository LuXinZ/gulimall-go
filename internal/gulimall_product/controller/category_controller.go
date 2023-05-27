package category

import (
	"github.com/gin-gonic/gin"
	"gulimall-product/internal/gulimall_product/service"
	"gulimall-product/internal/pkg/core"
)

type CategoryController struct {
	s service.Category
}

func New() *CategoryController {
	return &CategoryController{}
}
func (ct *CategoryController) List(c *gin.Context) {
	list := ct.s.ListWithTree()
	core.WriteResponse(c, nil, list)
}
