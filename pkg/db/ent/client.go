// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/empty"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/walletnode"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CoinInfo is the client for interacting with the CoinInfo builders.
	CoinInfo *CoinInfoClient
	// Empty is the client for interacting with the Empty builders.
	Empty *EmptyClient
	// Review is the client for interacting with the Review builders.
	Review *ReviewClient
	// Transaction is the client for interacting with the Transaction builders.
	Transaction *TransactionClient
	// WalletNode is the client for interacting with the WalletNode builders.
	WalletNode *WalletNodeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CoinInfo = NewCoinInfoClient(c.config)
	c.Empty = NewEmptyClient(c.config)
	c.Review = NewReviewClient(c.config)
	c.Transaction = NewTransactionClient(c.config)
	c.WalletNode = NewWalletNodeClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		CoinInfo:    NewCoinInfoClient(cfg),
		Empty:       NewEmptyClient(cfg),
		Review:      NewReviewClient(cfg),
		Transaction: NewTransactionClient(cfg),
		WalletNode:  NewWalletNodeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:      cfg,
		CoinInfo:    NewCoinInfoClient(cfg),
		Empty:       NewEmptyClient(cfg),
		Review:      NewReviewClient(cfg),
		Transaction: NewTransactionClient(cfg),
		WalletNode:  NewWalletNodeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CoinInfo.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CoinInfo.Use(hooks...)
	c.Empty.Use(hooks...)
	c.Review.Use(hooks...)
	c.Transaction.Use(hooks...)
	c.WalletNode.Use(hooks...)
}

// CoinInfoClient is a client for the CoinInfo schema.
type CoinInfoClient struct {
	config
}

