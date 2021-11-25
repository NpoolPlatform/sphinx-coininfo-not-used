// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/sphinx-coininfo/pkg/db/ent/walletnode"
	"github.com/google/uuid"
)

// WalletNodeUpdate is the builder for updating WalletNode entities.
type WalletNodeUpdate struct {
	config
	hooks    []Hook
	mutation *WalletNodeMutation
}

// Where appends a list predicates to the WalletNodeUpdate builder.
func (wnu *WalletNodeUpdate) Where(ps ...predicate.WalletNode) *WalletNodeUpdate {
	wnu.mutation.Where(ps...)
	return wnu
}

// SetUUID sets the "uuid" field.
func (wnu *WalletNodeUpdate) SetUUID(s string) *WalletNodeUpdate {
	wnu.mutation.SetUUID(s)
	return wnu
}

// SetLocation sets the "location" field.
func (wnu *WalletNodeUpdate) SetLocation(s string) *WalletNodeUpdate {
	wnu.mutation.SetLocation(s)
	return wnu
}

// SetHostVendor sets the "host_vendor" field.
func (wnu *WalletNodeUpdate) SetHostVendor(s string) *WalletNodeUpdate {
	wnu.mutation.SetHostVendor(s)
	return wnu
}

// SetPublicIP sets the "public_ip" field.
func (wnu *WalletNodeUpdate) SetPublicIP(s string) *WalletNodeUpdate {
	wnu.mutation.SetPublicIP(s)
	return wnu
}

// SetLocalIP sets the "local_ip" field.
func (wnu *WalletNodeUpdate) SetLocalIP(s string) *WalletNodeUpdate {
	wnu.mutation.SetLocalIP(s)
	return wnu
}

// SetCreatetimeUtc sets the "createtime_utc" field.
func (wnu *WalletNodeUpdate) SetCreatetimeUtc(i int64) *WalletNodeUpdate {
	wnu.mutation.ResetCreatetimeUtc()
	wnu.mutation.SetCreatetimeUtc(i)
	return wnu
}

// AddCreatetimeUtc adds i to the "createtime_utc" field.
func (wnu *WalletNodeUpdate) AddCreatetimeUtc(i int64) *WalletNodeUpdate {
	wnu.mutation.AddCreatetimeUtc(i)
	return wnu
}

// SetLastOnlineTimeUtc sets the "last_online_time_utc" field.
func (wnu *WalletNodeUpdate) SetLastOnlineTimeUtc(i int64) *WalletNodeUpdate {
	wnu.mutation.ResetLastOnlineTimeUtc()
	wnu.mutation.SetLastOnlineTimeUtc(i)
	return wnu
}

// AddLastOnlineTimeUtc adds i to the "last_online_time_utc" field.
func (wnu *WalletNodeUpdate) AddLastOnlineTimeUtc(i int64) *WalletNodeUpdate {
	wnu.mutation.AddLastOnlineTimeUtc(i)
	return wnu
}

// AddCoinIDs adds the "coin" edge to the CoinInfo entity by IDs.
func (wnu *WalletNodeUpdate) AddCoinIDs(ids ...uuid.UUID) *WalletNodeUpdate {
	wnu.mutation.AddCoinIDs(ids...)
	return wnu
}

// AddCoin adds the "coin" edges to the CoinInfo entity.
func (wnu *WalletNodeUpdate) AddCoin(c ...*CoinInfo) *WalletNodeUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wnu.AddCoinIDs(ids...)
}

// Mutation returns the WalletNodeMutation object of the builder.
func (wnu *WalletNodeUpdate) Mutation() *WalletNodeMutation {
	return wnu.mutation
}

// ClearCoin clears all "coin" edges to the CoinInfo entity.
func (wnu *WalletNodeUpdate) ClearCoin() *WalletNodeUpdate {
	wnu.mutation.ClearCoin()
	return wnu
}

// RemoveCoinIDs removes the "coin" edge to CoinInfo entities by IDs.
func (wnu *WalletNodeUpdate) RemoveCoinIDs(ids ...uuid.UUID) *WalletNodeUpdate {
	wnu.mutation.RemoveCoinIDs(ids...)
	return wnu
}

