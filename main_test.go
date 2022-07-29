package main_test

import (
	"testing"

	builder "github.com/egawata/go-sqlbuilder"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	b := &builder.Builder{}
	b = b.Select("col1")
	assert.Equal(t, "SELECT col1", b.ToSQL())
}

func TestSelectFrom(t *testing.T) {
	b := &builder.Builder{}
	b = b.Select("col1").From("table1")
	assert.Equal(t, "SELECT col1 FROM table1", b.ToSQL())
}

func TestWhereRaw(t *testing.T) {
	b := &builder.Builder{}
	b = b.Select("col1").From("table1").WhereRaw("col2 = value2")
	assert.Equal(t, "SELECT col1 FROM table1 WHERE col2 = value2", b.ToSQL())
}

func TestWhereObj(t *testing.T) {
	b := &builder.Builder{}
	b = b.Select("col1").From("table1").Where(builder.Eq("col2", "value2"))
	assert.Equal(t, "SELECT col1 FROM table1 WHERE col2 = value2", b.ToSQL())
}

func TestAddWhere(t *testing.T) {
	b := &builder.Builder{}
	b = b.Select("col1").
		From("table1").
		Where(builder.Eq("col2", "value2")).
		AddWhere(builder.Eq("col3", "value3"))
	assert.Equal(t, "SELECT col1 FROM table1 WHERE col2 = value2 AND col3 = value3", b.ToSQL())
}

func TestInnerJoin(t *testing.T) {
	b := &builder.Builder{}
	b = b.Select("col1").From("table1").InnerJoin("table2 ON table1.id = table2.table1id")
	assert.Equal(t, "SELECT col1 FROM table1 INNER JOIN table2 ON table1.id = table2.table1id", b.ToSQL())
}
