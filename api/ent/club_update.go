// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Wanted-Linx/linx-backend/api/ent/club"
	"github.com/Wanted-Linx/linx-backend/api/ent/clubmember"
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
	"github.com/Wanted-Linx/linx-backend/api/ent/student"
)

// ClubUpdate is the builder for updating Club entities.
type ClubUpdate struct {
	config
	hooks    []Hook
	mutation *ClubMutation
}

// Where appends a list predicates to the ClubUpdate builder.
func (cu *ClubUpdate) Where(ps ...predicate.Club) *ClubUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ClubUpdate) SetName(s string) *ClubUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetOrganization sets the "organization" field.
func (cu *ClubUpdate) SetOrganization(s string) *ClubUpdate {
	cu.mutation.SetOrganization(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *ClubUpdate) SetDescription(s string) *ClubUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetProfileImage sets the "profile_image" field.
func (cu *ClubUpdate) SetProfileImage(s string) *ClubUpdate {
	cu.mutation.SetProfileImage(s)
	return cu
}

// SetNillableProfileImage sets the "profile_image" field if the given value is not nil.
func (cu *ClubUpdate) SetNillableProfileImage(s *string) *ClubUpdate {
	if s != nil {
		cu.SetProfileImage(*s)
	}
	return cu
}

// ClearProfileImage clears the value of the "profile_image" field.
func (cu *ClubUpdate) ClearProfileImage() *ClubUpdate {
	cu.mutation.ClearProfileImage()
	return cu
}

// SetProfileLink sets the "profile_link" field.
func (cu *ClubUpdate) SetProfileLink(s string) *ClubUpdate {
	cu.mutation.SetProfileLink(s)
	return cu
}

// SetNillableProfileLink sets the "profile_link" field if the given value is not nil.
func (cu *ClubUpdate) SetNillableProfileLink(s *string) *ClubUpdate {
	if s != nil {
		cu.SetProfileLink(*s)
	}
	return cu
}

// ClearProfileLink clears the value of the "profile_link" field.
func (cu *ClubUpdate) ClearProfileLink() *ClubUpdate {
	cu.mutation.ClearProfileLink()
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *ClubUpdate) SetCreatedAt(t time.Time) *ClubUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *ClubUpdate) SetNillableCreatedAt(t *time.Time) *ClubUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetLeaderID sets the "leader" edge to the Student entity by ID.
func (cu *ClubUpdate) SetLeaderID(id int) *ClubUpdate {
	cu.mutation.SetLeaderID(id)
	return cu
}

// SetLeader sets the "leader" edge to the Student entity.
func (cu *ClubUpdate) SetLeader(s *Student) *ClubUpdate {
	return cu.SetLeaderID(s.ID)
}

// AddClubMemberIDs adds the "club_member" edge to the ClubMember entity by IDs.
func (cu *ClubUpdate) AddClubMemberIDs(ids ...int) *ClubUpdate {
	cu.mutation.AddClubMemberIDs(ids...)
	return cu
}

// AddClubMember adds the "club_member" edges to the ClubMember entity.
func (cu *ClubUpdate) AddClubMember(c ...*ClubMember) *ClubUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddClubMemberIDs(ids...)
}

// Mutation returns the ClubMutation object of the builder.
func (cu *ClubUpdate) Mutation() *ClubMutation {
	return cu.mutation
}

// ClearLeader clears the "leader" edge to the Student entity.
func (cu *ClubUpdate) ClearLeader() *ClubUpdate {
	cu.mutation.ClearLeader()
	return cu
}

// ClearClubMember clears all "club_member" edges to the ClubMember entity.
func (cu *ClubUpdate) ClearClubMember() *ClubUpdate {
	cu.mutation.ClearClubMember()
	return cu
}

// RemoveClubMemberIDs removes the "club_member" edge to ClubMember entities by IDs.
func (cu *ClubUpdate) RemoveClubMemberIDs(ids ...int) *ClubUpdate {
	cu.mutation.RemoveClubMemberIDs(ids...)
	return cu
}

// RemoveClubMember removes "club_member" edges to ClubMember entities.
func (cu *ClubUpdate) RemoveClubMember(c ...*ClubMember) *ClubUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveClubMemberIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClubUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClubUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClubUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClubUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClubUpdate) check() error {
	if _, ok := cu.mutation.LeaderID(); cu.mutation.LeaderCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"leader\"")
	}
	return nil
}

