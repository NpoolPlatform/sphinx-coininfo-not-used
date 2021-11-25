// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/walletnode"
	"github.com/google/uuid"
)

// CoinInfoQuery is the builder for querying CoinInfo entities.
type CoinInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CoinInfo
	// eager-loading edges.
	withTransactions *TransactionQuery
	withReviews      *ReviewQuery
	withWalletNodes  *WalletNodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoinInfoQuery builder.
func (ciq *CoinInfoQuery) Where(ps ...predicate.CoinInfo) *CoinInfoQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit adds a limit step to the query.
func (ciq *CoinInfoQuery) Limit(limit int) *CoinInfoQuery {
	ciq.limit = &limit
	return ciq
}

// Offset adds an offset step to the query.
func (ciq *CoinInfoQuery) Offset(offset int) *CoinInfoQuery {
	ciq.offset = &offset
	return ciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciq *CoinInfoQuery) Unique(unique bool) *CoinInfoQuery {
	ciq.unique = &unique
	return ciq
}

// Order adds an order step to the query.
func (ciq *CoinInfoQuery) Order(o ...OrderFunc) *CoinInfoQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// QueryTransactions chains the current query on the "transactions" edge.
func (ciq *CoinInfoQuery) QueryTransactions() *TransactionQuery {
	query := &TransactionQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(coininfo.Table, coininfo.FieldID, selector),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, coininfo.TransactionsTable, coininfo.TransactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReviews chains the current query on the "reviews" edge.
func (ciq *CoinInfoQuery) QueryReviews() *ReviewQuery {
	query := &ReviewQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(coininfo.Table, coininfo.FieldID, selector),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, coininfo.ReviewsTable, coininfo.ReviewsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWalletNodes chains the current query on the "wallet_nodes" edge.
func (ciq *CoinInfoQuery) QueryWalletNodes() *WalletNodeQuery {
	query := &WalletNodeQuery{config: ciq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(coininfo.Table, coininfo.FieldID, selector),
			sqlgraph.To(walletnode.Table, walletnode.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, coininfo.WalletNodesTable, coininfo.WalletNodesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CoinInfo entity from the query.
// Returns a *NotFoundError when no CoinInfo was found.
func (ciq *CoinInfoQuery) First(ctx context.Context) (*CoinInfo, error) {
	nodes, err := ciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coininfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CoinInfoQuery) FirstX(ctx context.Context) *CoinInfo {
	node, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CoinInfo ID from the query.
// Returns a *NotFoundError when no CoinInfo ID was found.
func (ciq *CoinInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coininfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciq *CoinInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CoinInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CoinInfo entity is not found.
// Returns a *NotFoundError when no CoinInfo entities are found.
func (ciq *CoinInfoQuery) Only(ctx context.Context) (*CoinInfo, error) {
	nodes, err := ciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coininfo.Label}
	default:
		return nil, &NotSingularError{coininfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CoinInfoQuery) OnlyX(ctx context.Context) *CoinInfo {
	node, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CoinInfo ID in the query.
// Returns a *NotSingularError when exactly one CoinInfo ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ciq *CoinInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = &NotSingularError{coininfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CoinInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CoinInfos.
func (ciq *CoinInfoQuery) All(ctx context.Context) ([]*CoinInfo, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CoinInfoQuery) AllX(ctx context.Context) []*CoinInfo {
	nodes, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CoinInfo IDs.
func (ciq *CoinInfoQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ciq.Select(coininfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CoinInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CoinInfoQuery) Count(ctx context.Context) (int, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CoinInfoQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CoinInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CoinInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoinInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CoinInfoQuery) Clone() *CoinInfoQuery {
	if ciq == nil {
		return nil
	}
	return &CoinInfoQuery{
		config:           ciq.config,
		limit:            ciq.limit,
		offset:           ciq.offset,
		order:            append([]OrderFunc{}, ciq.order...),
		predicates:       append([]predicate.CoinInfo{}, ciq.predicates...),
		withTransactions: ciq.withTransactions.Clone(),
		withReviews:      ciq.withReviews.Clone(),
		withWalletNodes:  ciq.withWalletNodes.Clone(),
		// clone intermediate query.
		sql:  ciq.sql.Clone(),
		path: ciq.path,
	}
}

// WithTransactions tells the query-builder to eager-load the nodes that are connected to
// the "transactions" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CoinInfoQuery) WithTransactions(opts ...func(*TransactionQuery)) *CoinInfoQuery {
	query := &TransactionQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withTransactions = query
	return ciq
}

// WithReviews tells the query-builder to eager-load the nodes that are connected to
// the "reviews" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CoinInfoQuery) WithReviews(opts ...func(*ReviewQuery)) *CoinInfoQuery {
	query := &ReviewQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withReviews = query
	return ciq
}

// WithWalletNodes tells the query-builder to eager-load the nodes that are connected to
// the "wallet_nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CoinInfoQuery) WithWalletNodes(opts ...func(*WalletNodeQuery)) *CoinInfoQuery {
	query := &WalletNodeQuery{config: ciq.config}
	for _, opt := range opts {
		opt(query)
	}
	ciq.withWalletNodes = query
	return ciq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CoinTypeID int32 `json:"coin_type_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CoinInfo.Query().
//		GroupBy(coininfo.FieldCoinTypeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ciq *CoinInfoQuery) GroupBy(field string, fields ...string) *CoinInfoGroupBy {
	group := &CoinInfoGroupBy{config: ciq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CoinTypeID int32 `json:"coin_type_id,omitempty"`
//	}
//
//	client.CoinInfo.Query().
//		Select(coininfo.FieldCoinTypeID).
//		Scan(ctx, &v)
//
func (ciq *CoinInfoQuery) Select(fields ...string) *CoinInfoSelect {
	ciq.fields = append(ciq.fields, fields...)
	return &CoinInfoSelect{CoinInfoQuery: ciq}
}

func (ciq *CoinInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ciq.fields {
		if !coininfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *CoinInfoQuery) sqlAll(ctx context.Context) ([]*CoinInfo, error) {
	var (
		nodes       = []*CoinInfo{}
		_spec       = ciq.querySpec()
		loadedTypes = [3]bool{
			ciq.withTransactions != nil,
			ciq.withReviews != nil,
			ciq.withWalletNodes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CoinInfo{config: ciq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ciq.withTransactions; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*CoinInfo)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Transactions = []*Transaction{}
		}
		query.withFKs = true
		query.Where(predicate.Transaction(func(s *sql.Selector) {
			s.Where(sql.InValues(coininfo.TransactionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.coin_info_transactions
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "coin_info_transactions" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coin_info_transactions" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Transactions = append(node.Edges.Transactions, n)
		}
	}

	if query := ciq.withReviews; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[uuid.UUID]*CoinInfo, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Reviews = []*Review{}
		}
		var (
			edgeids []int32
			edges   = make(map[int32][]*CoinInfo)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   coininfo.ReviewsTable,
				Columns: coininfo.ReviewsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(coininfo.ReviewsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(uuid.UUID), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*uuid.UUID)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := *eout
				inValue := int32(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, ciq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "reviews": %w`, err)
		}
		query.Where(review.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "reviews" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Reviews = append(nodes[i].Edges.Reviews, n)
			}
		}
	}

	if query := ciq.withWalletNodes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[uuid.UUID]*CoinInfo, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.WalletNodes = []*WalletNode{}
		}
		var (
			edgeids []int32
			edges   = make(map[int32][]*CoinInfo)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   coininfo.WalletNodesTable,
				Columns: coininfo.WalletNodesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(coininfo.WalletNodesPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(uuid.UUID), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*uuid.UUID)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := *eout
				inValue := int32(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, ciq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "wallet_nodes": %w`, err)
		}
		query.Where(walletnode.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "wallet_nodes" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.WalletNodes = append(nodes[i].Edges.WalletNodes, n)
			}
		}
	}

	return nodes, nil
}

func (ciq *CoinInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CoinInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ciq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ciq *CoinInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coininfo.Table,
			Columns: coininfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coininfo.FieldID,
			},
		},
		From:   ciq.sql,
		Unique: true,
	}
	if unique := ciq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ciq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coininfo.FieldID)
		for i := range fields {
			if fields[i] != coininfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CoinInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(coininfo.Table)
	columns := ciq.fields
	if len(columns) == 0 {
		columns = coininfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CoinInfoGroupBy is the group-by builder for CoinInfo entities.
type CoinInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CoinInfoGroupBy) Aggregate(fns ...AggregateFunc) *CoinInfoGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the group-by query and scans the result into the given value.
func (cigb *CoinInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cigb.path(ctx)
	if err != nil {
		return err
	}
	cigb.sql = query
	return cigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cigb *CoinInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CoinInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := cigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cigb *CoinInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = fmt.Errorf("ent: CoinInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) StringX(ctx context.Context) string {
	v, err := cigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cigb *CoinInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CoinInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := cigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cigb *CoinInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = fmt.Errorf("ent: CoinInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) IntX(ctx context.Context) int {
	v, err := cigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cigb *CoinInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CoinInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cigb *CoinInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = fmt.Errorf("ent: CoinInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cigb *CoinInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cigb.fields) > 1 {
		return nil, errors.New("ent: CoinInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cigb *CoinInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = fmt.Errorf("ent: CoinInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cigb *CoinInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := cigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cigb *CoinInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cigb.fields {
		if !coininfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cigb *CoinInfoGroupBy) sqlQuery() *sql.Selector {
	selector := cigb.sql.Select()
	aggregation := make([]string, 0, len(cigb.fns))
	for _, fn := range cigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cigb.fields)+len(cigb.fns))
		for _, f := range cigb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cigb.fields...)...)
}

// CoinInfoSelect is the builder for selecting fields of CoinInfo entities.
type CoinInfoSelect struct {
	*CoinInfoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cis *CoinInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cis.prepareQuery(ctx); err != nil {
		return err
	}
	cis.sql = cis.CoinInfoQuery.sqlQuery(ctx)
	return cis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cis *CoinInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cis *CoinInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CoinInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cis *CoinInfoSelect) StringsX(ctx context.Context) []string {
	v, err := cis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cis *CoinInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = fmt.Errorf("ent: CoinInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cis *CoinInfoSelect) StringX(ctx context.Context) string {
	v, err := cis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cis *CoinInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CoinInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cis *CoinInfoSelect) IntsX(ctx context.Context) []int {
	v, err := cis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cis *CoinInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = fmt.Errorf("ent: CoinInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cis *CoinInfoSelect) IntX(ctx context.Context) int {
	v, err := cis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cis *CoinInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CoinInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cis *CoinInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cis *CoinInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = fmt.Errorf("ent: CoinInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cis *CoinInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := cis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cis *CoinInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cis.fields) > 1 {
		return nil, errors.New("ent: CoinInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cis *CoinInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := cis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cis *CoinInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coininfo.Label}
	default:
		err = fmt.Errorf("ent: CoinInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cis *CoinInfoSelect) BoolX(ctx context.Context) bool {
	v, err := cis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cis *CoinInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cis.sql.Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
