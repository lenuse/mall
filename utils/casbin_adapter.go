package utils

import (
	"errors"
	"fmt"
	"github.com/lenuse/mall/entity"
	"runtime"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

// Adapter represents the sqlx adapter for policy storage.
type Adapter struct {
	engine    sqlbuilder.Database
	tableName string
}

// AdapterOptions contains all possible configuration options.
type AdapterOptions struct {
	DriverName     string
	DataSourceName string
	TableName      string
	DB             sqlbuilder.Database
}

func finalizer(a *Adapter) {
	a.engine.Close()
}

func loadPolicyLine(line entity.CasbinRule, model model.Model) {
	lineText := line.PType
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}
	persist.LoadPolicyLine(lineText, model)
}

func savePolicyLine(ptype string, rule []string) entity.CasbinRule {
	line := entity.CasbinRule{}
	line.PType = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}
	return line
}

func (a *Adapter) dropTable() {
	_, err := a.engine.DeleteFrom(a.tableName).Exec()
	if err != nil {
		panic(err)
	}
}

func (a *Adapter) ensureTable() {
	_, err := a.engine.Query("SELECT 1 FROM ? LIMIT 1", a.tableName)
	if err != nil {
		panic(err)
	}
}

func (a *Adapter) insertPolicyLine(line entity.CasbinRule) (err error) {
	_, err = a.engine.InsertInto(a.tableName).Values(line).Exec()
	if err != nil {
		return
	}
	return
}

func (a *Adapter) deletePolicyLine(line entity.CasbinRule) (err error) {
	return a.engine.Collection(a.tableName).Find(line).Delete()
}

// NewAdapter is the constructor for Adapter
// Deprecated: Use NewAdapterFromOptions instead
func NewAdapter(driverName string, url db.ConnectionURL) *Adapter {
	database, err := open(driverName, url)
	if err != nil {
		panic(err)
	}
	a := &Adapter{
		engine:    database,
		tableName: "casbin_rule",
	}
	a.ensureTable()
	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)
	return a
}

func open(driverName string, url db.ConnectionURL) (sqlbuilder.Database, error) {
	switch driverName {
	case "postgresql":
		return postgresql.Open(url)
	case "mysql":
		return postgresql.Open(url)
	case "sqlite":
		return postgresql.Open(url)
	case "ql":
		return postgresql.Open(url)
	case "mssql":
		return postgresql.Open(url)
	case "mongo":
		return postgresql.Open(url)
	}
	return nil, errors.New("driver not found")
}

// NewAdapterByDB is the constructor for Adapter with existed connection
// Deprecated: Use NewAdapterFromOptions instead
//func NewAdapterByDB(db *sqlx.DB) *Adapter {
//	a := &Adapter{
//		db:        db,
//		tableName: "casbin_rule",
//	}
//	a.ensureTable()
//	return a
//}

// NewAdapterFromOptions is the constructor for Adapter with existed connection
//func NewAdapterFromOptions(opts *AdapterOptions) *Adapter {
//	a := &Adapter{tableName: "casbin_rule"}
//	if opts.TableName != "" {
//		a.tableName = opts.TableName
//	}
//	if opts.DB != nil {
//		a.engine = opts.DB
//	} else {
//		db, err := sqlx.Connect(opts.DriverName, opts.DataSourceName)
//		if err != nil {
//			panic(err)
//		}
//		a.engine = db
//		runtime.SetFinalizer(a, finalizer)
//	}
//	a.ensureTable()
//	return a
//}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	var lines []entity.CasbinRule
	err := a.engine.Collection(a.tableName).Find().All(&lines)
	if err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	return nil
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) (err error) {
	a.dropTable()
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			err = a.insertPolicyLine(line)
			if err != nil {
				return
			}
		}
	}
	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			err = a.insertPolicyLine(line)
			if err != nil {
				return
			}
		}
	}
	return
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) (err error) {
	line := savePolicyLine(ptype, rule)
	err = a.insertPolicyLine(line)
	if err != nil {
		return
	}
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) (err error) {
	line := savePolicyLine(ptype, rule)
	err = a.deletePolicyLine(line)
	if err != nil {
		return
	}
	return err
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) (err error) {
	line := entity.CasbinRule{}
	line.PType = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	err = a.rawDelete(&line)
	if err != nil {
		return
	}
	return
}

func (a *Adapter) rawDelete(line *entity.CasbinRule) (err error) {
	queryArgs := []interface{}{line.PType}
	query := fmt.Sprintf("DELETE FROM `%s` WHERE p_type = ?", a.tableName)
	if line.V0 != "" {
		query += " AND v0 = ?"
		queryArgs = append(queryArgs, line.V0)
	}
	if line.V1 != "" {
		query += " AND v1 = ?"
		queryArgs = append(queryArgs, line.V1)
	}
	if line.V2 != "" {
		query += " AND v2 = ?"
		queryArgs = append(queryArgs, line.V2)
	}
	if line.V3 != "" {
		query += " AND v3 = ?"
		queryArgs = append(queryArgs, line.V3)
	}
	if line.V4 != "" {
		query += " AND v4 = ?"
		queryArgs = append(queryArgs, line.V4)
	}
	if line.V5 != "" {
		query += " AND v5 = ?"
		queryArgs = append(queryArgs, line.V5)
	}
	_, err = a.engine.Exec(query, queryArgs...)
	if err != nil {
		return
	}
	return
}