// NewCoinInfoClient returns a client for the CoinInfo from the given config.
func NewCoinInfoClient(c config) *CoinInfoClient {
	return &CoinInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coininfo.Hooks(f(g(h())))`.
func (c *CoinInfoClient) Use(hooks ...Hook) {
	c.hooks.CoinInfo = append(c.hooks.CoinInfo, hooks...)
}

// Create returns a create builder for CoinInfo.
func (c *CoinInfoClient) Create() *CoinInfoCreate {
	mutation := newCoinInfoMutation(c.config, OpCreate)
	return &CoinInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CoinInfo entities.
func (c *CoinInfoClient) CreateBulk(builders ...*CoinInfoCreate) *CoinInfoCreateBulk {
	return &CoinInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CoinInfo.
func (c *CoinInfoClient) Update() *CoinInfoUpdate {
	mutation := newCoinInfoMutation(c.config, OpUpdate)
	return &CoinInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CoinInfoClient) UpdateOne(ci *CoinInfo) *CoinInfoUpdateOne {
	mutation := newCoinInfoMutation(c.config, OpUpdateOne, withCoinInfo(ci))
	return &CoinInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CoinInfoClient) UpdateOneID(id uuid.UUID) *CoinInfoUpdateOne {
	mutation := newCoinInfoMutation(c.config, OpUpdateOne, withCoinInfoID(id))
	return &CoinInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CoinInfo.
func (c *CoinInfoClient) Delete() *CoinInfoDelete {
	mutation := newCoinInfoMutation(c.config, OpDelete)
	return &CoinInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CoinInfoClient) DeleteOne(ci *CoinInfo) *CoinInfoDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CoinInfoClient) DeleteOneID(id uuid.UUID) *CoinInfoDeleteOne {
	builder := c.Delete().Where(coininfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CoinInfoDeleteOne{builder}
}

// Query returns a query builder for CoinInfo.
func (c *CoinInfoClient) Query() *CoinInfoQuery {
	return &CoinInfoQuery{
		config: c.config,
	}
}

// Get returns a CoinInfo entity by its id.
func (c *CoinInfoClient) Get(ctx context.Context, id uuid.UUID) (*CoinInfo, error) {
	return c.Query().Where(coininfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CoinInfoClient) GetX(ctx context.Context, id uuid.UUID) *CoinInfo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTransactions queries the transactions edge of a CoinInfo.
func (c *CoinInfoClient) QueryTransactions(ci *CoinInfo) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(coininfo.Table, coininfo.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, coininfo.TransactionsTable, coininfo.TransactionsColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReviews queries the reviews edge of a CoinInfo.
func (c *CoinInfoClient) QueryReviews(ci *CoinInfo) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(coininfo.Table, coininfo.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, coininfo.ReviewsTable, coininfo.ReviewsColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWalletNodes queries the wallet_nodes edge of a CoinInfo.
func (c *CoinInfoClient) QueryWalletNodes(ci *CoinInfo) *WalletNodeQuery {
	query := &WalletNodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(coininfo.Table, coininfo.FieldID, id),
			sqlgraph.To(walletnode.Table, walletnode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, coininfo.WalletNodesTable, coininfo.WalletNodesColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CoinInfoClient) Hooks() []Hook {
	return c.hooks.CoinInfo
}

// EmptyClient is a client for the Empty schema.
type EmptyClient struct {
	config
}

// NewEmptyClient returns a client for the Empty from the given config.
func NewEmptyClient(c config) *EmptyClient {
	return &EmptyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `empty.Hooks(f(g(h())))`.
func (c *EmptyClient) Use(hooks ...Hook) {
	c.hooks.Empty = append(c.hooks.Empty, hooks...)
}

// Create returns a create builder for Empty.
func (c *EmptyClient) Create() *EmptyCreate {
	mutation := newEmptyMutation(c.config, OpCreate)
	return &EmptyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Empty entities.
func (c *EmptyClient) CreateBulk(builders ...*EmptyCreate) *EmptyCreateBulk {
	return &EmptyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Empty.
func (c *EmptyClient) Update() *EmptyUpdate {
	mutation := newEmptyMutation(c.config, OpUpdate)
	return &EmptyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmptyClient) UpdateOne(e *Empty) *EmptyUpdateOne {
	mutation := newEmptyMutation(c.config, OpUpdateOne, withEmpty(e))
	return &EmptyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmptyClient) UpdateOneID(id int) *EmptyUpdateOne {
	mutation := newEmptyMutation(c.config, OpUpdateOne, withEmptyID(id))
	return &EmptyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Empty.
func (c *EmptyClient) Delete() *EmptyDelete {
	mutation := newEmptyMutation(c.config, OpDelete)
	return &EmptyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EmptyClient) DeleteOne(e *Empty) *EmptyDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EmptyClient) DeleteOneID(id int) *EmptyDeleteOne {
	builder := c.Delete().Where(empty.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmptyDeleteOne{builder}
}

// Query returns a query builder for Empty.
func (c *EmptyClient) Query() *EmptyQuery {
	return &EmptyQuery{
		config: c.config,
	}
}

// Get returns a Empty entity by its id.
func (c *EmptyClient) Get(ctx context.Context, id int) (*Empty, error) {
	return c.Query().Where(empty.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmptyClient) GetX(ctx context.Context, id int) *Empty {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EmptyClient) Hooks() []Hook {
	return c.hooks.Empty
}

// ReviewClient is a client for the Review schema.
type ReviewClient struct {
	config
}

// NewReviewClient returns a client for the Review from the given config.
func NewReviewClient(c config) *ReviewClient {
	return &ReviewClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `review.Hooks(f(g(h())))`.
func (c *ReviewClient) Use(hooks ...Hook) {
	c.hooks.Review = append(c.hooks.Review, hooks...)
}

// Create returns a create builder for Review.
func (c *ReviewClient) Create() *ReviewCreate {
	mutation := newReviewMutation(c.config, OpCreate)
	return &ReviewCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Review entities.
func (c *ReviewClient) CreateBulk(builders ...*ReviewCreate) *ReviewCreateBulk {
	return &ReviewCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Review.
func (c *ReviewClient) Update() *ReviewUpdate {
	mutation := newReviewMutation(c.config, OpUpdate)
	return &ReviewUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReviewClient) UpdateOne(r *Review) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReview(r))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReviewClient) UpdateOneID(id int32) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReviewID(id))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Review.
func (c *ReviewClient) Delete() *ReviewDelete {
	mutation := newReviewMutation(c.config, OpDelete)
	return &ReviewDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ReviewClient) DeleteOne(r *Review) *ReviewDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ReviewClient) DeleteOneID(id int32) *ReviewDeleteOne {
	builder := c.Delete().Where(review.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReviewDeleteOne{builder}
}

// Query returns a query builder for Review.
func (c *ReviewClient) Query() *ReviewQuery {
	return &ReviewQuery{
		config: c.config,
	}
}

// Get returns a Review entity by its id.
func (c *ReviewClient) Get(ctx context.Context, id int32) (*Review, error) {
	return c.Query().Where(review.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReviewClient) GetX(ctx context.Context, id int32) *Review {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTransaction queries the transaction edge of a Review.
func (c *ReviewClient) QueryTransaction(r *Review) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, review.TransactionTable, review.TransactionColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCoin queries the coin edge of a Review.
func (c *ReviewClient) QueryCoin(r *Review) *CoinInfoQuery {
	query := &CoinInfoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(coininfo.Table, coininfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, review.CoinTable, review.CoinColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReviewClient) Hooks() []Hook {
	return c.hooks.Review
}

// TransactionClient is a client for the Transaction schema.
type TransactionClient struct {
	config
}

// NewTransactionClient returns a client for the Transaction from the given config.
func NewTransactionClient(c config) *TransactionClient {
	return &TransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transaction.Hooks(f(g(h())))`.
func (c *TransactionClient) Use(hooks ...Hook) {
	c.hooks.Transaction = append(c.hooks.Transaction, hooks...)
}

// Create returns a create builder for Transaction.
func (c *TransactionClient) Create() *TransactionCreate {
	mutation := newTransactionMutation(c.config, OpCreate)
	return &TransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transaction entities.
func (c *TransactionClient) CreateBulk(builders ...*TransactionCreate) *TransactionCreateBulk {
	return &TransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transaction.
func (c *TransactionClient) Update() *TransactionUpdate {
	mutation := newTransactionMutation(c.config, OpUpdate)
	return &TransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransactionClient) UpdateOne(t *Transaction) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransaction(t))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransactionClient) UpdateOneID(id int32) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransactionID(id))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transaction.
