// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsSpuInfoDesc = "pms_spu_info_desc"

// PmsSpuInfoDesc mapped from table <pms_spu_info_desc>
type PmsSpuInfoDesc struct {
	SpuID   int64  `gorm:"column:spu_id;primaryKey;comment:商品id" json:"spu_id"`
	Decript string `gorm:"column:decript;comment:商品介绍" json:"decript"`
}

// TableName PmsSpuInfoDesc's table name
func (*PmsSpuInfoDesc) TableName() string {
	return TableNamePmsSpuInfoDesc
}
