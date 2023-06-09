// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsSkuSaleAttrValue = "pms_sku_sale_attr_value"

// PmsSkuSaleAttrValue mapped from table <pms_sku_sale_attr_value>
type PmsSkuSaleAttrValue struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`
	SkuID     int64  `gorm:"column:sku_id;comment:sku_id" json:"sku_id"`
	AttrID    int64  `gorm:"column:attr_id;comment:attr_id" json:"attr_id"`
	AttrName  string `gorm:"column:attr_name;comment:销售属性名" json:"attr_name"`
	AttrValue string `gorm:"column:attr_value;comment:销售属性值" json:"attr_value"`
	AttrSort  int32  `gorm:"column:attr_sort;comment:顺序" json:"attr_sort"`
}

// TableName PmsSkuSaleAttrValue's table name
func (*PmsSkuSaleAttrValue) TableName() string {
	return TableNamePmsSkuSaleAttrValue
}