// RemoveCoin removes "coin" edges to CoinInfo entities.
func (wnu *WalletNodeUpdate) RemoveCoin(c ...*CoinInfo) *WalletNodeUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wnu.RemoveCoinIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wnu *WalletNodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wnu.hooks) == 0 {
		affected, err = wnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wnu.mutation = mutation
			affected, err = wnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wnu.hooks) - 1; i >= 0; i-- {
			if wnu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wnu *WalletNodeUpdate) SaveX(ctx context.Context) int {
	affected, err := wnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wnu *WalletNodeUpdate) Exec(ctx context.Context) error {
	_, err := wnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnu *WalletNodeUpdate) ExecX(ctx context.Context) {
	if err := wnu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wnu *WalletNodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   walletnode.Table,
			Columns: walletnode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: walletnode.FieldID,
			},
		},
	}
	if ps := wnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wnu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldUUID,
		})
	}
	if value, ok := wnu.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldLocation,
		})
	}
	if value, ok := wnu.mutation.HostVendor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldHostVendor,
		})
	}
	if value, ok := wnu.mutation.PublicIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldPublicIP,
		})
	}
	if value, ok := wnu.mutation.LocalIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldLocalIP,
		})
	}
	if value, ok := wnu.mutation.CreatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: walletnode.FieldCreatetimeUtc,
		})
	}
	if value, ok := wnu.mutation.AddedCreatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: walletnode.FieldCreatetimeUtc,
		})
	}
	if value, ok := wnu.mutation.LastOnlineTimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: walletnode.FieldLastOnlineTimeUtc,
		})
	}
	if value, ok := wnu.mutation.AddedLastOnlineTimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: walletnode.FieldLastOnlineTimeUtc,
		})
	}
	if wnu.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   walletnode.CoinTable,
			Columns: walletnode.CoinPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnu.mutation.RemovedCoinIDs(); len(nodes) > 0 && !wnu.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   walletnode.CoinTable,
			Columns: walletnode.CoinPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnu.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   walletnode.CoinTable,
			Columns: walletnode.CoinPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{walletnode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WalletNodeUpdateOne is the builder for updating a single WalletNode entity.
type WalletNodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WalletNodeMutation
}

// SetUUID sets the "uuid" field.
func (wnuo *WalletNodeUpdateOne) SetUUID(s string) *WalletNodeUpdateOne {
	wnuo.mutation.SetUUID(s)
	return wnuo
}

// SetLocation sets the "location" field.
func (wnuo *WalletNodeUpdateOne) SetLocation(s string) *WalletNodeUpdateOne {
	wnuo.mutation.SetLocation(s)
	return wnuo
}

// SetHostVendor sets the "host_vendor" field.
func (wnuo *WalletNodeUpdateOne) SetHostVendor(s string) *WalletNodeUpdateOne {
	wnuo.mutation.SetHostVendor(s)
	return wnuo
}

// SetPublicIP sets the "public_ip" field.
func (wnuo *WalletNodeUpdateOne) SetPublicIP(s string) *WalletNodeUpdateOne {
	wnuo.mutation.SetPublicIP(s)
	return wnuo
}

// SetLocalIP sets the "local_ip" field.
func (wnuo *WalletNodeUpdateOne) SetLocalIP(s string) *WalletNodeUpdateOne {
	wnuo.mutation.SetLocalIP(s)
	return wnuo
}

// SetCreatetimeUtc sets the "createtime_utc" field.
func (wnuo *WalletNodeUpdateOne) SetCreatetimeUtc(i int64) *WalletNodeUpdateOne {
	wnuo.mutation.ResetCreatetimeUtc()
	wnuo.mutation.SetCreatetimeUtc(i)
	return wnuo
}

// AddCreatetimeUtc adds i to the "createtime_utc" field.
func (wnuo *WalletNodeUpdateOne) AddCreatetimeUtc(i int64) *WalletNodeUpdateOne {
	wnuo.mutation.AddCreatetimeUtc(i)
	return wnuo
}

// SetLastOnlineTimeUtc sets the "last_online_time_utc" field.
func (wnuo *WalletNodeUpdateOne) SetLastOnlineTimeUtc(i int64) *WalletNodeUpdateOne {
	wnuo.mutation.ResetLastOnlineTimeUtc()
	wnuo.mutation.SetLastOnlineTimeUtc(i)
	return wnuo
}

