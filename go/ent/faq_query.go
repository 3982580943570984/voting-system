// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"voting-system/ent/faq"
	"voting-system/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FAQQuery is the builder for querying FAQ entities.
type FAQQuery struct {
	config
	ctx        *QueryContext
	order      []faq.OrderOption
	inters     []Interceptor
	predicates []predicate.FAQ
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FAQQuery builder.
func (fq *FAQQuery) Where(ps ...predicate.FAQ) *FAQQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FAQQuery) Limit(limit int) *FAQQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FAQQuery) Offset(offset int) *FAQQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FAQQuery) Unique(unique bool) *FAQQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FAQQuery) Order(o ...faq.OrderOption) *FAQQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// First returns the first FAQ entity from the query.
// Returns a *NotFoundError when no FAQ was found.
func (fq *FAQQuery) First(ctx context.Context) (*FAQ, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{faq.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FAQQuery) FirstX(ctx context.Context) *FAQ {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FAQ ID from the query.
// Returns a *NotFoundError when no FAQ ID was found.
func (fq *FAQQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{faq.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FAQQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FAQ entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FAQ entity is found.
// Returns a *NotFoundError when no FAQ entities are found.
func (fq *FAQQuery) Only(ctx context.Context) (*FAQ, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{faq.Label}
	default:
		return nil, &NotSingularError{faq.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FAQQuery) OnlyX(ctx context.Context) *FAQ {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FAQ ID in the query.
// Returns a *NotSingularError when more than one FAQ ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FAQQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{faq.Label}
	default:
		err = &NotSingularError{faq.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FAQQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FAQs.
func (fq *FAQQuery) All(ctx context.Context) ([]*FAQ, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryAll)
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FAQ, *FAQQuery]()
	return withInterceptors[[]*FAQ](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FAQQuery) AllX(ctx context.Context) []*FAQ {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FAQ IDs.
func (fq *FAQQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryIDs)
	if err = fq.Select(faq.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FAQQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FAQQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryCount)
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FAQQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FAQQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FAQQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryExist)
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FAQQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FAQQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FAQQuery) Clone() *FAQQuery {
	if fq == nil {
		return nil
	}
	return &FAQQuery{
		config:     fq.config,
		ctx:        fq.ctx.Clone(),
		order:      append([]faq.OrderOption{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.FAQ{}, fq.predicates...),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (fq *FAQQuery) GroupBy(field string, fields ...string) *FAQGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FAQGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = faq.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (fq *FAQQuery) Select(fields ...string) *FAQSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FAQSelect{FAQQuery: fq}
	sbuild.label = faq.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FAQSelect configured with the given aggregations.
func (fq *FAQQuery) Aggregate(fns ...AggregateFunc) *FAQSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FAQQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !faq.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FAQQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FAQ, error) {
	var (
		nodes = []*FAQ{}
		_spec = fq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FAQ).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FAQ{config: fq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fq *FAQQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FAQQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(faq.Table, faq.Columns, sqlgraph.NewFieldSpec(faq.FieldID, field.TypeInt))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, faq.FieldID)
		for i := range fields {
			if fields[i] != faq.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FAQQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(faq.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = faq.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FAQGroupBy is the group-by builder for FAQ entities.
type FAQGroupBy struct {
	selector
	build *FAQQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FAQGroupBy) Aggregate(fns ...AggregateFunc) *FAQGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FAQGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, ent.OpQueryGroupBy)
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FAQQuery, *FAQGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FAQGroupBy) sqlScan(ctx context.Context, root *FAQQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FAQSelect is the builder for selecting fields of FAQ entities.
type FAQSelect struct {
	*FAQQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FAQSelect) Aggregate(fns ...AggregateFunc) *FAQSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FAQSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, ent.OpQuerySelect)
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FAQQuery, *FAQSelect](ctx, fs.FAQQuery, fs, fs.inters, v)
}

func (fs *FAQSelect) sqlScan(ctx context.Context, root *FAQQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
