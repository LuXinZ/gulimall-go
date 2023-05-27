package service

import (
	"gulimall-product/internal/gulimall_product/dao"
	"gulimall-product/internal/gulimall_product/model"
)

type Category struct {
}
type categoryEntity struct {
	model.PmsCategory
	Children []*categoryEntity `json:"children"`
}

func (c *Category) ListWithTree() interface{} {
	// 查找出所有的分类
	pmsCategories, err := dao.PmsCategory.Find()
	if err != nil {
		return nil
	}
	// 组装成父子的属性结构
	// 找出所有一级分类
	var categories []*categoryEntity
	var menus []*categoryEntity
	for _, pmsCategory := range pmsCategories {
		category := &categoryEntity{
			PmsCategory: *pmsCategory,
			Children:    make([]*categoryEntity, 0),
		}
		if pmsCategory.ParentCid == 0 {
			menus = append(categories, category)
		}

		categories = append(categories, category)
	}
	for _, menu := range menus {
		m := append(menu.Children, getChildren(menu, categories)...)
		menu.Children = m
	}
	return menus
}
func getChildren(root *categoryEntity, all []*categoryEntity) []*categoryEntity {
	var children []*categoryEntity
	for _, category := range all {
		if category.ParentCid == root.CatID {
			category.Children = getChildren(category, all)
			children = append(children, category)
		}
	}
	return children
}
