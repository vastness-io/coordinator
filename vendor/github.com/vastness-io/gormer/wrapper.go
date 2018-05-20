package gormer

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type DB interface {
	AddError(err error) error
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) DB
	AddIndex(indexName string, columns ...string) DB
	AddUniqueIndex(indexName string, columns ...string) DB
	Assign(attrs ...interface{}) DB
	Association(column string) *gorm.Association
	Attrs(attrs ...interface{}) DB
	AutoMigrate(values ...interface{}) DB
	Begin() DB
	BlockGlobalUpdate(enable bool) DB
	Callback() *gorm.Callback
	Close() error
	Commit() DB
	CommonDB() gorm.SQLCommon
	Count(value interface{}) DB
	Create(value interface{}) DB
	CreateTable(models ...interface{}) DB
	DB() *sql.DB
	Debug() DB
	Delete(value interface{}, where ...interface{}) DB
	Dialect() gorm.Dialect
	DropColumn(column string) DB
	DropTable(values ...interface{}) DB
	DropTableIfExists(values ...interface{}) DB
	Exec(sql string, values ...interface{}) DB
	Find(out interface{}, where ...interface{}) DB
	First(out interface{}, where ...interface{}) DB
	FirstOrCreate(out interface{}, where ...interface{}) DB
	FirstOrInit(out interface{}, where ...interface{}) DB
	Get(name string) (value interface{}, ok bool)
	GetErrors() []error
	Error() error
	Group(query string) DB
	HasBlockGlobalUpdate() bool
	HasTable(value interface{}) bool
	Having(query interface{}, values ...interface{}) DB
	InstantSet(name string, value interface{}) DB
	Joins(query string, args ...interface{}) DB
	Last(out interface{}, where ...interface{}) DB
	Limit(limit interface{}) DB
	LogMode(enable bool) DB
	Model(value interface{}) DB
	ModifyColumn(column string, typ string) DB
	New() *gorm.DB
	NewRecord(value interface{}) bool
	NewScope(value interface{}) *gorm.Scope
	Not(query interface{}, args ...interface{}) DB
	Offset(offset interface{}) DB
	Omit(columns ...string) DB
	Or(query interface{}, args ...interface{}) DB
	Order(value interface{}, reorder ...bool) DB
	Pluck(column string, value interface{}) DB
	Preload(column string, conditions ...interface{}) DB
	Raw(sql string, values ...interface{}) DB
	RecordNotFound() bool
	Related(value interface{}, foreignKeys ...string) DB
	RemoveForeignKey(field string, dest string) DB
	RemoveIndex(indexName string) DB
	Rollback() DB
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Save(value interface{}) DB
	Scan(dest interface{}) DB
	ScanRows(rows *sql.Rows, result interface{}) error
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) DB
	Select(query interface{}, args ...interface{}) DB
	Set(name string, value interface{}) DB
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	SetLogger(log *gorm.Logger)
	SingularTable(enable bool)
	Table(name string) DB
	Take(out interface{}, where ...interface{}) DB
	Unscoped() DB
	Update(attrs ...interface{}) DB
	UpdateColumn(attrs ...interface{}) DB
	UpdateColumns(values interface{}) DB
	Updates(values interface{}, ignoreProtectedAttrs ...bool) DB
	Where(query interface{}, args ...interface{}) DB
}

type gormer struct {
	w *gorm.DB
}

func Wrap(db *gorm.DB) DB {
	return &gormer{db}
}

func (g *gormer) Close() error {
	return g.w.Close()
}

func (g *gormer) DB() *sql.DB {
	return g.w.DB()
}

func (g *gormer) New() *gorm.DB {
	return g.w.New()
}

func (g *gormer) Dialect() gorm.Dialect {
	return g.w.Dialect()
}

func (g *gormer) Take(out interface{}, where ...interface{}) DB {
	return Wrap(g.w.Take(out, where))
}

func (g *gormer) RemoveForeignKey(field string, dest string) DB {
	return Wrap(g.w.RemoveForeignKey(field, dest))
}

func (g *gormer) BlockGlobalUpdate(enable bool) DB {
	return Wrap(g.w.BlockGlobalUpdate(enable))
}

