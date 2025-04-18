// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"service/internal/ent/times"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TimesCreate is the builder for creating a Times entity.
type TimesCreate struct {
	config
	mutation *TimesMutation
	hooks    []Hook
}

// SetUserId sets the "userId" field.
func (tc *TimesCreate) SetUserId(i int64) *TimesCreate {
	tc.mutation.SetUserId(i)
	return tc
}

// SetTimeSpace sets the "timeSpace" field.
func (tc *TimesCreate) SetTimeSpace(s string) *TimesCreate {
	tc.mutation.SetTimeSpace(s)
	return tc
}

// SetStartTime sets the "startTime" field.
func (tc *TimesCreate) SetStartTime(s string) *TimesCreate {
	tc.mutation.SetStartTime(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TimesCreate) SetID(i int64) *TimesCreate {
	tc.mutation.SetID(i)
	return tc
}

// Mutation returns the TimesMutation object of the builder.
func (tc *TimesCreate) Mutation() *TimesMutation {
	return tc.mutation
}

// Save creates the Times in the database.
func (tc *TimesCreate) Save(ctx context.Context) (*Times, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TimesCreate) SaveX(ctx context.Context) *Times {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TimesCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TimesCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TimesCreate) check() error {
	if _, ok := tc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "Times.userId"`)}
	}
	if _, ok := tc.mutation.TimeSpace(); !ok {
		return &ValidationError{Name: "timeSpace", err: errors.New(`ent: missing required field "Times.timeSpace"`)}
	}
	if _, ok := tc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "startTime", err: errors.New(`ent: missing required field "Times.startTime"`)}
	}
	return nil
}

func (tc *TimesCreate) sqlSave(ctx context.Context) (*Times, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TimesCreate) createSpec() (*Times, *sqlgraph.CreateSpec) {
	var (
		_node = &Times{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(times.Table, sqlgraph.NewFieldSpec(times.FieldID, field.TypeInt64))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.UserId(); ok {
		_spec.SetField(times.FieldUserId, field.TypeInt64, value)
		_node.UserId = value
	}
	if value, ok := tc.mutation.TimeSpace(); ok {
		_spec.SetField(times.FieldTimeSpace, field.TypeString, value)
		_node.TimeSpace = value
	}
	if value, ok := tc.mutation.StartTime(); ok {
		_spec.SetField(times.FieldStartTime, field.TypeString, value)
		_node.StartTime = value
	}
	return _node, _spec
}

// TimesCreateBulk is the builder for creating many Times entities in bulk.
type TimesCreateBulk struct {
	config
	err      error
	builders []*TimesCreate
}

// Save creates the Times entities in the database.
func (tcb *TimesCreateBulk) Save(ctx context.Context) ([]*Times, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Times, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TimesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TimesCreateBulk) SaveX(ctx context.Context) []*Times {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TimesCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TimesCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
