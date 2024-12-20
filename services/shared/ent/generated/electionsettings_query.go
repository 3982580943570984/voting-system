// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"
	"shared/ent/generated/election"
	"shared/ent/generated/electionsettings"
	"shared/ent/generated/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ElectionSettingsQuery is the builder for querying ElectionSettings entities.
type ElectionSettingsQuery struct {
	config
	ctx          *QueryContext
	order        []electionsettings.OrderOption
	inters       []Interceptor
	predicates   []predicate.ElectionSettings
	withElection *ElectionQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ElectionSettingsQuery builder.
func (esq *ElectionSettingsQuery) Where(ps ...predicate.ElectionSettings) *ElectionSettingsQuery {
	esq.predicates = append(esq.predicates, ps...)
	return esq
}

// Limit the number of records to be returned by this query.
func (esq *ElectionSettingsQuery) Limit(limit int) *ElectionSettingsQuery {
	esq.ctx.Limit = &limit
	return esq
}

// Offset to start from.
func (esq *ElectionSettingsQuery) Offset(offset int) *ElectionSettingsQuery {
	esq.ctx.Offset = &offset
	return esq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (esq *ElectionSettingsQuery) Unique(unique bool) *ElectionSettingsQuery {
	esq.ctx.Unique = &unique
	return esq
}

// Order specifies how the records should be ordered.
func (esq *ElectionSettingsQuery) Order(o ...electionsettings.OrderOption) *ElectionSettingsQuery {
	esq.order = append(esq.order, o...)
	return esq
}

// QueryElection chains the current query on the "election" edge.
func (esq *ElectionSettingsQuery) QueryElection() *ElectionQuery {
	query := (&ElectionClient{config: esq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := esq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := esq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(electionsettings.Table, electionsettings.FieldID, selector),
			sqlgraph.To(election.Table, election.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, electionsettings.ElectionTable, electionsettings.ElectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(esq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ElectionSettings entity from the query.
// Returns a *NotFoundError when no ElectionSettings was found.
func (esq *ElectionSettingsQuery) First(ctx context.Context) (*ElectionSettings, error) {
	nodes, err := esq.Limit(1).All(setContextOp(ctx, esq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{electionsettings.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (esq *ElectionSettingsQuery) FirstX(ctx context.Context) *ElectionSettings {
	node, err := esq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ElectionSettings ID from the query.
// Returns a *NotFoundError when no ElectionSettings ID was found.
func (esq *ElectionSettingsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = esq.Limit(1).IDs(setContextOp(ctx, esq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{electionsettings.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (esq *ElectionSettingsQuery) FirstIDX(ctx context.Context) int {
	id, err := esq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ElectionSettings entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ElectionSettings entity is found.
// Returns a *NotFoundError when no ElectionSettings entities are found.
func (esq *ElectionSettingsQuery) Only(ctx context.Context) (*ElectionSettings, error) {
	nodes, err := esq.Limit(2).All(setContextOp(ctx, esq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{electionsettings.Label}
	default:
		return nil, &NotSingularError{electionsettings.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (esq *ElectionSettingsQuery) OnlyX(ctx context.Context) *ElectionSettings {
	node, err := esq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ElectionSettings ID in the query.
// Returns a *NotSingularError when more than one ElectionSettings ID is found.
// Returns a *NotFoundError when no entities are found.
func (esq *ElectionSettingsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = esq.Limit(2).IDs(setContextOp(ctx, esq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{electionsettings.Label}
	default:
		err = &NotSingularError{electionsettings.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (esq *ElectionSettingsQuery) OnlyIDX(ctx context.Context) int {
	id, err := esq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ElectionSettingsSlice.
func (esq *ElectionSettingsQuery) All(ctx context.Context) ([]*ElectionSettings, error) {
	ctx = setContextOp(ctx, esq.ctx, ent.OpQueryAll)
	if err := esq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ElectionSettings, *ElectionSettingsQuery]()
	return withInterceptors[[]*ElectionSettings](ctx, esq, qr, esq.inters)
}

// AllX is like All, but panics if an error occurs.
func (esq *ElectionSettingsQuery) AllX(ctx context.Context) []*ElectionSettings {
	nodes, err := esq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ElectionSettings IDs.
func (esq *ElectionSettingsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if esq.ctx.Unique == nil && esq.path != nil {
		esq.Unique(true)
	}
	ctx = setContextOp(ctx, esq.ctx, ent.OpQueryIDs)
	if err = esq.Select(electionsettings.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (esq *ElectionSettingsQuery) IDsX(ctx context.Context) []int {
	ids, err := esq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (esq *ElectionSettingsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, esq.ctx, ent.OpQueryCount)
	if err := esq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, esq, querierCount[*ElectionSettingsQuery](), esq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (esq *ElectionSettingsQuery) CountX(ctx context.Context) int {
	count, err := esq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (esq *ElectionSettingsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, esq.ctx, ent.OpQueryExist)
	switch _, err := esq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (esq *ElectionSettingsQuery) ExistX(ctx context.Context) bool {
	exist, err := esq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ElectionSettingsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (esq *ElectionSettingsQuery) Clone() *ElectionSettingsQuery {
	if esq == nil {
		return nil
	}
	return &ElectionSettingsQuery{
		config:       esq.config,
		ctx:          esq.ctx.Clone(),
		order:        append([]electionsettings.OrderOption{}, esq.order...),
		inters:       append([]Interceptor{}, esq.inters...),
		predicates:   append([]predicate.ElectionSettings{}, esq.predicates...),
		withElection: esq.withElection.Clone(),
		// clone intermediate query.
		sql:  esq.sql.Clone(),
		path: esq.path,
	}
}

// WithElection tells the query-builder to eager-load the nodes that are connected to
// the "election" edge. The optional arguments are used to configure the query builder of the edge.
func (esq *ElectionSettingsQuery) WithElection(opts ...func(*ElectionQuery)) *ElectionSettingsQuery {
	query := (&ElectionClient{config: esq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	esq.withElection = query
	return esq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ElectionSettings.Query().
//		GroupBy(electionsettings.FieldCreateTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (esq *ElectionSettingsQuery) GroupBy(field string, fields ...string) *ElectionSettingsGroupBy {
	esq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ElectionSettingsGroupBy{build: esq}
	grbuild.flds = &esq.ctx.Fields
	grbuild.label = electionsettings.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.ElectionSettings.Query().
//		Select(electionsettings.FieldCreateTime).
//		Scan(ctx, &v)
func (esq *ElectionSettingsQuery) Select(fields ...string) *ElectionSettingsSelect {
	esq.ctx.Fields = append(esq.ctx.Fields, fields...)
	sbuild := &ElectionSettingsSelect{ElectionSettingsQuery: esq}
	sbuild.label = electionsettings.Label
	sbuild.flds, sbuild.scan = &esq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ElectionSettingsSelect configured with the given aggregations.
func (esq *ElectionSettingsQuery) Aggregate(fns ...AggregateFunc) *ElectionSettingsSelect {
	return esq.Select().Aggregate(fns...)
}

func (esq *ElectionSettingsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range esq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, esq); err != nil {
				return err
			}
		}
	}
	for _, f := range esq.ctx.Fields {
		if !electionsettings.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if esq.path != nil {
		prev, err := esq.path(ctx)
		if err != nil {
			return err
		}
		esq.sql = prev
	}
	return nil
}

func (esq *ElectionSettingsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ElectionSettings, error) {
	var (
		nodes       = []*ElectionSettings{}
		withFKs     = esq.withFKs
		_spec       = esq.querySpec()
		loadedTypes = [1]bool{
			esq.withElection != nil,
		}
	)
	if esq.withElection != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, electionsettings.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ElectionSettings).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ElectionSettings{config: esq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, esq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := esq.withElection; query != nil {
		if err := esq.loadElection(ctx, query, nodes, nil,
			func(n *ElectionSettings, e *Election) { n.Edges.Election = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (esq *ElectionSettingsQuery) loadElection(ctx context.Context, query *ElectionQuery, nodes []*ElectionSettings, init func(*ElectionSettings), assign func(*ElectionSettings, *Election)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ElectionSettings)
	for i := range nodes {
		if nodes[i].election_settings == nil {
			continue
		}
		fk := *nodes[i].election_settings
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
			return fmt.Errorf(`unexpected foreign-key "election_settings" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (esq *ElectionSettingsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := esq.querySpec()
	_spec.Node.Columns = esq.ctx.Fields
	if len(esq.ctx.Fields) > 0 {
		_spec.Unique = esq.ctx.Unique != nil && *esq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, esq.driver, _spec)
}

func (esq *ElectionSettingsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(electionsettings.Table, electionsettings.Columns, sqlgraph.NewFieldSpec(electionsettings.FieldID, field.TypeInt))
	_spec.From = esq.sql
	if unique := esq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if esq.path != nil {
		_spec.Unique = true
	}
	if fields := esq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, electionsettings.FieldID)
		for i := range fields {
			if fields[i] != electionsettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := esq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := esq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := esq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := esq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (esq *ElectionSettingsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(esq.driver.Dialect())
	t1 := builder.Table(electionsettings.Table)
	columns := esq.ctx.Fields
	if len(columns) == 0 {
		columns = electionsettings.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if esq.sql != nil {
		selector = esq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if esq.ctx.Unique != nil && *esq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range esq.predicates {
		p(selector)
	}
	for _, p := range esq.order {
		p(selector)
	}
	if offset := esq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := esq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ElectionSettingsGroupBy is the group-by builder for ElectionSettings entities.
type ElectionSettingsGroupBy struct {
	selector
	build *ElectionSettingsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (esgb *ElectionSettingsGroupBy) Aggregate(fns ...AggregateFunc) *ElectionSettingsGroupBy {
	esgb.fns = append(esgb.fns, fns...)
	return esgb
}

// Scan applies the selector query and scans the result into the given value.
func (esgb *ElectionSettingsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, esgb.build.ctx, ent.OpQueryGroupBy)
	if err := esgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ElectionSettingsQuery, *ElectionSettingsGroupBy](ctx, esgb.build, esgb, esgb.build.inters, v)
}

func (esgb *ElectionSettingsGroupBy) sqlScan(ctx context.Context, root *ElectionSettingsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(esgb.fns))
	for _, fn := range esgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*esgb.flds)+len(esgb.fns))
		for _, f := range *esgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*esgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := esgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ElectionSettingsSelect is the builder for selecting fields of ElectionSettings entities.
type ElectionSettingsSelect struct {
	*ElectionSettingsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ess *ElectionSettingsSelect) Aggregate(fns ...AggregateFunc) *ElectionSettingsSelect {
	ess.fns = append(ess.fns, fns...)
	return ess
}

// Scan applies the selector query and scans the result into the given value.
func (ess *ElectionSettingsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ess.ctx, ent.OpQuerySelect)
	if err := ess.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ElectionSettingsQuery, *ElectionSettingsSelect](ctx, ess.ElectionSettingsQuery, ess, ess.inters, v)
}

func (ess *ElectionSettingsSelect) sqlScan(ctx context.Context, root *ElectionSettingsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ess.fns))
	for _, fn := range ess.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ess.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ess.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
