// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsSkuImage = "pms_sku_images"

// PmsSkuImage mapped from table <pms_sku_images>
type PmsSkuImage struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`
	SkuID      int64  `gorm:"column:sku_id;comment:sku_id" json:"sku_id"`
	ImgURL     string `gorm:"column:img_url;comment:图片地址" json:"img_url"`
	ImgSort    int32  `gorm:"column:img_sort;comment:排序" json:"img_sort"`
	DefaultImg int32  `gorm:"column:default_img;comment:默认图[0 - 不是默认图，1 - 是默认图]" json:"default_img"`
}

// TableName PmsSkuImage's table name
func (*PmsSkuImage) TableName() string {
	return TableNamePmsSkuImage
}
