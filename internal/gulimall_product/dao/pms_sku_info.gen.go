// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gulimall-product/internal/gulimall_product/model"
)

func newPmsSkuInfo(db *gorm.DB, opts ...gen.DOOption) pmsSkuInfo {
	_pmsSkuInfo := pmsSkuInfo{}

	_pmsSkuInfo.pmsSkuInfoDo.UseDB(db, opts...)
	_pmsSkuInfo.pmsSkuInfoDo.UseModel(&model.PmsSkuInfo{})

	tableName := _pmsSkuInfo.pmsSkuInfoDo.TableName()
	_pmsSkuInfo.ALL = field.NewAsterisk(tableName)
	_pmsSkuInfo.SkuID = field.NewInt64(tableName, "sku_id")
	_pmsSkuInfo.SpuID = field.NewInt64(tableName, "spu_id")
	_pmsSkuInfo.SkuName = field.NewString(tableName, "sku_name")
	_pmsSkuInfo.SkuDesc = field.NewString(tableName, "sku_desc")
	_pmsSkuInfo.CatalogID = field.NewInt64(tableName, "catalog_id")
	_pmsSkuInfo.BrandID = field.NewInt64(tableName, "brand_id")
	_pmsSkuInfo.SkuDefaultImg = field.NewString(tableName, "sku_default_img")
	_pmsSkuInfo.SkuTitle = field.NewString(tableName, "sku_title")
	_pmsSkuInfo.SkuSubtitle = field.NewString(tableName, "sku_subtitle")
	_pmsSkuInfo.Price = field.NewFloat64(tableName, "price")
	_pmsSkuInfo.SaleCount = field.NewInt64(tableName, "sale_count")

	_pmsSkuInfo.fillFieldMap()

	return _pmsSkuInfo
}

type pmsSkuInfo struct {
	pmsSkuInfoDo

	ALL           field.Asterisk
	SkuID         field.Int64   // skuId
	SpuID         field.Int64   // spuId
	SkuName       field.String  // sku名称
	SkuDesc       field.String  // sku介绍描述
	CatalogID     field.Int64   // 所属分类id
	BrandID       field.Int64   // 品牌id
	SkuDefaultImg field.String  // 默认图片
	SkuTitle      field.String  // 标题
	SkuSubtitle   field.String  // 副标题
	Price         field.Float64 // 价格
	SaleCount     field.Int64   // 销量

	fieldMap map[string]field.Expr
}

func (p pmsSkuInfo) Table(newTableName string) *pmsSkuInfo {
	p.pmsSkuInfoDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsSkuInfo) As(alias string) *pmsSkuInfo {
	p.pmsSkuInfoDo.DO = *(p.pmsSkuInfoDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsSkuInfo) updateTableName(table string) *pmsSkuInfo {
	p.ALL = field.NewAsterisk(table)
	p.SkuID = field.NewInt64(table, "sku_id")
	p.SpuID = field.NewInt64(table, "spu_id")
	p.SkuName = field.NewString(table, "sku_name")
	p.SkuDesc = field.NewString(table, "sku_desc")
	p.CatalogID = field.NewInt64(table, "catalog_id")
	p.BrandID = field.NewInt64(table, "brand_id")
	p.SkuDefaultImg = field.NewString(table, "sku_default_img")
	p.SkuTitle = field.NewString(table, "sku_title")
	p.SkuSubtitle = field.NewString(table, "sku_subtitle")
	p.Price = field.NewFloat64(table, "price")
	p.SaleCount = field.NewInt64(table, "sale_count")

	p.fillFieldMap()

	return p
}

func (p *pmsSkuInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsSkuInfo) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 11)
	p.fieldMap["sku_id"] = p.SkuID
	p.fieldMap["spu_id"] = p.SpuID
	p.fieldMap["sku_name"] = p.SkuName
	p.fieldMap["sku_desc"] = p.SkuDesc
	p.fieldMap["catalog_id"] = p.CatalogID
	p.fieldMap["brand_id"] = p.BrandID
	p.fieldMap["sku_default_img"] = p.SkuDefaultImg
	p.fieldMap["sku_title"] = p.SkuTitle
	p.fieldMap["sku_subtitle"] = p.SkuSubtitle
	p.fieldMap["price"] = p.Price
	p.fieldMap["sale_count"] = p.SaleCount
}