// AddLastOnlineTimeUtc adds i to the "last_online_time_utc" field.
func (wnuo *WalletNodeUpdateOne) AddLastOnlineTimeUtc(i int64) *WalletNodeUpdateOne {
	wnuo.mutation.AddLastOnlineTimeUtc(i)
	return wnuo
}

// AddCoinIDs adds the "coin" edge to the CoinInfo entity by IDs.
func (wnuo *WalletNodeUpdateOne) AddCoinIDs(ids ...uuid.UUID) *WalletNodeUpdateOne {
	wnuo.mutation.AddCoinIDs(ids...)
	return wnuo
}

// AddCoin adds the "coin" edges to the CoinInfo entity.
func (wnuo *WalletNodeUpdateOne) AddCoin(c ...*CoinInfo) *WalletNodeUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wnuo.AddCoinIDs(ids...)
}

// Mutation returns the WalletNodeMutation object of the builder.
func (wnuo *WalletNodeUpdateOne) Mutation() *WalletNodeMutation {
	return wnuo.mutation
}

// ClearCoin clears all "coin" edges to the CoinInfo entity.
func (wnuo *WalletNodeUpdateOne) ClearCoin() *WalletNodeUpdateOne {
	wnuo.mutation.ClearCoin()
	return wnuo
}

// RemoveCoinIDs removes the "coin" edge to CoinInfo entities by IDs.
func (wnuo *WalletNodeUpdateOne) RemoveCoinIDs(ids ...uuid.UUID) *WalletNodeUpdateOne {
	wnuo.mutation.RemoveCoinIDs(ids...)
	return wnuo
}

// RemoveCoin removes "coin" edges to CoinInfo entities.
func (wnuo *WalletNodeUpdateOne) RemoveCoin(c ...*CoinInfo) *WalletNodeUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wnuo.RemoveCoinIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wnuo *WalletNodeUpdateOne) Select(field string, fields ...string) *WalletNodeUpdateOne {
	wnuo.fields = append([]string{field}, fields...)
	return wnuo
}

// Save executes the query and returns the updated WalletNode entity.
func (wnuo *WalletNodeUpdateOne) Save(ctx context.Context) (*WalletNode, error) {
	var (
		err  error
		node *WalletNode
	)
	if len(wnuo.hooks) == 0 {
		node, err = wnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wnuo.mutation = mutation
			node, err = wnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wnuo.hooks) - 1; i >= 0; i-- {
			if wnuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wnuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wnuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wnuo *WalletNodeUpdateOne) SaveX(ctx context.Context) *WalletNode {
	node, err := wnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wnuo *WalletNodeUpdateOne) Exec(ctx context.Context) error {
	_, err := wnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnuo *WalletNodeUpdateOne) ExecX(ctx context.Context) {
	if err := wnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wnuo *WalletNodeUpdateOne) sqlSave(ctx context.Context) (_node *WalletNode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   walletnode.Table,
			Columns: walletnode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: walletnode.FieldID,
			},
		},
	}
	id, ok := wnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WalletNode.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, walletnode.FieldID)
		for _, f := range fields {
			if !walletnode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != walletnode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wnuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldUUID,
		})
	}
	if value, ok := wnuo.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldLocation,
		})
	}
	if value, ok := wnuo.mutation.HostVendor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldHostVendor,
		})
	}
	if value, ok := wnuo.mutation.PublicIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldPublicIP,
		})
	}
	if value, ok := wnuo.mutation.LocalIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: walletnode.FieldLocalIP,
		})
	}
	if value, ok := wnuo.mutation.CreatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: walletnode.FieldCreatetimeUtc,
		})
	}
	if value, ok := wnuo.mutation.AddedCreatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: walletnode.FieldCreatetimeUtc,
		})
	}
	if value, ok := wnuo.mutation.LastOnlineTimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: walletnode.FieldLastOnlineTimeUtc,
		})
	}
	if value, ok := wnuo.mutation.AddedLastOnlineTimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: walletnode.FieldLastOnlineTimeUtc,
		})
	}
	if wnuo.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   walletnode.CoinTable,
			Columns: walletnode.CoinPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnuo.mutation.RemovedCoinIDs(); len(nodes) > 0 && !wnuo.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   walletnode.CoinTable,
			Columns: walletnode.CoinPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnuo.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   walletnode.CoinTable,
			Columns: walletnode.CoinPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WalletNode{config: wnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{walletnode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
