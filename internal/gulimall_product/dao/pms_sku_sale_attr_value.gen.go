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

func newPmsSkuSaleAttrValue(db *gorm.DB, opts ...gen.DOOption) pmsSkuSaleAttrValue {
	_pmsSkuSaleAttrValue := pmsSkuSaleAttrValue{}

	_pmsSkuSaleAttrValue.pmsSkuSaleAttrValueDo.UseDB(db, opts...)
	_pmsSkuSaleAttrValue.pmsSkuSaleAttrValueDo.UseModel(&model.PmsSkuSaleAttrValue{})

	tableName := _pmsSkuSaleAttrValue.pmsSkuSaleAttrValueDo.TableName()
	_pmsSkuSaleAttrValue.ALL = field.NewAsterisk(tableName)
	_pmsSkuSaleAttrValue.ID = field.NewInt64(tableName, "id")
	_pmsSkuSaleAttrValue.SkuID = field.NewInt64(tableName, "sku_id")
	_pmsSkuSaleAttrValue.AttrID = field.NewInt64(tableName, "attr_id")
	_pmsSkuSaleAttrValue.AttrName = field.NewString(tableName, "attr_name")
	_pmsSkuSaleAttrValue.AttrValue = field.NewString(tableName, "attr_value")
	_pmsSkuSaleAttrValue.AttrSort = field.NewInt32(tableName, "attr_sort")

	_pmsSkuSaleAttrValue.fillFieldMap()

	return _pmsSkuSaleAttrValue
}

type pmsSkuSaleAttrValue struct {
	pmsSkuSaleAttrValueDo

	ALL       field.Asterisk
	ID        field.Int64  // id
	SkuID     field.Int64  // sku_id
	AttrID    field.Int64  // attr_id
	AttrName  field.String // 销售属性名
	AttrValue field.String // 销售属性值
	AttrSort  field.Int32  // 顺序

	fieldMap map[string]field.Expr
}

func (p pmsSkuSaleAttrValue) Table(newTableName string) *pmsSkuSaleAttrValue {
	p.pmsSkuSaleAttrValueDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsSkuSaleAttrValue) As(alias string) *pmsSkuSaleAttrValue {
	p.pmsSkuSaleAttrValueDo.DO = *(p.pmsSkuSaleAttrValueDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsSkuSaleAttrValue) updateTableName(table string) *pmsSkuSaleAttrValue {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.SkuID = field.NewInt64(table, "sku_id")
	p.AttrID = field.NewInt64(table, "attr_id")
	p.AttrName = field.NewString(table, "attr_name")
	p.AttrValue = field.NewString(table, "attr_value")
	p.AttrSort = field.NewInt32(table, "attr_sort")

	p.fillFieldMap()

	return p
}

func (p *pmsSkuSaleAttrValue) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsSkuSaleAttrValue) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["sku_id"] = p.SkuID
	p.fieldMap["attr_id"] = p.AttrID
	p.fieldMap["attr_name"] = p.AttrName
	p.fieldMap["attr_value"] = p.AttrValue
	p.fieldMap["attr_sort"] = p.AttrSort
}

func (p pmsSkuSaleAttrValue) clone(db *gorm.DB) pmsSkuSaleAttrValue {
	p.pmsSkuSaleAttrValueDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsSkuSaleAttrValue) replaceDB(db *gorm.DB) pmsSkuSaleAttrValue {
	p.pmsSkuSaleAttrValueDo.ReplaceDB(db)
	return p
}

type pmsSkuSaleAttrValueDo struct{ gen.DO }