func (g *gormer) HasBlockGlobalUpdate() bool {
	return g.w.HasBlockGlobalUpdate()
}

func (g *gormer) NewScope(value interface{}) *gorm.Scope {
	return g.w.NewScope(value)
}

func (g *gormer) CommonDB() gorm.SQLCommon {
	return g.w.CommonDB()
}

func (g *gormer) Callback() *gorm.Callback {
	return g.w.Callback()
}

func (g *gormer) SetLogger(log *gorm.Logger) {
	g.w.SetLogger(log)
}

func (g *gormer) LogMode(enable bool) DB {
	return Wrap(g.w.LogMode(enable))
}

func (g *gormer) SingularTable(enable bool) {
	g.w.SingularTable(enable)
}

func (g *gormer) Where(query interface{}, args ...interface{}) DB {
	return Wrap(g.w.Where(query, args...))
}

func (g *gormer) Or(query interface{}, args ...interface{}) DB {
	return Wrap(g.w.Or(query, args...))
}

func (g *gormer) Not(query interface{}, args ...interface{}) DB {
	return Wrap(g.w.Not(query, args...))
}

func (g *gormer) Limit(value interface{}) DB {
	return Wrap(g.w.Limit(value))
}

func (g *gormer) Offset(value interface{}) DB {
	return Wrap(g.w.Offset(value))
}

func (g *gormer) Order(value interface{}, reorder ...bool) DB {
	return Wrap(g.w.Order(value, reorder...))
}

func (g *gormer) Select(query interface{}, args ...interface{}) DB {
	return Wrap(g.w.Select(query, args...))
}

func (g *gormer) Omit(columns ...string) DB {
	return Wrap(g.w.Omit(columns...))
}

func (g *gormer) Group(query string) DB {
	return Wrap(g.w.Group(query))
}

func (g *gormer) Having(query interface{}, values ...interface{}) DB {
	return Wrap(g.w.Having(query, values...))
}

func (g *gormer) Joins(query string, args ...interface{}) DB {
	return Wrap(g.w.Joins(query, args...))
}

func (g *gormer) Scopes(funcs ...func(*gorm.DB) *gorm.DB) DB {
	return Wrap(g.w.Scopes(funcs...))
}

func (g *gormer) Unscoped() DB {
	return Wrap(g.w.Unscoped())
}

func (g *gormer) Attrs(attrs ...interface{}) DB {
	return Wrap(g.w.Attrs(attrs...))
}

func (g *gormer) Assign(attrs ...interface{}) DB {
	return Wrap(g.w.Assign(attrs...))
}

func (g *gormer) First(out interface{}, where ...interface{}) DB {
	return Wrap(g.w.First(out, where...))
}

func (g *gormer) Last(out interface{}, where ...interface{}) DB {
	return Wrap(g.w.Last(out, where...))
}

func (g *gormer) Find(out interface{}, where ...interface{}) DB {
	return Wrap(g.w.Find(out, where...))
}

func (g *gormer) Scan(dest interface{}) DB {
	return Wrap(g.w.Scan(dest))
}

func (g *gormer) Row() *sql.Row {
	return g.w.Row()
}

func (g *gormer) Rows() (*sql.Rows, error) {
	return g.w.Rows()
}

func (g *gormer) ScanRows(rows *sql.Rows, result interface{}) error {
	return g.w.ScanRows(rows, result)
}

func (g *gormer) Pluck(column string, value interface{}) DB {
	return Wrap(g.w.Pluck(column, value))
}

func (g *gormer) Count(value interface{}) DB {
	return Wrap(g.w.Count(value))
}

func (g *gormer) Related(value interface{}, foreignKeys ...string) DB {
	return Wrap(g.w.Related(value, foreignKeys...))
}

func (g *gormer) FirstOrInit(out interface{}, where ...interface{}) DB {
	return Wrap(g.w.FirstOrInit(out, where...))
}

func (g *gormer) FirstOrCreate(out interface{}, where ...interface{}) DB {
	return Wrap(g.w.FirstOrCreate(out, where...))
}

func (g *gormer) Update(attrs ...interface{}) DB {
	return Wrap(g.w.Update(attrs...))
}

