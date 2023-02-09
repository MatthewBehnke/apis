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

	"github.com/MatthewBehnke/apis/pkg/database/ent/predicate"
	"github.com/MatthewBehnke/apis/pkg/database/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPasswordHash sets the "password_hash" field.
func (uu *UserUpdate) SetPasswordHash(s string) *UserUpdate {
	uu.mutation.SetPasswordHash(s)
	return uu
}

// SetNillablePasswordHash sets the "password_hash" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePasswordHash(s *string) *UserUpdate {
	if s != nil {
		uu.SetPasswordHash(*s)
	}
	return uu
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (uu *UserUpdate) ClearPasswordHash() *UserUpdate {
	uu.mutation.ClearPasswordHash()
	return uu
}

// SetAttemptCount sets the "attempt_count" field.
func (uu *UserUpdate) SetAttemptCount(i int) *UserUpdate {
	uu.mutation.ResetAttemptCount()
	uu.mutation.SetAttemptCount(i)
	return uu
}

// SetNillableAttemptCount sets the "attempt_count" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAttemptCount(i *int) *UserUpdate {
	if i != nil {
		uu.SetAttemptCount(*i)
	}
	return uu
}

// AddAttemptCount adds i to the "attempt_count" field.
func (uu *UserUpdate) AddAttemptCount(i int) *UserUpdate {
	uu.mutation.AddAttemptCount(i)
	return uu
}

// ClearAttemptCount clears the value of the "attempt_count" field.
func (uu *UserUpdate) ClearAttemptCount() *UserUpdate {
	uu.mutation.ClearAttemptCount()
	return uu
}

// SetLastAttempt sets the "last_attempt" field.
func (uu *UserUpdate) SetLastAttempt(t time.Time) *UserUpdate {
	uu.mutation.SetLastAttempt(t)
	return uu
}

// SetNillableLastAttempt sets the "last_attempt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastAttempt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastAttempt(*t)
	}
	return uu
}

// ClearLastAttempt clears the value of the "last_attempt" field.
func (uu *UserUpdate) ClearLastAttempt() *UserUpdate {
	uu.mutation.ClearLastAttempt()
	return uu
}

// SetLocked sets the "locked" field.
func (uu *UserUpdate) SetLocked(t time.Time) *UserUpdate {
	uu.mutation.SetLocked(t)
	return uu
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLocked(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLocked(*t)
	}
	return uu
}

// ClearLocked clears the value of the "locked" field.
func (uu *UserUpdate) ClearLocked() *UserUpdate {
	uu.mutation.ClearLocked()
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(s string) *UserUpdate {
	uu.mutation.SetRole(s)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(s *string) *UserUpdate {
	if s != nil {
		uu.SetRole(*s)
	}
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if uu.mutation.PasswordHashCleared() {
		_spec.ClearField(user.FieldPasswordHash, field.TypeString)
	}
	if value, ok := uu.mutation.AttemptCount(); ok {
		_spec.SetField(user.FieldAttemptCount, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedAttemptCount(); ok {
		_spec.AddField(user.FieldAttemptCount, field.TypeInt, value)
	}
	if uu.mutation.AttemptCountCleared() {
		_spec.ClearField(user.FieldAttemptCount, field.TypeInt)
	}
	if value, ok := uu.mutation.LastAttempt(); ok {
		_spec.SetField(user.FieldLastAttempt, field.TypeTime, value)
	}
	if uu.mutation.LastAttemptCleared() {
		_spec.ClearField(user.FieldLastAttempt, field.TypeTime)
	}
	if value, ok := uu.mutation.Locked(); ok {
		_spec.SetField(user.FieldLocked, field.TypeTime, value)
	}
	if uu.mutation.LockedCleared() {
		_spec.ClearField(user.FieldLocked, field.TypeTime)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPasswordHash sets the "password_hash" field.
func (uuo *UserUpdateOne) SetPasswordHash(s string) *UserUpdateOne {
	uuo.mutation.SetPasswordHash(s)
	return uuo
}

// SetNillablePasswordHash sets the "password_hash" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePasswordHash(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPasswordHash(*s)
	}
	return uuo
}

// ClearPasswordHash clears the value of the "password_hash" field.
func (uuo *UserUpdateOne) ClearPasswordHash() *UserUpdateOne {
	uuo.mutation.ClearPasswordHash()
	return uuo
}

// SetAttemptCount sets the "attempt_count" field.
func (uuo *UserUpdateOne) SetAttemptCount(i int) *UserUpdateOne {
	uuo.mutation.ResetAttemptCount()
	uuo.mutation.SetAttemptCount(i)
	return uuo
}

// SetNillableAttemptCount sets the "attempt_count" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAttemptCount(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetAttemptCount(*i)
	}
	return uuo
}

// AddAttemptCount adds i to the "attempt_count" field.
func (uuo *UserUpdateOne) AddAttemptCount(i int) *UserUpdateOne {
	uuo.mutation.AddAttemptCount(i)
	return uuo
}

// ClearAttemptCount clears the value of the "attempt_count" field.
func (uuo *UserUpdateOne) ClearAttemptCount() *UserUpdateOne {
	uuo.mutation.ClearAttemptCount()
	return uuo
}

// SetLastAttempt sets the "last_attempt" field.
func (uuo *UserUpdateOne) SetLastAttempt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastAttempt(t)
	return uuo
}

// SetNillableLastAttempt sets the "last_attempt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastAttempt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastAttempt(*t)
	}
	return uuo
}

// ClearLastAttempt clears the value of the "last_attempt" field.
func (uuo *UserUpdateOne) ClearLastAttempt() *UserUpdateOne {
	uuo.mutation.ClearLastAttempt()
	return uuo
}

// SetLocked sets the "locked" field.
func (uuo *UserUpdateOne) SetLocked(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLocked(t)
	return uuo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLocked(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLocked(*t)
	}
	return uuo
}

// ClearLocked clears the value of the "locked" field.
func (uuo *UserUpdateOne) ClearLocked() *UserUpdateOne {
	uuo.mutation.ClearLocked()
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(s string) *UserUpdateOne {
	uuo.mutation.SetRole(s)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRole(*s)
	}
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if uuo.mutation.PasswordHashCleared() {
		_spec.ClearField(user.FieldPasswordHash, field.TypeString)
	}
	if value, ok := uuo.mutation.AttemptCount(); ok {
		_spec.SetField(user.FieldAttemptCount, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedAttemptCount(); ok {
		_spec.AddField(user.FieldAttemptCount, field.TypeInt, value)
	}
	if uuo.mutation.AttemptCountCleared() {
		_spec.ClearField(user.FieldAttemptCount, field.TypeInt)
	}
	if value, ok := uuo.mutation.LastAttempt(); ok {
		_spec.SetField(user.FieldLastAttempt, field.TypeTime, value)
	}
	if uuo.mutation.LastAttemptCleared() {
		_spec.ClearField(user.FieldLastAttempt, field.TypeTime)
	}
	if value, ok := uuo.mutation.Locked(); ok {
		_spec.SetField(user.FieldLocked, field.TypeTime, value)
	}
	if uuo.mutation.LockedCleared() {
		_spec.ClearField(user.FieldLocked, field.TypeTime)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeString, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
