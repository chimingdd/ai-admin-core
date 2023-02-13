// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
	"github.com/suyuan32/simple-admin-core/pkg/ent/position"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
)

// PositionUpdate is the builder for updating Position entities.
type PositionUpdate struct {
	config
	hooks    []Hook
	mutation *PositionMutation
}

// Where appends a list predicates to the PositionUpdate builder.
func (pu *PositionUpdate) Where(ps ...predicate.Position) *PositionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PositionUpdate) SetUpdatedAt(t time.Time) *PositionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PositionUpdate) SetStatus(u uint8) *PositionUpdate {
	pu.mutation.ResetStatus()
	pu.mutation.SetStatus(u)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableStatus(u *uint8) *PositionUpdate {
	if u != nil {
		pu.SetStatus(*u)
	}
	return pu
}

// AddStatus adds u to the "status" field.
func (pu *PositionUpdate) AddStatus(u int8) *PositionUpdate {
	pu.mutation.AddStatus(u)
	return pu
}

// ClearStatus clears the value of the "status" field.
func (pu *PositionUpdate) ClearStatus() *PositionUpdate {
	pu.mutation.ClearStatus()
	return pu
}

// SetSort sets the "sort" field.
func (pu *PositionUpdate) SetSort(u uint32) *PositionUpdate {
	pu.mutation.ResetSort()
	pu.mutation.SetSort(u)
	return pu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableSort(u *uint32) *PositionUpdate {
	if u != nil {
		pu.SetSort(*u)
	}
	return pu
}

// AddSort adds u to the "sort" field.
func (pu *PositionUpdate) AddSort(u int32) *PositionUpdate {
	pu.mutation.AddSort(u)
	return pu
}

// SetName sets the "name" field.
func (pu *PositionUpdate) SetName(s string) *PositionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetCode sets the "code" field.
func (pu *PositionUpdate) SetCode(s string) *PositionUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetRemark sets the "remark" field.
func (pu *PositionUpdate) SetRemark(s string) *PositionUpdate {
	pu.mutation.SetRemark(s)
	return pu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (pu *PositionUpdate) AddUserIDs(ids ...uuid.UUID) *PositionUpdate {
	pu.mutation.AddUserIDs(ids...)
	return pu
}

// AddUsers adds the "users" edges to the User entity.
func (pu *PositionUpdate) AddUsers(u ...*User) *PositionUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddUserIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pu *PositionUpdate) Mutation() *PositionMutation {
	return pu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (pu *PositionUpdate) ClearUsers() *PositionUpdate {
	pu.mutation.ClearUsers()
	return pu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (pu *PositionUpdate) RemoveUserIDs(ids ...uuid.UUID) *PositionUpdate {
	pu.mutation.RemoveUserIDs(ids...)
	return pu
}

// RemoveUsers removes "users" edges to User entities.
func (pu *PositionUpdate) RemoveUsers(u ...*User) *PositionUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PositionUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks[int, PositionMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PositionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PositionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PositionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PositionUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := position.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   position.Table,
			Columns: position.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: position.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(position.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.AddField(position.FieldStatus, field.TypeUint8, value)
	}
	if pu.mutation.StatusCleared() {
		_spec.ClearField(position.FieldStatus, field.TypeUint8)
	}
	if value, ok := pu.mutation.Sort(); ok {
		_spec.SetField(position.FieldSort, field.TypeUint32, value)
	}
	if value, ok := pu.mutation.AddedSort(); ok {
		_spec.AddField(position.FieldSort, field.TypeUint32, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
	}
	if pu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !pu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PositionUpdateOne is the builder for updating a single Position entity.
type PositionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PositionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PositionUpdateOne) SetUpdatedAt(t time.Time) *PositionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetStatus sets the "status" field.
func (puo *PositionUpdateOne) SetStatus(u uint8) *PositionUpdateOne {
	puo.mutation.ResetStatus()
	puo.mutation.SetStatus(u)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableStatus(u *uint8) *PositionUpdateOne {
	if u != nil {
		puo.SetStatus(*u)
	}
	return puo
}

// AddStatus adds u to the "status" field.
func (puo *PositionUpdateOne) AddStatus(u int8) *PositionUpdateOne {
	puo.mutation.AddStatus(u)
	return puo
}

// ClearStatus clears the value of the "status" field.
func (puo *PositionUpdateOne) ClearStatus() *PositionUpdateOne {
	puo.mutation.ClearStatus()
	return puo
}

// SetSort sets the "sort" field.
func (puo *PositionUpdateOne) SetSort(u uint32) *PositionUpdateOne {
	puo.mutation.ResetSort()
	puo.mutation.SetSort(u)
	return puo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableSort(u *uint32) *PositionUpdateOne {
	if u != nil {
		puo.SetSort(*u)
	}
	return puo
}

// AddSort adds u to the "sort" field.
func (puo *PositionUpdateOne) AddSort(u int32) *PositionUpdateOne {
	puo.mutation.AddSort(u)
	return puo
}

// SetName sets the "name" field.
func (puo *PositionUpdateOne) SetName(s string) *PositionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetCode sets the "code" field.
func (puo *PositionUpdateOne) SetCode(s string) *PositionUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetRemark sets the "remark" field.
func (puo *PositionUpdateOne) SetRemark(s string) *PositionUpdateOne {
	puo.mutation.SetRemark(s)
	return puo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (puo *PositionUpdateOne) AddUserIDs(ids ...uuid.UUID) *PositionUpdateOne {
	puo.mutation.AddUserIDs(ids...)
	return puo
}

// AddUsers adds the "users" edges to the User entity.
func (puo *PositionUpdateOne) AddUsers(u ...*User) *PositionUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddUserIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (puo *PositionUpdateOne) Mutation() *PositionMutation {
	return puo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (puo *PositionUpdateOne) ClearUsers() *PositionUpdateOne {
	puo.mutation.ClearUsers()
	return puo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (puo *PositionUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *PositionUpdateOne {
	puo.mutation.RemoveUserIDs(ids...)
	return puo
}

// RemoveUsers removes "users" edges to User entities.
func (puo *PositionUpdateOne) RemoveUsers(u ...*User) *PositionUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PositionUpdateOne) Select(field string, fields ...string) *PositionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Position entity.
func (puo *PositionUpdateOne) Save(ctx context.Context) (*Position, error) {
	puo.defaults()
	return withHooks[*Position, PositionMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PositionUpdateOne) SaveX(ctx context.Context) *Position {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PositionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PositionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PositionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := position.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PositionUpdateOne) sqlSave(ctx context.Context) (_node *Position, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   position.Table,
			Columns: position.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: position.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Position.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, position.FieldID)
		for _, f := range fields {
			if !position.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != position.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(position.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.AddField(position.FieldStatus, field.TypeUint8, value)
	}
	if puo.mutation.StatusCleared() {
		_spec.ClearField(position.FieldStatus, field.TypeUint8)
	}
	if value, ok := puo.mutation.Sort(); ok {
		_spec.SetField(position.FieldSort, field.TypeUint32, value)
	}
	if value, ok := puo.mutation.AddedSort(); ok {
		_spec.AddField(position.FieldSort, field.TypeUint32, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
	}
	if puo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !puo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Position{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