func (g *gormer) Updates(values interface{}, ignoreProtectedAttrs ...bool) DB {
	return Wrap(g.w.Updates(values, ignoreProtectedAttrs...))
}

func (g *gormer) UpdateColumn(attrs ...interface{}) DB {
	return Wrap(g.w.UpdateColumn(attrs...))
}

func (g *gormer) UpdateColumns(values interface{}) DB {
	return Wrap(g.w.UpdateColumns(values))
}

func (g *gormer) Save(value interface{}) DB {
	return Wrap(g.w.Save(value))
}

func (g *gormer) Create(value interface{}) DB {
	return Wrap(g.w.Create(value))
}

func (g *gormer) Delete(value interface{}, where ...interface{}) DB {
	return Wrap(g.w.Delete(value, where...))
}

func (g *gormer) Raw(sql string, values ...interface{}) DB {
	return Wrap(g.w.Raw(sql, values...))
}

func (g *gormer) Exec(sql string, values ...interface{}) DB {
	return Wrap(g.w.Exec(sql, values...))
}

func (g *gormer) Model(value interface{}) DB {
	return Wrap(g.w.Model(value))
}

func (g *gormer) Table(name string) DB {
	return Wrap(g.w.Table(name))
}

func (g *gormer) Debug() DB {
	return Wrap(g.w.Debug())
}

func (g *gormer) Begin() DB {
	return Wrap(g.w.Begin())
}

func (g *gormer) Commit() DB {
	return Wrap(g.w.Commit())
}

func (g *gormer) Rollback() DB {
	return Wrap(g.w.Rollback())
}

func (g *gormer) NewRecord(value interface{}) bool {
	return g.w.NewRecord(value)
}

func (g *gormer) RecordNotFound() bool {
	return g.w.RecordNotFound()
}

func (g *gormer) CreateTable(values ...interface{}) DB {
	return Wrap(g.w.CreateTable(values...))
}

func (g *gormer) DropTable(values ...interface{}) DB {
	return Wrap(g.w.DropTable(values...))
}

func (g *gormer) DropTableIfExists(values ...interface{}) DB {
	return Wrap(g.w.DropTableIfExists(values...))
}

func (g *gormer) HasTable(value interface{}) bool {
	return g.w.HasTable(value)
}

func (g *gormer) AutoMigrate(values ...interface{}) DB {
	return Wrap(g.w.AutoMigrate(values...))
}

func (g *gormer) ModifyColumn(column string, typ string) DB {
	return Wrap(g.w.ModifyColumn(column, typ))
}

func (g *gormer) DropColumn(column string) DB {
	return Wrap(g.w.DropColumn(column))
}

func (g *gormer) AddIndex(indexName string, columns ...string) DB {
	return Wrap(g.w.AddIndex(indexName, columns...))
}

func (g *gormer) AddUniqueIndex(indexName string, columns ...string) DB {
	return Wrap(g.w.AddUniqueIndex(indexName, columns...))
}

func (g *gormer) RemoveIndex(indexName string) DB {
	return Wrap(g.w.RemoveIndex(indexName))
}

func (g *gormer) Association(column string) *gorm.Association {
	return g.w.Association(column)
}

func (g *gormer) Preload(column string, conditions ...interface{}) DB {
	return Wrap(g.w.Preload(column, conditions...))
}

func (g *gormer) Set(name string, value interface{}) DB {
	return Wrap(g.w.Set(name, value))
}

func (g *gormer) InstantSet(name string, value interface{}) DB {
	return Wrap(g.w.InstantSet(name, value))
}

func (g *gormer) Get(name string) (interface{}, bool) {
	return g.w.Get(name)
}

func (g *gormer) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	g.w.SetJoinTableHandler(source, column, handler)
}

func (g *gormer) AddForeignKey(field string, dest string, onDelete string, onUpdate string) DB {
	return Wrap(g.w.AddForeignKey(field, dest, onDelete, onUpdate))
}

func (g *gormer) AddError(err error) error {
	return g.w.AddError(err)
}

func (g *gormer) GetErrors() (errors []error) {
	return g.w.GetErrors()
}

func (g *gormer) Error() error {
	return g.w.Error
}