func (c *TransactionClient) Delete() *TransactionDelete {
	mutation := newTransactionMutation(c.config, OpDelete)
	return &TransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TransactionClient) DeleteOne(t *Transaction) *TransactionDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TransactionClient) DeleteOneID(id int32) *TransactionDeleteOne {
	builder := c.Delete().Where(transaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransactionDeleteOne{builder}
}

// Query returns a query builder for Transaction.
func (c *TransactionClient) Query() *TransactionQuery {
	return &TransactionQuery{
		config: c.config,
	}
}

// Get returns a Transaction entity by its id.
func (c *TransactionClient) Get(ctx context.Context, id int32) (*Transaction, error) {
	return c.Query().Where(transaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransactionClient) GetX(ctx context.Context, id int32) *Transaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCoin queries the coin edge of a Transaction.
func (c *TransactionClient) QueryCoin(t *Transaction) *CoinInfoQuery {
	query := &CoinInfoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(coininfo.Table, coininfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transaction.CoinTable, transaction.CoinColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReview queries the review edge of a Transaction.
func (c *TransactionClient) QueryReview(t *Transaction) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transaction.ReviewTable, transaction.ReviewColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TransactionClient) Hooks() []Hook {
	return c.hooks.Transaction
}

// WalletNodeClient is a client for the WalletNode schema.
type WalletNodeClient struct {
	config
}

// NewWalletNodeClient returns a client for the WalletNode from the given config.
func NewWalletNodeClient(c config) *WalletNodeClient {
	return &WalletNodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `walletnode.Hooks(f(g(h())))`.
func (c *WalletNodeClient) Use(hooks ...Hook) {
	c.hooks.WalletNode = append(c.hooks.WalletNode, hooks...)
}

// Create returns a create builder for WalletNode.
func (c *WalletNodeClient) Create() *WalletNodeCreate {
	mutation := newWalletNodeMutation(c.config, OpCreate)
	return &WalletNodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WalletNode entities.
func (c *WalletNodeClient) CreateBulk(builders ...*WalletNodeCreate) *WalletNodeCreateBulk {
	return &WalletNodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WalletNode.
func (c *WalletNodeClient) Update() *WalletNodeUpdate {
	mutation := newWalletNodeMutation(c.config, OpUpdate)
	return &WalletNodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WalletNodeClient) UpdateOne(wn *WalletNode) *WalletNodeUpdateOne {
	mutation := newWalletNodeMutation(c.config, OpUpdateOne, withWalletNode(wn))
	return &WalletNodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WalletNodeClient) UpdateOneID(id int32) *WalletNodeUpdateOne {
	mutation := newWalletNodeMutation(c.config, OpUpdateOne, withWalletNodeID(id))
	return &WalletNodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WalletNode.
func (c *WalletNodeClient) Delete() *WalletNodeDelete {
	mutation := newWalletNodeMutation(c.config, OpDelete)
	return &WalletNodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WalletNodeClient) DeleteOne(wn *WalletNode) *WalletNodeDeleteOne {
	return c.DeleteOneID(wn.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WalletNodeClient) DeleteOneID(id int32) *WalletNodeDeleteOne {
	builder := c.Delete().Where(walletnode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WalletNodeDeleteOne{builder}
}

// Query returns a query builder for WalletNode.
func (c *WalletNodeClient) Query() *WalletNodeQuery {
	return &WalletNodeQuery{
		config: c.config,
	}
}

// Get returns a WalletNode entity by its id.
func (c *WalletNodeClient) Get(ctx context.Context, id int32) (*WalletNode, error) {
	return c.Query().Where(walletnode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WalletNodeClient) GetX(ctx context.Context, id int32) *WalletNode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCoin queries the coin edge of a WalletNode.
func (c *WalletNodeClient) QueryCoin(wn *WalletNode) *CoinInfoQuery {
	query := &CoinInfoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := wn.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(walletnode.Table, walletnode.FieldID, id),
			sqlgraph.To(coininfo.Table, coininfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, walletnode.CoinTable, walletnode.CoinColumn),
		)
		fromV = sqlgraph.Neighbors(wn.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WalletNodeClient) Hooks() []Hook {
	return c.hooks.WalletNode
}
