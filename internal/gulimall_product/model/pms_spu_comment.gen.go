// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePmsSpuComment = "pms_spu_comment"

// PmsSpuComment mapped from table <pms_spu_comment>
type PmsSpuComment struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`
	SkuID          int64     `gorm:"column:sku_id;comment:sku_id" json:"sku_id"`
	SpuID          int64     `gorm:"column:spu_id;comment:spu_id" json:"spu_id"`
	SpuName        string    `gorm:"column:spu_name;comment:商品名字" json:"spu_name"`
	MemberNickName string    `gorm:"column:member_nick_name;comment:会员昵称" json:"member_nick_name"`
	Star           bool      `gorm:"column:star;comment:星级" json:"star"`
	MemberIP       string    `gorm:"column:member_ip;comment:会员ip" json:"member_ip"`
	CreateTime     time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	ShowStatus     bool      `gorm:"column:show_status;comment:显示状态[0-不显示，1-显示]" json:"show_status"`
	SpuAttributes  string    `gorm:"column:spu_attributes;comment:购买时属性组合" json:"spu_attributes"`
	LikesCount     int32     `gorm:"column:likes_count;comment:点赞数" json:"likes_count"`
	ReplyCount     int32     `gorm:"column:reply_count;comment:回复数" json:"reply_count"`
	Resources      string    `gorm:"column:resources;comment:评论图片/视频[json数据；[{type:文件类型,url:资源路径}]]" json:"resources"`
	Content        string    `gorm:"column:content;comment:内容" json:"content"`
	MemberIcon     string    `gorm:"column:member_icon;comment:用户头像" json:"member_icon"`
	CommentType    int32     `gorm:"column:comment_type;comment:评论类型[0 - 对商品的直接评论，1 - 对评论的回复]" json:"comment_type"`
}

// TableName PmsSpuComment's table name
func (*PmsSpuComment) TableName() string {
	return TableNamePmsSpuComment
}
