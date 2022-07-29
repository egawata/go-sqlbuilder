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

func TestSelectFromWhere(t *testing.T) {
	b := &builder.Builder{}
	b = b.Select("col1").From("table1").Where("col2 = value2")
	assert.Equal(t, "SELECT col1 FROM table1 WHERE col2 = value2", b.ToSQL())
}
