package category

import (
	"github.com/gin-gonic/gin"
	"gulimall-product/internal/gulimall_product/model"
	"gulimall-product/internal/gulimall_product/service"
	"gulimall-product/internal/pkg/core"
	"gulimall-product/internal/pkg/errno"
	"strconv"
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
func (ct *CategoryController) Delete(c *gin.Context) {
	var ids []int64
	if err := c.ShouldBindJSON(&ids); err != nil {
		core.WriteResponse(c, errno.ErrBind, "fail")
		return
	}
	ct.s.RemoveMenuByIds(ids)
	core.WriteResponse(c, nil, "success")
}
func (ct *CategoryController) Save(c *gin.Context) {
	var category model.PmsCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return
	}
	ct.s.Save(&category)
	core.WriteResponse(c, nil, "success")
}
func (ct *CategoryController) Info(c *gin.Context) {
	catId, err := strconv.ParseInt(c.Param("catId"), 10, 64)

	if err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter, nil)
		return
	}
	category, err := ct.s.GetById(catId)
	if err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter, nil)
		return
	}
	core.WriteResponse(c, nil, category)
}
func (ct *CategoryController) Update(c *gin.Context) {
	var category *model.PmsCategory
	if err := c.ShouldBindJSON(category); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	ct.s.UpdateCascade(category)
	core.WriteResponse(c, nil, nil)
}

// 批量修改节点
func (ct *CategoryController) UpdateNodes(c *gin.Context) {
	var categorys []*model.PmsCategory
	if err := c.ShouldBindJSON(categorys); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	ct.s.UpdateBatchById(categorys)
	core.WriteResponse(c, nil, nil)
}
