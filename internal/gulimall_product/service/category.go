package service

import (
	"gulimall-product/internal/gulimall_product/dao"
	"gulimall-product/internal/gulimall_product/model"
	"sort"
)

type Category struct {
}
type categoryEntity struct {
	model.PmsCategory
	Children []*categoryEntity `json:"children"`
}
type ByID []*categoryEntity

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].Sort < a[j].Sort }
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
	sort.Sort(ByID(menus))

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
	sort.Sort(ByID(children))
	return children
}
func (c *Category) RemoveMenuByIds(catIds []int64) {
	_, err := dao.PmsCategory.Where(dao.PmsCategory.CatID.In(catIds...)).Update(dao.PmsCategory.ShowStatus, 0)
	if err != nil {
		return
	}
}
func (c *Category) Save(category *model.PmsCategory) {
	err := dao.PmsCategory.Create(category)
	if err != nil {
		return
	}
}
func (c *Category) GetById(catId int64) (*model.PmsCategory, error) {
	category, err := dao.PmsCategory.Where(dao.PmsCategory.CatID.Eq(catId)).First()
	if err != nil {
		return nil, err
	}
	return category, nil
}

// 修改 级联更新所有数据
func (c *Category) UpdateCascade(category *model.PmsCategory) {
	dao.Q.Transaction(func(tx *dao.Query) error {

		err := c.updateById(category)
		if err != nil {
			return err
		}
		if category.Name != "" {
			service := CategoryBrandRelationService{}
			err = service.UpdateCategory(category)
			if err != nil {
				return err
			}

		}
		return nil
	})

}
func (c *Category) updateById(category *model.PmsCategory) error {
	_, err := dao.PmsCategory.Where(dao.PmsCategory.CatID.Eq(category.CatID)).Updates(category)
	if err != nil {
		return err
	}
	return nil
}
func (c *Category) UpdateBatchById(category []*model.PmsCategory) {
	var categoryIds []int64
	for _, pmsCategory := range category {
		categoryIds = append(categoryIds, pmsCategory.CatID)
	}
	_, err := dao.PmsCategory.Where(dao.PmsCategory.CatID.In(categoryIds...)).Updates(category)
	if err != nil {
		return
	}
}
