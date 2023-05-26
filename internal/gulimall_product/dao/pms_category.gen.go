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

func newPmsCategory(db *gorm.DB, opts ...gen.DOOption) pmsCategory {
	_pmsCategory := pmsCategory{}

	_pmsCategory.pmsCategoryDo.UseDB(db, opts...)
	_pmsCategory.pmsCategoryDo.UseModel(&model.PmsCategory{})

	tableName := _pmsCategory.pmsCategoryDo.TableName()
	_pmsCategory.ALL = field.NewAsterisk(tableName)
	_pmsCategory.CatID = field.NewInt64(tableName, "cat_id")
	_pmsCategory.Name = field.NewString(tableName, "name")
	_pmsCategory.ParentCid = field.NewInt64(tableName, "parent_cid")
	_pmsCategory.CatLevel = field.NewInt32(tableName, "cat_level")
	_pmsCategory.ShowStatus = field.NewInt32(tableName, "show_status")
	_pmsCategory.Sort = field.NewInt32(tableName, "sort")
	_pmsCategory.Icon = field.NewString(tableName, "icon")
	_pmsCategory.ProductUnit = field.NewString(tableName, "product_unit")
	_pmsCategory.ProductCount = field.NewInt32(tableName, "product_count")

	_pmsCategory.fillFieldMap()

	return _pmsCategory
}

type pmsCategory struct {
	pmsCategoryDo

	ALL          field.Asterisk
	CatID        field.Int64  // 分类id
	Name         field.String // 分类名称
	ParentCid    field.Int64  // 父分类id
	CatLevel     field.Int32  // 层级
	ShowStatus   field.Int32  // 是否显示[0-不显示，1显示]
	Sort         field.Int32  // 排序
	Icon         field.String // 图标地址
	ProductUnit  field.String // 计量单位
	ProductCount field.Int32  // 商品数量

	fieldMap map[string]field.Expr
}

func (p pmsCategory) Table(newTableName string) *pmsCategory {
	p.pmsCategoryDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsCategory) As(alias string) *pmsCategory {
	p.pmsCategoryDo.DO = *(p.pmsCategoryDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsCategory) updateTableName(table string) *pmsCategory {
	p.ALL = field.NewAsterisk(table)
	p.CatID = field.NewInt64(table, "cat_id")
	p.Name = field.NewString(table, "name")
	p.ParentCid = field.NewInt64(table, "parent_cid")
	p.CatLevel = field.NewInt32(table, "cat_level")
	p.ShowStatus = field.NewInt32(table, "show_status")
	p.Sort = field.NewInt32(table, "sort")
	p.Icon = field.NewString(table, "icon")
	p.ProductUnit = field.NewString(table, "product_unit")
	p.ProductCount = field.NewInt32(table, "product_count")

	p.fillFieldMap()

	return p
}

func (p *pmsCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsCategory) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["cat_id"] = p.CatID
	p.fieldMap["name"] = p.Name
	p.fieldMap["parent_cid"] = p.ParentCid
	p.fieldMap["cat_level"] = p.CatLevel
	p.fieldMap["show_status"] = p.ShowStatus
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["icon"] = p.Icon
	p.fieldMap["product_unit"] = p.ProductUnit
	p.fieldMap["product_count"] = p.ProductCount
}

func (p pmsCategory) clone(db *gorm.DB) pmsCategory {
	p.pmsCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsCategory) replaceDB(db *gorm.DB) pmsCategory {
	p.pmsCategoryDo.ReplaceDB(db)
	return p
}

type pmsCategoryDo struct{ gen.DO }

func (p pmsCategoryDo) Debug() *pmsCategoryDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsCategoryDo) WithContext(ctx context.Context) *pmsCategoryDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsCategoryDo) ReadDB() *pmsCategoryDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsCategoryDo) WriteDB() *pmsCategoryDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsCategoryDo) Session(config *gorm.Session) *pmsCategoryDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsCategoryDo) Clauses(conds ...clause.Expression) *pmsCategoryDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsCategoryDo) Returning(value interface{}, columns ...string) *pmsCategoryDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsCategoryDo) Not(conds ...gen.Condition) *pmsCategoryDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsCategoryDo) Or(conds ...gen.Condition) *pmsCategoryDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsCategoryDo) Select(conds ...field.Expr) *pmsCategoryDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsCategoryDo) Where(conds ...gen.Condition) *pmsCategoryDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsCategoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pmsCategoryDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pmsCategoryDo) Order(conds ...field.Expr) *pmsCategoryDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsCategoryDo) Distinct(cols ...field.Expr) *pmsCategoryDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsCategoryDo) Omit(cols ...field.Expr) *pmsCategoryDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsCategoryDo) Join(table schema.Tabler, on ...field.Expr) *pmsCategoryDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pmsCategoryDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *pmsCategoryDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsCategoryDo) Group(cols ...field.Expr) *pmsCategoryDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsCategoryDo) Having(conds ...gen.Condition) *pmsCategoryDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsCategoryDo) Limit(limit int) *pmsCategoryDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsCategoryDo) Offset(offset int) *pmsCategoryDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pmsCategoryDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsCategoryDo) Unscoped() *pmsCategoryDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsCategoryDo) Create(values ...*model.PmsCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsCategoryDo) CreateInBatches(values []*model.PmsCategory, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsCategoryDo) Save(values ...*model.PmsCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsCategoryDo) First() (*model.PmsCategory, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsCategory), nil
	}
}

func (p pmsCategoryDo) Take() (*model.PmsCategory, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsCategory), nil
	}
}

func (p pmsCategoryDo) Last() (*model.PmsCategory, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsCategory), nil
	}
}

func (p pmsCategoryDo) Find() ([]*model.PmsCategory, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsCategory), err
}

func (p pmsCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsCategory, err error) {
	buf := make([]*model.PmsCategory, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsCategoryDo) FindInBatches(result *[]*model.PmsCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsCategoryDo) Attrs(attrs ...field.AssignExpr) *pmsCategoryDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsCategoryDo) Assign(attrs ...field.AssignExpr) *pmsCategoryDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsCategoryDo) Joins(fields ...field.RelationField) *pmsCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsCategoryDo) Preload(fields ...field.RelationField) *pmsCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsCategoryDo) FirstOrInit() (*model.PmsCategory, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsCategory), nil
	}
}

func (p pmsCategoryDo) FirstOrCreate() (*model.PmsCategory, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsCategory), nil
	}
}

func (p pmsCategoryDo) FindByPage(offset int, limit int) (result []*model.PmsCategory, count int64, err error) {
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

func (p pmsCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsCategoryDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsCategoryDo) Delete(models ...*model.PmsCategory) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsCategoryDo) withDO(do gen.Dao) *pmsCategoryDo {
	p.DO = *do.(*gen.DO)
	return p
}