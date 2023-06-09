// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsAttr = "pms_attr"

// PmsAttr mapped from table <pms_attr>
type PmsAttr struct {
	AttrID      int64  `gorm:"column:attr_id;primaryKey;autoIncrement:true;comment:属性id" json:"attr_id"`
	AttrName    string `gorm:"column:attr_name;comment:属性名" json:"attr_name"`
	SearchType  int32  `gorm:"column:search_type;comment:是否需要检索[0-不需要，1-需要]" json:"search_type"`
	ValueType   int32  `gorm:"column:value_type;comment:值类型[0-为单个值，1-可以选择多个值]" json:"value_type"`
	Icon        string `gorm:"column:icon;comment:属性图标" json:"icon"`
	ValueSelect string `gorm:"column:value_select;comment:可选值列表[用逗号分隔]" json:"value_select"`
	AttrType    int32  `gorm:"column:attr_type;comment:属性类型[0-销售属性，1-基本属性，2-既是销售属性又是基本属性]" json:"attr_type"`
	Enable      int64  `gorm:"column:enable;comment:启用状态[0 - 禁用，1 - 启用]" json:"enable"`
	CatelogID   int64  `gorm:"column:catelog_id;comment:所属分类" json:"catelog_id"`
	ShowDesc    int32  `gorm:"column:show_desc;comment:快速展示【是否展示在介绍上；0-否 1-是】，在sku中仍然可以调整" json:"show_desc"`
}

// TableName PmsAttr's table name
func (*PmsAttr) TableName() string {
	return TableNamePmsAttr
}