func (p pmsSkuInfo) clone(db *gorm.DB) pmsSkuInfo {
	p.pmsSkuInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsSkuInfo) replaceDB(db *gorm.DB) pmsSkuInfo {
	p.pmsSkuInfoDo.ReplaceDB(db)
	return p
}

type pmsSkuInfoDo struct{ gen.DO }

func (p pmsSkuInfoDo) Debug() *pmsSkuInfoDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsSkuInfoDo) WithContext(ctx context.Context) *pmsSkuInfoDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsSkuInfoDo) ReadDB() *pmsSkuInfoDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsSkuInfoDo) WriteDB() *pmsSkuInfoDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsSkuInfoDo) Session(config *gorm.Session) *pmsSkuInfoDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsSkuInfoDo) Clauses(conds ...clause.Expression) *pmsSkuInfoDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsSkuInfoDo) Returning(value interface{}, columns ...string) *pmsSkuInfoDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsSkuInfoDo) Not(conds ...gen.Condition) *pmsSkuInfoDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsSkuInfoDo) Or(conds ...gen.Condition) *pmsSkuInfoDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsSkuInfoDo) Select(conds ...field.Expr) *pmsSkuInfoDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsSkuInfoDo) Where(conds ...gen.Condition) *pmsSkuInfoDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsSkuInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pmsSkuInfoDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pmsSkuInfoDo) Order(conds ...field.Expr) *pmsSkuInfoDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsSkuInfoDo) Distinct(cols ...field.Expr) *pmsSkuInfoDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsSkuInfoDo) Omit(cols ...field.Expr) *pmsSkuInfoDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsSkuInfoDo) Join(table schema.Tabler, on ...field.Expr) *pmsSkuInfoDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsSkuInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pmsSkuInfoDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsSkuInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *pmsSkuInfoDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsSkuInfoDo) Group(cols ...field.Expr) *pmsSkuInfoDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsSkuInfoDo) Having(conds ...gen.Condition) *pmsSkuInfoDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsSkuInfoDo) Limit(limit int) *pmsSkuInfoDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsSkuInfoDo) Offset(offset int) *pmsSkuInfoDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsSkuInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pmsSkuInfoDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsSkuInfoDo) Unscoped() *pmsSkuInfoDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsSkuInfoDo) Create(values ...*model.PmsSkuInfo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsSkuInfoDo) CreateInBatches(values []*model.PmsSkuInfo, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsSkuInfoDo) Save(values ...*model.PmsSkuInfo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsSkuInfoDo) First() (*model.PmsSkuInfo, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuInfo), nil
	}
}

func (p pmsSkuInfoDo) Take() (*model.PmsSkuInfo, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuInfo), nil
	}
}

func (p pmsSkuInfoDo) Last() (*model.PmsSkuInfo, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuInfo), nil
	}
}

func (p pmsSkuInfoDo) Find() ([]*model.PmsSkuInfo, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsSkuInfo), err
}

func (p pmsSkuInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsSkuInfo, err error) {
	buf := make([]*model.PmsSkuInfo, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsSkuInfoDo) FindInBatches(result *[]*model.PmsSkuInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsSkuInfoDo) Attrs(attrs ...field.AssignExpr) *pmsSkuInfoDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsSkuInfoDo) Assign(attrs ...field.AssignExpr) *pmsSkuInfoDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsSkuInfoDo) Joins(fields ...field.RelationField) *pmsSkuInfoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsSkuInfoDo) Preload(fields ...field.RelationField) *pmsSkuInfoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsSkuInfoDo) FirstOrInit() (*model.PmsSkuInfo, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuInfo), nil
	}
}

func (p pmsSkuInfoDo) FirstOrCreate() (*model.PmsSkuInfo, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuInfo), nil
	}
}

func (p pmsSkuInfoDo) FindByPage(offset int, limit int) (result []*model.PmsSkuInfo, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pmsSkuInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsSkuInfoDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsSkuInfoDo) Delete(models ...*model.PmsSkuInfo) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsSkuInfoDo) withDO(do gen.Dao) *pmsSkuInfoDo {
	p.DO = *do.(*gen.DO)
	return p
}