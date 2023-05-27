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

func newPmsAttr(db *gorm.DB, opts ...gen.DOOption) pmsAttr {
	_pmsAttr := pmsAttr{}

	_pmsAttr.pmsAttrDo.UseDB(db, opts...)
	_pmsAttr.pmsAttrDo.UseModel(&model.PmsAttr{})

	tableName := _pmsAttr.pmsAttrDo.TableName()
	_pmsAttr.ALL = field.NewAsterisk(tableName)
	_pmsAttr.AttrID = field.NewInt64(tableName, "attr_id")
	_pmsAttr.AttrName = field.NewString(tableName, "attr_name")
	_pmsAttr.SearchType = field.NewInt32(tableName, "search_type")
	_pmsAttr.ValueType = field.NewInt32(tableName, "value_type")
	_pmsAttr.Icon = field.NewString(tableName, "icon")
	_pmsAttr.ValueSelect = field.NewString(tableName, "value_select")
	_pmsAttr.AttrType = field.NewInt32(tableName, "attr_type")
	_pmsAttr.Enable = field.NewInt64(tableName, "enable")
	_pmsAttr.CatelogID = field.NewInt64(tableName, "catelog_id")
	_pmsAttr.ShowDesc = field.NewInt32(tableName, "show_desc")

	_pmsAttr.fillFieldMap()

	return _pmsAttr
}

type pmsAttr struct {
	pmsAttrDo

	ALL         field.Asterisk
	AttrID      field.Int64  // 属性id
	AttrName    field.String // 属性名
	SearchType  field.Int32  // 是否需要检索[0-不需要，1-需要]
	ValueType   field.Int32  // 值类型[0-为单个值，1-可以选择多个值]
	Icon        field.String // 属性图标
	ValueSelect field.String // 可选值列表[用逗号分隔]
	AttrType    field.Int32  // 属性类型[0-销售属性，1-基本属性，2-既是销售属性又是基本属性]
	Enable      field.Int64  // 启用状态[0 - 禁用，1 - 启用]
	CatelogID   field.Int64  // 所属分类
	ShowDesc    field.Int32  // 快速展示【是否展示在介绍上；0-否 1-是】，在sku中仍然可以调整

	fieldMap map[string]field.Expr
}

func (p pmsAttr) Table(newTableName string) *pmsAttr {
	p.pmsAttrDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsAttr) As(alias string) *pmsAttr {
	p.pmsAttrDo.DO = *(p.pmsAttrDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsAttr) updateTableName(table string) *pmsAttr {
	p.ALL = field.NewAsterisk(table)
	p.AttrID = field.NewInt64(table, "attr_id")
	p.AttrName = field.NewString(table, "attr_name")
	p.SearchType = field.NewInt32(table, "search_type")
	p.ValueType = field.NewInt32(table, "value_type")
	p.Icon = field.NewString(table, "icon")
	p.ValueSelect = field.NewString(table, "value_select")
	p.AttrType = field.NewInt32(table, "attr_type")
	p.Enable = field.NewInt64(table, "enable")
	p.CatelogID = field.NewInt64(table, "catelog_id")
	p.ShowDesc = field.NewInt32(table, "show_desc")

	p.fillFieldMap()

	return p
}

func (p *pmsAttr) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsAttr) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["attr_id"] = p.AttrID
	p.fieldMap["attr_name"] = p.AttrName
	p.fieldMap["search_type"] = p.SearchType
	p.fieldMap["value_type"] = p.ValueType
	p.fieldMap["icon"] = p.Icon
	p.fieldMap["value_select"] = p.ValueSelect
	p.fieldMap["attr_type"] = p.AttrType
	p.fieldMap["enable"] = p.Enable
	p.fieldMap["catelog_id"] = p.CatelogID
	p.fieldMap["show_desc"] = p.ShowDesc
}

func (p pmsAttr) clone(db *gorm.DB) pmsAttr {
	p.pmsAttrDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsAttr) replaceDB(db *gorm.DB) pmsAttr {
	p.pmsAttrDo.ReplaceDB(db)
	return p
}

type pmsAttrDo struct{ gen.DO }

func (p pmsAttrDo) Debug() *pmsAttrDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsAttrDo) WithContext(ctx context.Context) *pmsAttrDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsAttrDo) ReadDB() *pmsAttrDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsAttrDo) WriteDB() *pmsAttrDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsAttrDo) Session(config *gorm.Session) *pmsAttrDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsAttrDo) Clauses(conds ...clause.Expression) *pmsAttrDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsAttrDo) Returning(value interface{}, columns ...string) *pmsAttrDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsAttrDo) Not(conds ...gen.Condition) *pmsAttrDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsAttrDo) Or(conds ...gen.Condition) *pmsAttrDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsAttrDo) Select(conds ...field.Expr) *pmsAttrDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsAttrDo) Where(conds ...gen.Condition) *pmsAttrDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsAttrDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pmsAttrDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pmsAttrDo) Order(conds ...field.Expr) *pmsAttrDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsAttrDo) Distinct(cols ...field.Expr) *pmsAttrDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsAttrDo) Omit(cols ...field.Expr) *pmsAttrDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsAttrDo) Join(table schema.Tabler, on ...field.Expr) *pmsAttrDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsAttrDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pmsAttrDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsAttrDo) RightJoin(table schema.Tabler, on ...field.Expr) *pmsAttrDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsAttrDo) Group(cols ...field.Expr) *pmsAttrDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsAttrDo) Having(conds ...gen.Condition) *pmsAttrDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsAttrDo) Limit(limit int) *pmsAttrDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsAttrDo) Offset(offset int) *pmsAttrDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsAttrDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pmsAttrDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsAttrDo) Unscoped() *pmsAttrDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsAttrDo) Create(values ...*model.PmsAttr) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsAttrDo) CreateInBatches(values []*model.PmsAttr, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsAttrDo) Save(values ...*model.PmsAttr) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsAttrDo) First() (*model.PmsAttr, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttr), nil
	}
}

func (p pmsAttrDo) Take() (*model.PmsAttr, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttr), nil
	}
}

func (p pmsAttrDo) Last() (*model.PmsAttr, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttr), nil
	}
}

func (p pmsAttrDo) Find() ([]*model.PmsAttr, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsAttr), err
}

func (p pmsAttrDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsAttr, err error) {
	buf := make([]*model.PmsAttr, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsAttrDo) FindInBatches(result *[]*model.PmsAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsAttrDo) Attrs(attrs ...field.AssignExpr) *pmsAttrDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsAttrDo) Assign(attrs ...field.AssignExpr) *pmsAttrDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsAttrDo) Joins(fields ...field.RelationField) *pmsAttrDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsAttrDo) Preload(fields ...field.RelationField) *pmsAttrDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsAttrDo) FirstOrInit() (*model.PmsAttr, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttr), nil
	}
}

func (p pmsAttrDo) FirstOrCreate() (*model.PmsAttr, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttr), nil
	}
}

func (p pmsAttrDo) FindByPage(offset int, limit int) (result []*model.PmsAttr, count int64, err error) {
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

func (p pmsAttrDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsAttrDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsAttrDo) Delete(models ...*model.PmsAttr) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsAttrDo) withDO(do gen.Dao) *pmsAttrDo {
	p.DO = *do.(*gen.DO)
	return p
}