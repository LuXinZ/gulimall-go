// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsProductAttrValue = "pms_product_attr_value"

// PmsProductAttrValue mapped from table <pms_product_attr_value>
type PmsProductAttrValue struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`
	SpuID     int64  `gorm:"column:spu_id;comment:商品id" json:"spu_id"`
	AttrID    int64  `gorm:"column:attr_id;comment:属性id" json:"attr_id"`
	AttrName  string `gorm:"column:attr_name;comment:属性名" json:"attr_name"`
	AttrValue string `gorm:"column:attr_value;comment:属性值" json:"attr_value"`
	AttrSort  int32  `gorm:"column:attr_sort;comment:顺序" json:"attr_sort"`
	QuickShow int32  `gorm:"column:quick_show;comment:快速展示【是否展示在介绍上；0-否 1-是】" json:"quick_show"`
}

// TableName PmsProductAttrValue's table name
func (*PmsProductAttrValue) TableName() string {
	return TableNamePmsProductAttrValue
}
