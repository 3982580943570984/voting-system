// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"
	"voting-system/ent/generated/election"
	"voting-system/ent/generated/electionfilters"
	"voting-system/ent/generated/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ElectionFiltersQuery is the builder for querying ElectionFilters entities.
type ElectionFiltersQuery struct {
	config
	ctx          *QueryContext
	order        []electionfilters.OrderOption
	inters       []Interceptor
	predicates   []predicate.ElectionFilters
	withElection *ElectionQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ElectionFiltersQuery builder.
func (efq *ElectionFiltersQuery) Where(ps ...predicate.ElectionFilters) *ElectionFiltersQuery {
	efq.predicates = append(efq.predicates, ps...)
	return efq
}

// Limit the number of records to be returned by this query.
func (efq *ElectionFiltersQuery) Limit(limit int) *ElectionFiltersQuery {
	efq.ctx.Limit = &limit
	return efq
}

// Offset to start from.
func (efq *ElectionFiltersQuery) Offset(offset int) *ElectionFiltersQuery {
	efq.ctx.Offset = &offset
	return efq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (efq *ElectionFiltersQuery) Unique(unique bool) *ElectionFiltersQuery {
	efq.ctx.Unique = &unique
	return efq
}

// Order specifies how the records should be ordered.
func (efq *ElectionFiltersQuery) Order(o ...electionfilters.OrderOption) *ElectionFiltersQuery {
	efq.order = append(efq.order, o...)
	return efq
}

// QueryElection chains the current query on the "election" edge.
func (efq *ElectionFiltersQuery) QueryElection() *ElectionQuery {
	query := (&ElectionClient{config: efq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := efq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := efq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(electionfilters.Table, electionfilters.FieldID, selector),
			sqlgraph.To(election.Table, election.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, electionfilters.ElectionTable, electionfilters.ElectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(efq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ElectionFilters entity from the query.
// Returns a *NotFoundError when no ElectionFilters was found.
func (efq *ElectionFiltersQuery) First(ctx context.Context) (*ElectionFilters, error) {
	nodes, err := efq.Limit(1).All(setContextOp(ctx, efq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{electionfilters.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (efq *ElectionFiltersQuery) FirstX(ctx context.Context) *ElectionFilters {
	node, err := efq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ElectionFilters ID from the query.
// Returns a *NotFoundError when no ElectionFilters ID was found.
func (efq *ElectionFiltersQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = efq.Limit(1).IDs(setContextOp(ctx, efq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{electionfilters.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (efq *ElectionFiltersQuery) FirstIDX(ctx context.Context) int {
	id, err := efq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ElectionFilters entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ElectionFilters entity is found.
// Returns a *NotFoundError when no ElectionFilters entities are found.
func (efq *ElectionFiltersQuery) Only(ctx context.Context) (*ElectionFilters, error) {
	nodes, err := efq.Limit(2).All(setContextOp(ctx, efq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{electionfilters.Label}
	default:
		return nil, &NotSingularError{electionfilters.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (efq *ElectionFiltersQuery) OnlyX(ctx context.Context) *ElectionFilters {
	node, err := efq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ElectionFilters ID in the query.
// Returns a *NotSingularError when more than one ElectionFilters ID is found.
// Returns a *NotFoundError when no entities are found.
func (efq *ElectionFiltersQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = efq.Limit(2).IDs(setContextOp(ctx, efq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{electionfilters.Label}
	default:
		err = &NotSingularError{electionfilters.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (efq *ElectionFiltersQuery) OnlyIDX(ctx context.Context) int {
	id, err := efq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ElectionFiltersSlice.
func (efq *ElectionFiltersQuery) All(ctx context.Context) ([]*ElectionFilters, error) {
	ctx = setContextOp(ctx, efq.ctx, ent.OpQueryAll)
	if err := efq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ElectionFilters, *ElectionFiltersQuery]()
	return withInterceptors[[]*ElectionFilters](ctx, efq, qr, efq.inters)
}

// AllX is like All, but panics if an error occurs.
func (efq *ElectionFiltersQuery) AllX(ctx context.Context) []*ElectionFilters {
	nodes, err := efq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ElectionFilters IDs.
func (efq *ElectionFiltersQuery) IDs(ctx context.Context) (ids []int, err error) {
	if efq.ctx.Unique == nil && efq.path != nil {
		efq.Unique(true)
	}
	ctx = setContextOp(ctx, efq.ctx, ent.OpQueryIDs)
	if err = efq.Select(electionfilters.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (efq *ElectionFiltersQuery) IDsX(ctx context.Context) []int {
	ids, err := efq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (efq *ElectionFiltersQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, efq.ctx, ent.OpQueryCount)
	if err := efq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, efq, querierCount[*ElectionFiltersQuery](), efq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (efq *ElectionFiltersQuery) CountX(ctx context.Context) int {
	count, err := efq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (efq *ElectionFiltersQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, efq.ctx, ent.OpQueryExist)
	switch _, err := efq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (efq *ElectionFiltersQuery) ExistX(ctx context.Context) bool {
	exist, err := efq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ElectionFiltersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (efq *ElectionFiltersQuery) Clone() *ElectionFiltersQuery {
	if efq == nil {
		return nil
	}
	return &ElectionFiltersQuery{
		config:       efq.config,
		ctx:          efq.ctx.Clone(),
		order:        append([]electionfilters.OrderOption{}, efq.order...),
		inters:       append([]Interceptor{}, efq.inters...),
		predicates:   append([]predicate.ElectionFilters{}, efq.predicates...),
		withElection: efq.withElection.Clone(),
		// clone intermediate query.
		sql:  efq.sql.Clone(),
		path: efq.path,
	}
}

// WithElection tells the query-builder to eager-load the nodes that are connected to
// the "election" edge. The optional arguments are used to configure the query builder of the edge.
func (efq *ElectionFiltersQuery) WithElection(opts ...func(*ElectionQuery)) *ElectionFiltersQuery {
	query := (&ElectionClient{config: efq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	efq.withElection = query
	return efq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HasFirstName bool `json:"has_first_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ElectionFilters.Query().
//		GroupBy(electionfilters.FieldHasFirstName).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (efq *ElectionFiltersQuery) GroupBy(field string, fields ...string) *ElectionFiltersGroupBy {
	efq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ElectionFiltersGroupBy{build: efq}
	grbuild.flds = &efq.ctx.Fields
	grbuild.label = electionfilters.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HasFirstName bool `json:"has_first_name,omitempty"`
//	}
//
//	client.ElectionFilters.Query().
//		Select(electionfilters.FieldHasFirstName).
//		Scan(ctx, &v)
func (efq *ElectionFiltersQuery) Select(fields ...string) *ElectionFiltersSelect {
	efq.ctx.Fields = append(efq.ctx.Fields, fields...)
	sbuild := &ElectionFiltersSelect{ElectionFiltersQuery: efq}
	sbuild.label = electionfilters.Label
	sbuild.flds, sbuild.scan = &efq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ElectionFiltersSelect configured with the given aggregations.
func (efq *ElectionFiltersQuery) Aggregate(fns ...AggregateFunc) *ElectionFiltersSelect {
	return efq.Select().Aggregate(fns...)
}

func (efq *ElectionFiltersQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range efq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, efq); err != nil {
				return err
			}
		}
	}
	for _, f := range efq.ctx.Fields {
		if !electionfilters.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if efq.path != nil {
		prev, err := efq.path(ctx)
		if err != nil {
			return err
		}
		efq.sql = prev
	}
	return nil
}

func (efq *ElectionFiltersQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ElectionFilters, error) {
	var (
		nodes       = []*ElectionFilters{}
		withFKs     = efq.withFKs
		_spec       = efq.querySpec()
		loadedTypes = [1]bool{
			efq.withElection != nil,
		}
	)
	if efq.withElection != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, electionfilters.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ElectionFilters).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ElectionFilters{config: efq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, efq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := efq.withElection; query != nil {
		if err := efq.loadElection(ctx, query, nodes, nil,
			func(n *ElectionFilters, e *Election) { n.Edges.Election = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (efq *ElectionFiltersQuery) loadElection(ctx context.Context, query *ElectionQuery, nodes []*ElectionFilters, init func(*ElectionFilters), assign func(*ElectionFilters, *Election)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ElectionFilters)
	for i := range nodes {
		if nodes[i].election_filters == nil {
			continue
		}
		fk := *nodes[i].election_filters
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(election.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "election_filters" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (efq *ElectionFiltersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := efq.querySpec()
	_spec.Node.Columns = efq.ctx.Fields
	if len(efq.ctx.Fields) > 0 {
		_spec.Unique = efq.ctx.Unique != nil && *efq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, efq.driver, _spec)
}

func (efq *ElectionFiltersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(electionfilters.Table, electionfilters.Columns, sqlgraph.NewFieldSpec(electionfilters.FieldID, field.TypeInt))
	_spec.From = efq.sql
	if unique := efq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if efq.path != nil {
		_spec.Unique = true
	}
	if fields := efq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, electionfilters.FieldID)
		for i := range fields {
			if fields[i] != electionfilters.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := efq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := efq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := efq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := efq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (efq *ElectionFiltersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(efq.driver.Dialect())
	t1 := builder.Table(electionfilters.Table)
	columns := efq.ctx.Fields
	if len(columns) == 0 {
		columns = electionfilters.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if efq.sql != nil {
		selector = efq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if efq.ctx.Unique != nil && *efq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range efq.predicates {
		p(selector)
	}
	for _, p := range efq.order {
		p(selector)
	}
	if offset := efq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := efq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ElectionFiltersGroupBy is the group-by builder for ElectionFilters entities.
type ElectionFiltersGroupBy struct {
	selector
	build *ElectionFiltersQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (efgb *ElectionFiltersGroupBy) Aggregate(fns ...AggregateFunc) *ElectionFiltersGroupBy {
	efgb.fns = append(efgb.fns, fns...)
	return efgb
}

// Scan applies the selector query and scans the result into the given value.
func (efgb *ElectionFiltersGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, efgb.build.ctx, ent.OpQueryGroupBy)
	if err := efgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ElectionFiltersQuery, *ElectionFiltersGroupBy](ctx, efgb.build, efgb, efgb.build.inters, v)
}

func (efgb *ElectionFiltersGroupBy) sqlScan(ctx context.Context, root *ElectionFiltersQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(efgb.fns))
	for _, fn := range efgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*efgb.flds)+len(efgb.fns))
		for _, f := range *efgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*efgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := efgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ElectionFiltersSelect is the builder for selecting fields of ElectionFilters entities.
type ElectionFiltersSelect struct {
	*ElectionFiltersQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (efs *ElectionFiltersSelect) Aggregate(fns ...AggregateFunc) *ElectionFiltersSelect {
	efs.fns = append(efs.fns, fns...)
	return efs
}

// Scan applies the selector query and scans the result into the given value.
func (efs *ElectionFiltersSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, efs.ctx, ent.OpQuerySelect)
	if err := efs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ElectionFiltersQuery, *ElectionFiltersSelect](ctx, efs.ElectionFiltersQuery, efs, efs.inters, v)
}

func (efs *ElectionFiltersSelect) sqlScan(ctx context.Context, root *ElectionFiltersQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(efs.fns))
	for _, fn := range efs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*efs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := efs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
