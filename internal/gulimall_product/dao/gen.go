// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                        = new(Query)
	PmsAttr                  *pmsAttr
	PmsAttrAttrgroupRelation *pmsAttrAttrgroupRelation
	PmsAttrGroup             *pmsAttrGroup
	PmsBrand                 *pmsBrand
	PmsCategory              *pmsCategory
	PmsCategoryBrandRelation *pmsCategoryBrandRelation
	PmsCommentReplay         *pmsCommentReplay
	PmsProductAttrValue      *pmsProductAttrValue
	PmsSkuImage              *pmsSkuImage
	PmsSkuInfo               *pmsSkuInfo
	PmsSkuSaleAttrValue      *pmsSkuSaleAttrValue
	PmsSpuComment            *pmsSpuComment
	PmsSpuImage              *pmsSpuImage
	PmsSpuInfo               *pmsSpuInfo
	PmsSpuInfoDesc           *pmsSpuInfoDesc
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	PmsAttr = &Q.PmsAttr
	PmsAttrAttrgroupRelation = &Q.PmsAttrAttrgroupRelation
	PmsAttrGroup = &Q.PmsAttrGroup
	PmsBrand = &Q.PmsBrand
	PmsCategory = &Q.PmsCategory
	PmsCategoryBrandRelation = &Q.PmsCategoryBrandRelation
	PmsCommentReplay = &Q.PmsCommentReplay
	PmsProductAttrValue = &Q.PmsProductAttrValue
	PmsSkuImage = &Q.PmsSkuImage
	PmsSkuInfo = &Q.PmsSkuInfo
	PmsSkuSaleAttrValue = &Q.PmsSkuSaleAttrValue
	PmsSpuComment = &Q.PmsSpuComment
	PmsSpuImage = &Q.PmsSpuImage
	PmsSpuInfo = &Q.PmsSpuInfo
	PmsSpuInfoDesc = &Q.PmsSpuInfoDesc
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                       db,
		PmsAttr:                  newPmsAttr(db, opts...),
		PmsAttrAttrgroupRelation: newPmsAttrAttrgroupRelation(db, opts...),
		PmsAttrGroup:             newPmsAttrGroup(db, opts...),
		PmsBrand:                 newPmsBrand(db, opts...),
		PmsCategory:              newPmsCategory(db, opts...),
		PmsCategoryBrandRelation: newPmsCategoryBrandRelation(db, opts...),
		PmsCommentReplay:         newPmsCommentReplay(db, opts...),
		PmsProductAttrValue:      newPmsProductAttrValue(db, opts...),
		PmsSkuImage:              newPmsSkuImage(db, opts...),
		PmsSkuInfo:               newPmsSkuInfo(db, opts...),
		PmsSkuSaleAttrValue:      newPmsSkuSaleAttrValue(db, opts...),
		PmsSpuComment:            newPmsSpuComment(db, opts...),
		PmsSpuImage:              newPmsSpuImage(db, opts...),
		PmsSpuInfo:               newPmsSpuInfo(db, opts...),
		PmsSpuInfoDesc:           newPmsSpuInfoDesc(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	PmsAttr                  pmsAttr
	PmsAttrAttrgroupRelation pmsAttrAttrgroupRelation
	PmsAttrGroup             pmsAttrGroup
	PmsBrand                 pmsBrand
	PmsCategory              pmsCategory
	PmsCategoryBrandRelation pmsCategoryBrandRelation
	PmsCommentReplay         pmsCommentReplay
	PmsProductAttrValue      pmsProductAttrValue
	PmsSkuImage              pmsSkuImage
	PmsSkuInfo               pmsSkuInfo
	PmsSkuSaleAttrValue      pmsSkuSaleAttrValue
	PmsSpuComment            pmsSpuComment
	PmsSpuImage              pmsSpuImage
	PmsSpuInfo               pmsSpuInfo
	PmsSpuInfoDesc           pmsSpuInfoDesc
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                       db,
		PmsAttr:                  q.PmsAttr.clone(db),
		PmsAttrAttrgroupRelation: q.PmsAttrAttrgroupRelation.clone(db),
		PmsAttrGroup:             q.PmsAttrGroup.clone(db),
		PmsBrand:                 q.PmsBrand.clone(db),
		PmsCategory:              q.PmsCategory.clone(db),
		PmsCategoryBrandRelation: q.PmsCategoryBrandRelation.clone(db),
		PmsCommentReplay:         q.PmsCommentReplay.clone(db),
		PmsProductAttrValue:      q.PmsProductAttrValue.clone(db),
		PmsSkuImage:              q.PmsSkuImage.clone(db),
		PmsSkuInfo:               q.PmsSkuInfo.clone(db),
		PmsSkuSaleAttrValue:      q.PmsSkuSaleAttrValue.clone(db),
		PmsSpuComment:            q.PmsSpuComment.clone(db),
		PmsSpuImage:              q.PmsSpuImage.clone(db),
		PmsSpuInfo:               q.PmsSpuInfo.clone(db),
		PmsSpuInfoDesc:           q.PmsSpuInfoDesc.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                       db,
		PmsAttr:                  q.PmsAttr.replaceDB(db),
		PmsAttrAttrgroupRelation: q.PmsAttrAttrgroupRelation.replaceDB(db),
		PmsAttrGroup:             q.PmsAttrGroup.replaceDB(db),
		PmsBrand:                 q.PmsBrand.replaceDB(db),
		PmsCategory:              q.PmsCategory.replaceDB(db),
		PmsCategoryBrandRelation: q.PmsCategoryBrandRelation.replaceDB(db),
		PmsCommentReplay:         q.PmsCommentReplay.replaceDB(db),
		PmsProductAttrValue:      q.PmsProductAttrValue.replaceDB(db),
		PmsSkuImage:              q.PmsSkuImage.replaceDB(db),
		PmsSkuInfo:               q.PmsSkuInfo.replaceDB(db),
		PmsSkuSaleAttrValue:      q.PmsSkuSaleAttrValue.replaceDB(db),
		PmsSpuComment:            q.PmsSpuComment.replaceDB(db),
		PmsSpuImage:              q.PmsSpuImage.replaceDB(db),
		PmsSpuInfo:               q.PmsSpuInfo.replaceDB(db),
		PmsSpuInfoDesc:           q.PmsSpuInfoDesc.replaceDB(db),
	}
}