func (p pmsSkuSaleAttrValueDo) Debug() *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsSkuSaleAttrValueDo) WithContext(ctx context.Context) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsSkuSaleAttrValueDo) ReadDB() *pmsSkuSaleAttrValueDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsSkuSaleAttrValueDo) WriteDB() *pmsSkuSaleAttrValueDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsSkuSaleAttrValueDo) Session(config *gorm.Session) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsSkuSaleAttrValueDo) Clauses(conds ...clause.Expression) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsSkuSaleAttrValueDo) Returning(value interface{}, columns ...string) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsSkuSaleAttrValueDo) Not(conds ...gen.Condition) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsSkuSaleAttrValueDo) Or(conds ...gen.Condition) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsSkuSaleAttrValueDo) Select(conds ...field.Expr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsSkuSaleAttrValueDo) Where(conds ...gen.Condition) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsSkuSaleAttrValueDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pmsSkuSaleAttrValueDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pmsSkuSaleAttrValueDo) Order(conds ...field.Expr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsSkuSaleAttrValueDo) Distinct(cols ...field.Expr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsSkuSaleAttrValueDo) Omit(cols ...field.Expr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsSkuSaleAttrValueDo) Join(table schema.Tabler, on ...field.Expr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsSkuSaleAttrValueDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsSkuSaleAttrValueDo) RightJoin(table schema.Tabler, on ...field.Expr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsSkuSaleAttrValueDo) Group(cols ...field.Expr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsSkuSaleAttrValueDo) Having(conds ...gen.Condition) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsSkuSaleAttrValueDo) Limit(limit int) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsSkuSaleAttrValueDo) Offset(offset int) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsSkuSaleAttrValueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsSkuSaleAttrValueDo) Unscoped() *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsSkuSaleAttrValueDo) Create(values ...*model.PmsSkuSaleAttrValue) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsSkuSaleAttrValueDo) CreateInBatches(values []*model.PmsSkuSaleAttrValue, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsSkuSaleAttrValueDo) Save(values ...*model.PmsSkuSaleAttrValue) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsSkuSaleAttrValueDo) First() (*model.PmsSkuSaleAttrValue, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuSaleAttrValue), nil
	}
}

func (p pmsSkuSaleAttrValueDo) Take() (*model.PmsSkuSaleAttrValue, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuSaleAttrValue), nil
	}
}

func (p pmsSkuSaleAttrValueDo) Last() (*model.PmsSkuSaleAttrValue, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuSaleAttrValue), nil
	}
}

func (p pmsSkuSaleAttrValueDo) Find() ([]*model.PmsSkuSaleAttrValue, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsSkuSaleAttrValue), err
}

func (p pmsSkuSaleAttrValueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsSkuSaleAttrValue, err error) {
	buf := make([]*model.PmsSkuSaleAttrValue, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsSkuSaleAttrValueDo) FindInBatches(result *[]*model.PmsSkuSaleAttrValue, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsSkuSaleAttrValueDo) Attrs(attrs ...field.AssignExpr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsSkuSaleAttrValueDo) Assign(attrs ...field.AssignExpr) *pmsSkuSaleAttrValueDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsSkuSaleAttrValueDo) Joins(fields ...field.RelationField) *pmsSkuSaleAttrValueDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsSkuSaleAttrValueDo) Preload(fields ...field.RelationField) *pmsSkuSaleAttrValueDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsSkuSaleAttrValueDo) FirstOrInit() (*model.PmsSkuSaleAttrValue, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuSaleAttrValue), nil
	}
}

func (p pmsSkuSaleAttrValueDo) FirstOrCreate() (*model.PmsSkuSaleAttrValue, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuSaleAttrValue), nil
	}
}

func (p pmsSkuSaleAttrValueDo) FindByPage(offset int, limit int) (result []*model.PmsSkuSaleAttrValue, count int64, err error) {
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

func (p pmsSkuSaleAttrValueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsSkuSaleAttrValueDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsSkuSaleAttrValueDo) Delete(models ...*model.PmsSkuSaleAttrValue) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsSkuSaleAttrValueDo) withDO(do gen.Dao) *pmsSkuSaleAttrValueDo {
	p.DO = *do.(*gen.DO)
	return p
}