func (cu *ClubUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   club.Table,
			Columns: club.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: club.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldName,
		})
	}
	if value, ok := cu.mutation.Organization(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldOrganization,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldDescription,
		})
	}
	if value, ok := cu.mutation.ProfileImage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldProfileImage,
		})
	}
	if cu.mutation.ProfileImageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: club.FieldProfileImage,
		})
	}
	if value, ok := cu.mutation.ProfileLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldProfileLink,
		})
	}
	if cu.mutation.ProfileLinkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: club.FieldProfileLink,
		})
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: club.FieldCreatedAt,
		})
	}
	if cu.mutation.LeaderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   club.LeaderTable,
			Columns: []string{club.LeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   club.LeaderTable,
			Columns: []string{club.LeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ClubMemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   club.ClubMemberTable,
			Columns: []string{club.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedClubMemberIDs(); len(nodes) > 0 && !cu.mutation.ClubMemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   club.ClubMemberTable,
			Columns: []string{club.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClubMemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   club.ClubMemberTable,
			Columns: []string{club.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{club.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ClubUpdateOne is the builder for updating a single Club entity.
type ClubUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClubMutation
}

// SetName sets the "name" field.
func (cuo *ClubUpdateOne) SetName(s string) *ClubUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetOrganization sets the "organization" field.
func (cuo *ClubUpdateOne) SetOrganization(s string) *ClubUpdateOne {
	cuo.mutation.SetOrganization(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ClubUpdateOne) SetDescription(s string) *ClubUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetProfileImage sets the "profile_image" field.
func (cuo *ClubUpdateOne) SetProfileImage(s string) *ClubUpdateOne {
	cuo.mutation.SetProfileImage(s)
	return cuo
}

// SetNillableProfileImage sets the "profile_image" field if the given value is not nil.
func (cuo *ClubUpdateOne) SetNillableProfileImage(s *string) *ClubUpdateOne {
	if s != nil {
		cuo.SetProfileImage(*s)
	}
	return cuo
}

// ClearProfileImage clears the value of the "profile_image" field.
func (cuo *ClubUpdateOne) ClearProfileImage() *ClubUpdateOne {
	cuo.mutation.ClearProfileImage()
	return cuo
}

// SetProfileLink sets the "profile_link" field.
func (cuo *ClubUpdateOne) SetProfileLink(s string) *ClubUpdateOne {
	cuo.mutation.SetProfileLink(s)
	return cuo
}

// SetNillableProfileLink sets the "profile_link" field if the given value is not nil.
func (cuo *ClubUpdateOne) SetNillableProfileLink(s *string) *ClubUpdateOne {
	if s != nil {
		cuo.SetProfileLink(*s)
	}
	return cuo
}

// ClearProfileLink clears the value of the "profile_link" field.
func (cuo *ClubUpdateOne) ClearProfileLink() *ClubUpdateOne {
	cuo.mutation.ClearProfileLink()
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *ClubUpdateOne) SetCreatedAt(t time.Time) *ClubUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *ClubUpdateOne) SetNillableCreatedAt(t *time.Time) *ClubUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetLeaderID sets the "leader" edge to the Student entity by ID.
func (cuo *ClubUpdateOne) SetLeaderID(id int) *ClubUpdateOne {
	cuo.mutation.SetLeaderID(id)
	return cuo
}

// SetLeader sets the "leader" edge to the Student entity.
func (cuo *ClubUpdateOne) SetLeader(s *Student) *ClubUpdateOne {
	return cuo.SetLeaderID(s.ID)
}

// AddClubMemberIDs adds the "club_member" edge to the ClubMember entity by IDs.
func (cuo *ClubUpdateOne) AddClubMemberIDs(ids ...int) *ClubUpdateOne {
	cuo.mutation.AddClubMemberIDs(ids...)
	return cuo
}

// AddClubMember adds the "club_member" edges to the ClubMember entity.
func (cuo *ClubUpdateOne) AddClubMember(c ...*ClubMember) *ClubUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddClubMemberIDs(ids...)
}

// Mutation returns the ClubMutation object of the builder.
func (cuo *ClubUpdateOne) Mutation() *ClubMutation {
	return cuo.mutation
}

// ClearLeader clears the "leader" edge to the Student entity.
func (cuo *ClubUpdateOne) ClearLeader() *ClubUpdateOne {
	cuo.mutation.ClearLeader()
	return cuo
}

// ClearClubMember clears all "club_member" edges to the ClubMember entity.
func (cuo *ClubUpdateOne) ClearClubMember() *ClubUpdateOne {
	cuo.mutation.ClearClubMember()
	return cuo
}

// RemoveClubMemberIDs removes the "club_member" edge to ClubMember entities by IDs.
func (cuo *ClubUpdateOne) RemoveClubMemberIDs(ids ...int) *ClubUpdateOne {
	cuo.mutation.RemoveClubMemberIDs(ids...)
	return cuo
}

// RemoveClubMember removes "club_member" edges to ClubMember entities.
func (cuo *ClubUpdateOne) RemoveClubMember(c ...*ClubMember) *ClubUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveClubMemberIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClubUpdateOne) Select(field string, fields ...string) *ClubUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Club entity.
func (cuo *ClubUpdateOne) Save(ctx context.Context) (*Club, error) {
	var (
		err  error
		node *Club
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClubUpdateOne) SaveX(ctx context.Context) *Club {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClubUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClubUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClubUpdateOne) check() error {
	if _, ok := cuo.mutation.LeaderID(); cuo.mutation.LeaderCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"leader\"")
	}
	return nil
}

func (cuo *ClubUpdateOne) sqlSave(ctx context.Context) (_node *Club, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   club.Table,
			Columns: club.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: club.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Club.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, club.FieldID)
		for _, f := range fields {
			if !club.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != club.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldName,
		})
	}
	if value, ok := cuo.mutation.Organization(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldOrganization,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.ProfileImage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldProfileImage,
		})
	}
	if cuo.mutation.ProfileImageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: club.FieldProfileImage,
		})
	}
	if value, ok := cuo.mutation.ProfileLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldProfileLink,
		})
	}
	if cuo.mutation.ProfileLinkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: club.FieldProfileLink,
		})
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: club.FieldCreatedAt,
		})
	}
	if cuo.mutation.LeaderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   club.LeaderTable,
			Columns: []string{club.LeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   club.LeaderTable,
			Columns: []string{club.LeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ClubMemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   club.ClubMemberTable,
			Columns: []string{club.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedClubMemberIDs(); len(nodes) > 0 && !cuo.mutation.ClubMemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   club.ClubMemberTable,
			Columns: []string{club.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClubMemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   club.ClubMemberTable,
			Columns: []string{club.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Club{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{club.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}