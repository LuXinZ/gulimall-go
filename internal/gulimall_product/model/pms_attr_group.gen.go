// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsAttrGroup = "pms_attr_group"

// PmsAttrGroup mapped from table <pms_attr_group>
type PmsAttrGroup struct {
	AttrGroupID   int64  `gorm:"column:attr_group_id;primaryKey;autoIncrement:true;comment:分组id" json:"attr_group_id"`
	AttrGroupName string `gorm:"column:attr_group_name;comment:组名" json:"attr_group_name"`
	Sort          int32  `gorm:"column:sort;comment:排序" json:"sort"`
	Descript      string `gorm:"column:descript;comment:描述" json:"descript"`
	Icon          string `gorm:"column:icon;comment:组图标" json:"icon"`
	CatelogID     int64  `gorm:"column:catelog_id;comment:所属分类id" json:"catelog_id"`
}

// TableName PmsAttrGroup's table name
func (*PmsAttrGroup) TableName() string {
	return TableNamePmsAttrGroup
}
