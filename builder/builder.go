package builder

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	selectCols []string
	fromTable  string
	where      string
	orderby    string
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) Build() string {
	return fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s",
		strings.Join(qb.selectCols, ", "),
		qb.fromTable,
		qb.where,
		qb.orderby)
}

func (qb *QueryBuilder) Select(cols []string) {
	qb.selectCols = cols
}

func (qb *QueryBuilder) FromTable(table string) {
	qb.fromTable = table
}

func (qb *QueryBuilder) Where(condition string) {
	qb.where = condition
}

func (qb *QueryBuilder) OrderBy(order string) {
	qb.orderby = order
}
