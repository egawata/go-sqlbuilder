package main

import (
	"strings"
)

func main() {

}

type Parts interface {
	ToSQL() string
}

type Builder struct {
	pSelect PartSelect
	pFroms  []*PartFrom
	pWheres whereClause
}

type PartSelect struct {
	Columns []string
}

type PartFrom struct {
	Table string
}

type whereClause interface {
	_where()
	ToSQL() string
}

type isWhere struct{}

func (i *isWhere) _where() {}

type PartWhere struct {
	isWhere
	op  string
	lhv string
	rhv string
}

func (p *PartWhere) ToSQL() string {
	panic("PartWhere.ToSQL() is not implemented yet")
}

type PartWhereRaw struct {
	isWhere
	exp string
}

func (p *PartWhereRaw) ToSQL() string {
	return " WHERE " + p.exp
}

func (p *PartSelect) ToSQL() string {
	return "SELECT " + strings.Join(p.Columns, ", ")
}

func (b *Builder) Select(s string) *Builder {
	b.pSelect = PartSelect{Columns: []string{s}}
	return b
}

func (b *Builder) From(s string) *Builder {
	b.pFroms = append(b.pFroms, &PartFrom{Table: s})
	return b
}

func (b *Builder) Where(s string) *Builder {
	b.pWheres = &PartWhereRaw{exp: s}
	return b
}

func (b *Builder) ToSQL() string {
	sql := b.pSelect.ToSQL()
	for _, f := range b.pFroms {
		sql += " FROM " + f.Table
	}
	if b.pWheres != nil {
		sql += b.pWheres.ToSQL()
	}
	return sql
}
