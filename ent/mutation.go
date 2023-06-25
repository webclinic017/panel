// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"panel/ent/predicate"
	"panel/ent/role"
	"panel/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRole = "Role"
	TypeUser = "User"
)

// RoleMutation represents an operation that mutates the Role nodes in the graph.
type RoleMutation struct {
	config
	op                Op
	typ               string
	id                *int
	created_at        *time.Time
	updated_at        *time.Time
	name              *string
	permissions       *[]string
	appendpermissions []string
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*Role, error)
	predicates        []predicate.Role
}

var _ ent.Mutation = (*RoleMutation)(nil)

// roleOption allows management of the mutation configuration using functional options.
type roleOption func(*RoleMutation)

// newRoleMutation creates new mutation for the Role entity.
func newRoleMutation(c config, op Op, opts ...roleOption) *RoleMutation {
	m := &RoleMutation{
		config:        c,
		op:            op,
		typ:           TypeRole,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRoleID sets the ID field of the mutation.
func withRoleID(id int) roleOption {
	return func(m *RoleMutation) {
		var (
			err   error
			once  sync.Once
			value *Role
		)
		m.oldValue = func(ctx context.Context) (*Role, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Role.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRole sets the old Role of the mutation.
func withRole(node *Role) roleOption {
	return func(m *RoleMutation) {
		m.oldValue = func(context.Context) (*Role, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RoleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RoleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RoleMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RoleMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Role.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *RoleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *RoleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *RoleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *RoleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *RoleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *RoleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *RoleMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *RoleMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *RoleMutation) ResetName() {
	m.name = nil
}

// SetPermissions sets the "permissions" field.
func (m *RoleMutation) SetPermissions(s []string) {
	m.permissions = &s
	m.appendpermissions = nil
}

// Permissions returns the value of the "permissions" field in the mutation.
func (m *RoleMutation) Permissions() (r []string, exists bool) {
	v := m.permissions
	if v == nil {
		return
	}
	return *v, true
}

// OldPermissions returns the old "permissions" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldPermissions(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPermissions is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPermissions requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPermissions: %w", err)
	}
	return oldValue.Permissions, nil
}

// AppendPermissions adds s to the "permissions" field.
func (m *RoleMutation) AppendPermissions(s []string) {
	m.appendpermissions = append(m.appendpermissions, s...)
}

// AppendedPermissions returns the list of values that were appended to the "permissions" field in this mutation.
func (m *RoleMutation) AppendedPermissions() ([]string, bool) {
	if len(m.appendpermissions) == 0 {
		return nil, false
	}
	return m.appendpermissions, true
}

// ClearPermissions clears the value of the "permissions" field.
func (m *RoleMutation) ClearPermissions() {
	m.permissions = nil
	m.appendpermissions = nil
	m.clearedFields[role.FieldPermissions] = struct{}{}
}

// PermissionsCleared returns if the "permissions" field was cleared in this mutation.
func (m *RoleMutation) PermissionsCleared() bool {
	_, ok := m.clearedFields[role.FieldPermissions]
	return ok
}

// ResetPermissions resets all changes to the "permissions" field.
func (m *RoleMutation) ResetPermissions() {
	m.permissions = nil
	m.appendpermissions = nil
	delete(m.clearedFields, role.FieldPermissions)
}

// Where appends a list predicates to the RoleMutation builder.
func (m *RoleMutation) Where(ps ...predicate.Role) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the RoleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *RoleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Role, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *RoleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *RoleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Role).
func (m *RoleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RoleMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, role.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, role.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, role.FieldName)
	}
	if m.permissions != nil {
		fields = append(fields, role.FieldPermissions)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RoleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case role.FieldCreatedAt:
		return m.CreatedAt()
	case role.FieldUpdatedAt:
		return m.UpdatedAt()
	case role.FieldName:
		return m.Name()
	case role.FieldPermissions:
		return m.Permissions()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RoleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case role.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case role.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case role.FieldName:
		return m.OldName(ctx)
	case role.FieldPermissions:
		return m.OldPermissions(ctx)
	}
	return nil, fmt.Errorf("unknown Role field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case role.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case role.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case role.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case role.FieldPermissions:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPermissions(v)
		return nil
	}
	return fmt.Errorf("unknown Role field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RoleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RoleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Role numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RoleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(role.FieldPermissions) {
		fields = append(fields, role.FieldPermissions)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RoleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RoleMutation) ClearField(name string) error {
	switch name {
	case role.FieldPermissions:
		m.ClearPermissions()
		return nil
	}
	return fmt.Errorf("unknown Role nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RoleMutation) ResetField(name string) error {
	switch name {
	case role.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case role.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case role.FieldName:
		m.ResetName()
		return nil
	case role.FieldPermissions:
		m.ResetPermissions()
		return nil
	}
	return fmt.Errorf("unknown Role field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RoleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RoleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RoleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RoleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RoleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RoleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RoleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Role unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RoleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Role edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	uuid          *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	email         *string
	password      *string
	pfp           *[]byte
	name          *string
	clearedFields map[string]struct{}
	role          *int
	clearedrole   bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUUID sets the "uuid" field.
func (m *UserMutation) SetUUID(u uuid.UUID) {
	m.uuid = &u
}

// UUID returns the value of the "uuid" field in the mutation.
func (m *UserMutation) UUID() (r uuid.UUID, exists bool) {
	v := m.uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUUID returns the old "uuid" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUUID: %w", err)
	}
	return oldValue.UUID, nil
}

// ResetUUID resets all changes to the "uuid" field.
func (m *UserMutation) ResetUUID() {
	m.uuid = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetPfp sets the "pfp" field.
func (m *UserMutation) SetPfp(b []byte) {
	m.pfp = &b
}

// Pfp returns the value of the "pfp" field in the mutation.
func (m *UserMutation) Pfp() (r []byte, exists bool) {
	v := m.pfp
	if v == nil {
		return
	}
	return *v, true
}

// OldPfp returns the old "pfp" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPfp(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPfp is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPfp requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPfp: %w", err)
	}
	return oldValue.Pfp, nil
}

// ResetPfp resets all changes to the "pfp" field.
func (m *UserMutation) ResetPfp() {
	m.pfp = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetRoleID sets the "role_id" field.
func (m *UserMutation) SetRoleID(i int) {
	m.role = &i
}

// RoleID returns the value of the "role_id" field in the mutation.
func (m *UserMutation) RoleID() (r int, exists bool) {
	v := m.role
	if v == nil {
		return
	}
	return *v, true
}

// OldRoleID returns the old "role_id" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRoleID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRoleID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRoleID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRoleID: %w", err)
	}
	return oldValue.RoleID, nil
}

// ResetRoleID resets all changes to the "role_id" field.
func (m *UserMutation) ResetRoleID() {
	m.role = nil
}

// ClearRole clears the "role" edge to the Role entity.
func (m *UserMutation) ClearRole() {
	m.clearedrole = true
}

// RoleCleared reports if the "role" edge to the Role entity was cleared.
func (m *UserMutation) RoleCleared() bool {
	return m.clearedrole
}

// RoleIDs returns the "role" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// RoleID instead. It exists only for internal usage by the builders.
func (m *UserMutation) RoleIDs() (ids []int) {
	if id := m.role; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetRole resets all changes to the "role" edge.
func (m *UserMutation) ResetRole() {
	m.role = nil
	m.clearedrole = false
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.uuid != nil {
		fields = append(fields, user.FieldUUID)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.pfp != nil {
		fields = append(fields, user.FieldPfp)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.role != nil {
		fields = append(fields, user.FieldRoleID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUUID:
		return m.UUID()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	case user.FieldPfp:
		return m.Pfp()
	case user.FieldName:
		return m.Name()
	case user.FieldRoleID:
		return m.RoleID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUUID:
		return m.OldUUID(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldPfp:
		return m.OldPfp(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldRoleID:
		return m.OldRoleID(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUUID(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldPfp:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPfp(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldRoleID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRoleID(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUUID:
		m.ResetUUID()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldPfp:
		m.ResetPfp()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldRoleID:
		m.ResetRoleID()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.role != nil {
		edges = append(edges, user.EdgeRole)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeRole:
		if id := m.role; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedrole {
		edges = append(edges, user.EdgeRole)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeRole:
		return m.clearedrole
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeRole:
		m.ClearRole()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeRole:
		m.ResetRole()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
