package service

import (
	"gulimall-product/internal/gulimall_product/dao"
	"gulimall-product/internal/gulimall_product/model"
)

type CategoryBrandRelationService struct {
}

func (s *CategoryBrandRelationService) UpdateCategory(category *model.PmsCategory) error {
	_, err := dao.PmsCategoryBrandRelation.Where(dao.PmsCategory.CatID.Eq(category.CatID)).Updates(model.PmsCategoryBrandRelation{CatelogName: category.Name, CatelogID: category.CatID})
	if err != nil {
		return err
	}
	return nil
}
