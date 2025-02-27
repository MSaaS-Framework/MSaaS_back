// Code generated by ent, DO NOT EDIT.

package ent

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/database"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/generalspec"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/predicate"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/service"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DatabaseUpdate is the builder for updating Database entities.
type DatabaseUpdate struct {
	config
	hooks    []Hook
	mutation *DatabaseMutation
}

// Where appends a list predicates to the DatabaseUpdate builder.
func (du *DatabaseUpdate) Where(ps ...predicate.Database) *DatabaseUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetConnectionPath sets the "connection_path" field.
func (du *DatabaseUpdate) SetConnectionPath(s string) *DatabaseUpdate {
	du.mutation.SetConnectionPath(s)
	return du
}

// SetNillableConnectionPath sets the "connection_path" field if the given value is not nil.
func (du *DatabaseUpdate) SetNillableConnectionPath(s *string) *DatabaseUpdate {
	if s != nil {
		du.SetConnectionPath(*s)
	}
	return du
}

// SetPassword sets the "password" field.
func (du *DatabaseUpdate) SetPassword(s string) *DatabaseUpdate {
	du.mutation.SetPassword(s)
	return du
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (du *DatabaseUpdate) SetNillablePassword(s *string) *DatabaseUpdate {
	if s != nil {
		du.SetPassword(*s)
	}
	return du
}

// SetDbType sets the "db_type" field.
func (du *DatabaseUpdate) SetDbType(s string) *DatabaseUpdate {
	du.mutation.SetDbType(s)
	return du
}

// SetNillableDbType sets the "db_type" field if the given value is not nil.
func (du *DatabaseUpdate) SetNillableDbType(s *string) *DatabaseUpdate {
	if s != nil {
		du.SetDbType(*s)
	}
	return du
}

// SetServiceID sets the "service" edge to the Service entity by ID.
func (du *DatabaseUpdate) SetServiceID(id uuid.UUID) *DatabaseUpdate {
	du.mutation.SetServiceID(id)
	return du
}

// SetNillableServiceID sets the "service" edge to the Service entity by ID if the given value is not nil.
func (du *DatabaseUpdate) SetNillableServiceID(id *uuid.UUID) *DatabaseUpdate {
	if id != nil {
		du = du.SetServiceID(*id)
	}
	return du
}

// SetService sets the "service" edge to the Service entity.
func (du *DatabaseUpdate) SetService(s *Service) *DatabaseUpdate {
	return du.SetServiceID(s.ID)
}

// SetGeneralspecID sets the "generalspec" edge to the GeneralSpec entity by ID.
func (du *DatabaseUpdate) SetGeneralspecID(id int) *DatabaseUpdate {
	du.mutation.SetGeneralspecID(id)
	return du
}

// SetGeneralspec sets the "generalspec" edge to the GeneralSpec entity.
func (du *DatabaseUpdate) SetGeneralspec(g *GeneralSpec) *DatabaseUpdate {
	return du.SetGeneralspecID(g.ID)
}

// Mutation returns the DatabaseMutation object of the builder.
func (du *DatabaseUpdate) Mutation() *DatabaseMutation {
	return du.mutation
}

// ClearService clears the "service" edge to the Service entity.
func (du *DatabaseUpdate) ClearService() *DatabaseUpdate {
	du.mutation.ClearService()
	return du
}

// ClearGeneralspec clears the "generalspec" edge to the GeneralSpec entity.
func (du *DatabaseUpdate) ClearGeneralspec() *DatabaseUpdate {
	du.mutation.ClearGeneralspec()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DatabaseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DatabaseUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DatabaseUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DatabaseUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DatabaseUpdate) check() error {
	if v, ok := du.mutation.ConnectionPath(); ok {
		if err := database.ConnectionPathValidator(v); err != nil {
			return &ValidationError{Name: "connection_path", err: fmt.Errorf(`ent: validator failed for field "Database.connection_path": %w`, err)}
		}
	}
	if v, ok := du.mutation.DbType(); ok {
		if err := database.DbTypeValidator(v); err != nil {
			return &ValidationError{Name: "db_type", err: fmt.Errorf(`ent: validator failed for field "Database.db_type": %w`, err)}
		}
	}
	if du.mutation.GeneralspecCleared() && len(du.mutation.GeneralspecIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Database.generalspec"`)
	}
	return nil
}

func (du *DatabaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(database.Table, database.Columns, sqlgraph.NewFieldSpec(database.FieldID, field.TypeUUID))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.ConnectionPath(); ok {
		_spec.SetField(database.FieldConnectionPath, field.TypeString, value)
	}
	if value, ok := du.mutation.Password(); ok {
		_spec.SetField(database.FieldPassword, field.TypeString, value)
	}
	if value, ok := du.mutation.DbType(); ok {
		_spec.SetField(database.FieldDbType, field.TypeString, value)
	}
	if du.mutation.ServiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   database.ServiceTable,
			Columns: []string{database.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   database.ServiceTable,
			Columns: []string{database.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.GeneralspecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   database.GeneralspecTable,
			Columns: []string{database.GeneralspecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalspec.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.GeneralspecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   database.GeneralspecTable,
			Columns: []string{database.GeneralspecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalspec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{database.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DatabaseUpdateOne is the builder for updating a single Database entity.
type DatabaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DatabaseMutation
}

// SetConnectionPath sets the "connection_path" field.
func (duo *DatabaseUpdateOne) SetConnectionPath(s string) *DatabaseUpdateOne {
	duo.mutation.SetConnectionPath(s)
	return duo
}

// SetNillableConnectionPath sets the "connection_path" field if the given value is not nil.
func (duo *DatabaseUpdateOne) SetNillableConnectionPath(s *string) *DatabaseUpdateOne {
	if s != nil {
		duo.SetConnectionPath(*s)
	}
	return duo
}

// SetPassword sets the "password" field.
func (duo *DatabaseUpdateOne) SetPassword(s string) *DatabaseUpdateOne {
	duo.mutation.SetPassword(s)
	return duo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (duo *DatabaseUpdateOne) SetNillablePassword(s *string) *DatabaseUpdateOne {
	if s != nil {
		duo.SetPassword(*s)
	}
	return duo
}

// SetDbType sets the "db_type" field.
func (duo *DatabaseUpdateOne) SetDbType(s string) *DatabaseUpdateOne {
	duo.mutation.SetDbType(s)
	return duo
}

// SetNillableDbType sets the "db_type" field if the given value is not nil.
func (duo *DatabaseUpdateOne) SetNillableDbType(s *string) *DatabaseUpdateOne {
	if s != nil {
		duo.SetDbType(*s)
	}
	return duo
}

// SetServiceID sets the "service" edge to the Service entity by ID.
func (duo *DatabaseUpdateOne) SetServiceID(id uuid.UUID) *DatabaseUpdateOne {
	duo.mutation.SetServiceID(id)
	return duo
}

// SetNillableServiceID sets the "service" edge to the Service entity by ID if the given value is not nil.
func (duo *DatabaseUpdateOne) SetNillableServiceID(id *uuid.UUID) *DatabaseUpdateOne {
	if id != nil {
		duo = duo.SetServiceID(*id)
	}
	return duo
}

// SetService sets the "service" edge to the Service entity.
func (duo *DatabaseUpdateOne) SetService(s *Service) *DatabaseUpdateOne {
	return duo.SetServiceID(s.ID)
}

// SetGeneralspecID sets the "generalspec" edge to the GeneralSpec entity by ID.
func (duo *DatabaseUpdateOne) SetGeneralspecID(id int) *DatabaseUpdateOne {
	duo.mutation.SetGeneralspecID(id)
	return duo
}

// SetGeneralspec sets the "generalspec" edge to the GeneralSpec entity.
func (duo *DatabaseUpdateOne) SetGeneralspec(g *GeneralSpec) *DatabaseUpdateOne {
	return duo.SetGeneralspecID(g.ID)
}

// Mutation returns the DatabaseMutation object of the builder.
func (duo *DatabaseUpdateOne) Mutation() *DatabaseMutation {
	return duo.mutation
}

// ClearService clears the "service" edge to the Service entity.
func (duo *DatabaseUpdateOne) ClearService() *DatabaseUpdateOne {
	duo.mutation.ClearService()
	return duo
}

// ClearGeneralspec clears the "generalspec" edge to the GeneralSpec entity.
func (duo *DatabaseUpdateOne) ClearGeneralspec() *DatabaseUpdateOne {
	duo.mutation.ClearGeneralspec()
	return duo
}

// Where appends a list predicates to the DatabaseUpdate builder.
func (duo *DatabaseUpdateOne) Where(ps ...predicate.Database) *DatabaseUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DatabaseUpdateOne) Select(field string, fields ...string) *DatabaseUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Database entity.
func (duo *DatabaseUpdateOne) Save(ctx context.Context) (*Database, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DatabaseUpdateOne) SaveX(ctx context.Context) *Database {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DatabaseUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DatabaseUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DatabaseUpdateOne) check() error {
	if v, ok := duo.mutation.ConnectionPath(); ok {
		if err := database.ConnectionPathValidator(v); err != nil {
			return &ValidationError{Name: "connection_path", err: fmt.Errorf(`ent: validator failed for field "Database.connection_path": %w`, err)}
		}
	}
	if v, ok := duo.mutation.DbType(); ok {
		if err := database.DbTypeValidator(v); err != nil {
			return &ValidationError{Name: "db_type", err: fmt.Errorf(`ent: validator failed for field "Database.db_type": %w`, err)}
		}
	}
	if duo.mutation.GeneralspecCleared() && len(duo.mutation.GeneralspecIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Database.generalspec"`)
	}
	return nil
}

func (duo *DatabaseUpdateOne) sqlSave(ctx context.Context) (_node *Database, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(database.Table, database.Columns, sqlgraph.NewFieldSpec(database.FieldID, field.TypeUUID))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Database.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, database.FieldID)
		for _, f := range fields {
			if !database.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != database.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.ConnectionPath(); ok {
		_spec.SetField(database.FieldConnectionPath, field.TypeString, value)
	}
	if value, ok := duo.mutation.Password(); ok {
		_spec.SetField(database.FieldPassword, field.TypeString, value)
	}
	if value, ok := duo.mutation.DbType(); ok {
		_spec.SetField(database.FieldDbType, field.TypeString, value)
	}
	if duo.mutation.ServiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   database.ServiceTable,
			Columns: []string{database.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   database.ServiceTable,
			Columns: []string{database.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.GeneralspecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   database.GeneralspecTable,
			Columns: []string{database.GeneralspecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalspec.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.GeneralspecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   database.GeneralspecTable,
			Columns: []string{database.GeneralspecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalspec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Database{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{database.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