type queryCtx struct {
	PmsAttr                  *pmsAttrDo
	PmsAttrAttrgroupRelation *pmsAttrAttrgroupRelationDo
	PmsAttrGroup             *pmsAttrGroupDo
	PmsBrand                 *pmsBrandDo
	PmsCategory              *pmsCategoryDo
	PmsCategoryBrandRelation *pmsCategoryBrandRelationDo
	PmsCommentReplay         *pmsCommentReplayDo
	PmsProductAttrValue      *pmsProductAttrValueDo
	PmsSkuImage              *pmsSkuImageDo
	PmsSkuInfo               *pmsSkuInfoDo
	PmsSkuSaleAttrValue      *pmsSkuSaleAttrValueDo
	PmsSpuComment            *pmsSpuCommentDo
	PmsSpuImage              *pmsSpuImageDo
	PmsSpuInfo               *pmsSpuInfoDo
	PmsSpuInfoDesc           *pmsSpuInfoDescDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		PmsAttr:                  q.PmsAttr.WithContext(ctx),
		PmsAttrAttrgroupRelation: q.PmsAttrAttrgroupRelation.WithContext(ctx),
		PmsAttrGroup:             q.PmsAttrGroup.WithContext(ctx),
		PmsBrand:                 q.PmsBrand.WithContext(ctx),
		PmsCategory:              q.PmsCategory.WithContext(ctx),
		PmsCategoryBrandRelation: q.PmsCategoryBrandRelation.WithContext(ctx),
		PmsCommentReplay:         q.PmsCommentReplay.WithContext(ctx),
		PmsProductAttrValue:      q.PmsProductAttrValue.WithContext(ctx),
		PmsSkuImage:              q.PmsSkuImage.WithContext(ctx),
		PmsSkuInfo:               q.PmsSkuInfo.WithContext(ctx),
		PmsSkuSaleAttrValue:      q.PmsSkuSaleAttrValue.WithContext(ctx),
		PmsSpuComment:            q.PmsSpuComment.WithContext(ctx),
		PmsSpuImage:              q.PmsSpuImage.WithContext(ctx),
		PmsSpuInfo:               q.PmsSpuInfo.WithContext(ctx),
		PmsSpuInfoDesc:           q.PmsSpuInfoDesc.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
