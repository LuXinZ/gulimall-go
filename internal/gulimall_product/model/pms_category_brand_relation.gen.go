// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsCategoryBrandRelation = "pms_category_brand_relation"

// PmsCategoryBrandRelation mapped from table <pms_category_brand_relation>
type PmsCategoryBrandRelation struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BrandID     int64  `gorm:"column:brand_id;comment:品牌id" json:"brand_id"`
	CatelogID   int64  `gorm:"column:catelog_id;comment:分类id" json:"catelog_id"`
	BrandName   string `gorm:"column:brand_name" json:"brand_name"`
	CatelogName string `gorm:"column:catelog_name" json:"catelog_name"`
}

// TableName PmsCategoryBrandRelation's table name
func (*PmsCategoryBrandRelation) TableName() string {
	return TableNamePmsCategoryBrandRelation
}
