// Code generated by ent, DO NOT EDIT.

package ent

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/database"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/generalspec"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/predicate"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/service"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DatabaseQuery is the builder for querying Database entities.
type DatabaseQuery struct {
	config
	ctx             *QueryContext
	order           []database.OrderOption
	inters          []Interceptor
	predicates      []predicate.Database
	withService     *ServiceQuery
	withGeneralspec *GeneralSpecQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DatabaseQuery builder.
func (dq *DatabaseQuery) Where(ps ...predicate.Database) *DatabaseQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DatabaseQuery) Limit(limit int) *DatabaseQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DatabaseQuery) Offset(offset int) *DatabaseQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DatabaseQuery) Unique(unique bool) *DatabaseQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DatabaseQuery) Order(o ...database.OrderOption) *DatabaseQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryService chains the current query on the "service" edge.
func (dq *DatabaseQuery) QueryService() *ServiceQuery {
	query := (&ServiceClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(database.Table, database.FieldID, selector),
			sqlgraph.To(service.Table, service.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, database.ServiceTable, database.ServiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGeneralspec chains the current query on the "generalspec" edge.
func (dq *DatabaseQuery) QueryGeneralspec() *GeneralSpecQuery {
	query := (&GeneralSpecClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(database.Table, database.FieldID, selector),
			sqlgraph.To(generalspec.Table, generalspec.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, database.GeneralspecTable, database.GeneralspecColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Database entity from the query.
// Returns a *NotFoundError when no Database was found.
func (dq *DatabaseQuery) First(ctx context.Context) (*Database, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{database.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DatabaseQuery) FirstX(ctx context.Context) *Database {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Database ID from the query.
// Returns a *NotFoundError when no Database ID was found.
func (dq *DatabaseQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{database.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DatabaseQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Database entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Database entity is found.
// Returns a *NotFoundError when no Database entities are found.
func (dq *DatabaseQuery) Only(ctx context.Context) (*Database, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{database.Label}
	default:
		return nil, &NotSingularError{database.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DatabaseQuery) OnlyX(ctx context.Context) *Database {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Database ID in the query.
// Returns a *NotSingularError when more than one Database ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DatabaseQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{database.Label}
	default:
		err = &NotSingularError{database.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DatabaseQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Databases.
func (dq *DatabaseQuery) All(ctx context.Context) ([]*Database, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryAll)
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Database, *DatabaseQuery]()
	return withInterceptors[[]*Database](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DatabaseQuery) AllX(ctx context.Context) []*Database {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Database IDs.
func (dq *DatabaseQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryIDs)
	if err = dq.Select(database.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DatabaseQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DatabaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryCount)
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DatabaseQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DatabaseQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DatabaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryExist)
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DatabaseQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DatabaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DatabaseQuery) Clone() *DatabaseQuery {
	if dq == nil {
		return nil
	}
	return &DatabaseQuery{
		config:          dq.config,
		ctx:             dq.ctx.Clone(),
		order:           append([]database.OrderOption{}, dq.order...),
		inters:          append([]Interceptor{}, dq.inters...),
		predicates:      append([]predicate.Database{}, dq.predicates...),
		withService:     dq.withService.Clone(),
		withGeneralspec: dq.withGeneralspec.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithService tells the query-builder to eager-load the nodes that are connected to
// the "service" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DatabaseQuery) WithService(opts ...func(*ServiceQuery)) *DatabaseQuery {
	query := (&ServiceClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withService = query
	return dq
}

// WithGeneralspec tells the query-builder to eager-load the nodes that are connected to
// the "generalspec" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DatabaseQuery) WithGeneralspec(opts ...func(*GeneralSpecQuery)) *DatabaseQuery {
	query := (&GeneralSpecClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withGeneralspec = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ConnectionPath string `json:"connection_path,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Database.Query().
//		GroupBy(database.FieldConnectionPath).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DatabaseQuery) GroupBy(field string, fields ...string) *DatabaseGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DatabaseGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = database.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ConnectionPath string `json:"connection_path,omitempty"`
//	}
//
//	client.Database.Query().
//		Select(database.FieldConnectionPath).
//		Scan(ctx, &v)
func (dq *DatabaseQuery) Select(fields ...string) *DatabaseSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DatabaseSelect{DatabaseQuery: dq}
	sbuild.label = database.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DatabaseSelect configured with the given aggregations.
func (dq *DatabaseQuery) Aggregate(fns ...AggregateFunc) *DatabaseSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DatabaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !database.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DatabaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Database, error) {
	var (
		nodes       = []*Database{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [2]bool{
			dq.withService != nil,
			dq.withGeneralspec != nil,
		}
	)
	if dq.withService != nil || dq.withGeneralspec != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, database.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Database).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Database{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withService; query != nil {
		if err := dq.loadService(ctx, query, nodes, nil,
			func(n *Database, e *Service) { n.Edges.Service = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withGeneralspec; query != nil {
		if err := dq.loadGeneralspec(ctx, query, nodes, nil,
			func(n *Database, e *GeneralSpec) { n.Edges.Generalspec = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DatabaseQuery) loadService(ctx context.Context, query *ServiceQuery, nodes []*Database, init func(*Database), assign func(*Database, *Service)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Database)
	for i := range nodes {
		if nodes[i].service_databases == nil {
			continue
		}
		fk := *nodes[i].service_databases
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(service.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "service_databases" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DatabaseQuery) loadGeneralspec(ctx context.Context, query *GeneralSpecQuery, nodes []*Database, init func(*Database), assign func(*Database, *GeneralSpec)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Database)
	for i := range nodes {
		if nodes[i].general_spec_database == nil {
			continue
		}
		fk := *nodes[i].general_spec_database
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(generalspec.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "general_spec_database" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DatabaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DatabaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(database.Table, database.Columns, sqlgraph.NewFieldSpec(database.FieldID, field.TypeUUID))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, database.FieldID)
		for i := range fields {
			if fields[i] != database.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DatabaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(database.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = database.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DatabaseGroupBy is the group-by builder for Database entities.
type DatabaseGroupBy struct {
	selector
	build *DatabaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DatabaseGroupBy) Aggregate(fns ...AggregateFunc) *DatabaseGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DatabaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, ent.OpQueryGroupBy)
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DatabaseQuery, *DatabaseGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DatabaseGroupBy) sqlScan(ctx context.Context, root *DatabaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DatabaseSelect is the builder for selecting fields of Database entities.
type DatabaseSelect struct {
	*DatabaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DatabaseSelect) Aggregate(fns ...AggregateFunc) *DatabaseSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DatabaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, ent.OpQuerySelect)
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DatabaseQuery, *DatabaseSelect](ctx, ds.DatabaseQuery, ds, ds.inters, v)
}

func (ds *DatabaseSelect) sqlScan(ctx context.Context, root *DatabaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
