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
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
	"github.com/Wanted-Linx/linx-backend/api/ent/project"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectlog"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectlogfeedback"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectlogparticipant"
)

// ProjectLogUpdate is the builder for updating ProjectLog entities.
type ProjectLogUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectLogMutation
}

// Where appends a list predicates to the ProjectLogUpdate builder.
func (plu *ProjectLogUpdate) Where(ps ...predicate.ProjectLog) *ProjectLogUpdate {
	plu.mutation.Where(ps...)
	return plu
}

// SetTitle sets the "title" field.
func (plu *ProjectLogUpdate) SetTitle(s string) *ProjectLogUpdate {
	plu.mutation.SetTitle(s)
	return plu
}

// SetAuthor sets the "author" field.
func (plu *ProjectLogUpdate) SetAuthor(s string) *ProjectLogUpdate {
	plu.mutation.SetAuthor(s)
	return plu
}

// SetContent sets the "content" field.
func (plu *ProjectLogUpdate) SetContent(s string) *ProjectLogUpdate {
	plu.mutation.SetContent(s)
	return plu
}

// SetStartDate sets the "start_date" field.
func (plu *ProjectLogUpdate) SetStartDate(s string) *ProjectLogUpdate {
	plu.mutation.SetStartDate(s)
	return plu
}

// SetEndDate sets the "end_date" field.
func (plu *ProjectLogUpdate) SetEndDate(s string) *ProjectLogUpdate {
	plu.mutation.SetEndDate(s)
	return plu
}

// SetCreatedAt sets the "created_at" field.
func (plu *ProjectLogUpdate) SetCreatedAt(t time.Time) *ProjectLogUpdate {
	plu.mutation.SetCreatedAt(t)
	return plu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (plu *ProjectLogUpdate) SetNillableCreatedAt(t *time.Time) *ProjectLogUpdate {
	if t != nil {
		plu.SetCreatedAt(*t)
	}
	return plu
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (plu *ProjectLogUpdate) SetProjectID(id int) *ProjectLogUpdate {
	plu.mutation.SetProjectID(id)
	return plu
}

// SetProject sets the "project" edge to the Project entity.
func (plu *ProjectLogUpdate) SetProject(p *Project) *ProjectLogUpdate {
	return plu.SetProjectID(p.ID)
}

// AddProjectLogParticipantIDs adds the "project_log_participant" edge to the ProjectLogParticipant entity by IDs.
func (plu *ProjectLogUpdate) AddProjectLogParticipantIDs(ids ...int) *ProjectLogUpdate {
	plu.mutation.AddProjectLogParticipantIDs(ids...)
	return plu
}

// AddProjectLogParticipant adds the "project_log_participant" edges to the ProjectLogParticipant entity.
func (plu *ProjectLogUpdate) AddProjectLogParticipant(p ...*ProjectLogParticipant) *ProjectLogUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plu.AddProjectLogParticipantIDs(ids...)
}

// AddProjectLogFeedbackIDs adds the "project_log_feedback" edge to the ProjectLogFeedback entity by IDs.
func (plu *ProjectLogUpdate) AddProjectLogFeedbackIDs(ids ...int) *ProjectLogUpdate {
	plu.mutation.AddProjectLogFeedbackIDs(ids...)
	return plu
}

// AddProjectLogFeedback adds the "project_log_feedback" edges to the ProjectLogFeedback entity.
func (plu *ProjectLogUpdate) AddProjectLogFeedback(p ...*ProjectLogFeedback) *ProjectLogUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plu.AddProjectLogFeedbackIDs(ids...)
}

// Mutation returns the ProjectLogMutation object of the builder.
func (plu *ProjectLogUpdate) Mutation() *ProjectLogMutation {
	return plu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (plu *ProjectLogUpdate) ClearProject() *ProjectLogUpdate {
	plu.mutation.ClearProject()
	return plu
}

// ClearProjectLogParticipant clears all "project_log_participant" edges to the ProjectLogParticipant entity.
func (plu *ProjectLogUpdate) ClearProjectLogParticipant() *ProjectLogUpdate {
	plu.mutation.ClearProjectLogParticipant()
	return plu
}

// RemoveProjectLogParticipantIDs removes the "project_log_participant" edge to ProjectLogParticipant entities by IDs.
func (plu *ProjectLogUpdate) RemoveProjectLogParticipantIDs(ids ...int) *ProjectLogUpdate {
	plu.mutation.RemoveProjectLogParticipantIDs(ids...)
	return plu
}

// RemoveProjectLogParticipant removes "project_log_participant" edges to ProjectLogParticipant entities.
func (plu *ProjectLogUpdate) RemoveProjectLogParticipant(p ...*ProjectLogParticipant) *ProjectLogUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plu.RemoveProjectLogParticipantIDs(ids...)
}

// ClearProjectLogFeedback clears all "project_log_feedback" edges to the ProjectLogFeedback entity.
func (plu *ProjectLogUpdate) ClearProjectLogFeedback() *ProjectLogUpdate {
	plu.mutation.ClearProjectLogFeedback()
	return plu
}

// RemoveProjectLogFeedbackIDs removes the "project_log_feedback" edge to ProjectLogFeedback entities by IDs.
func (plu *ProjectLogUpdate) RemoveProjectLogFeedbackIDs(ids ...int) *ProjectLogUpdate {
	plu.mutation.RemoveProjectLogFeedbackIDs(ids...)
	return plu
}

// RemoveProjectLogFeedback removes "project_log_feedback" edges to ProjectLogFeedback entities.
func (plu *ProjectLogUpdate) RemoveProjectLogFeedback(p ...*ProjectLogFeedback) *ProjectLogUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return plu.RemoveProjectLogFeedbackIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (plu *ProjectLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(plu.hooks) == 0 {
		if err = plu.check(); err != nil {
			return 0, err
		}
		affected, err = plu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = plu.check(); err != nil {
				return 0, err
			}
			plu.mutation = mutation
			affected, err = plu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(plu.hooks) - 1; i >= 0; i-- {
			if plu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = plu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, plu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (plu *ProjectLogUpdate) SaveX(ctx context.Context) int {
	affected, err := plu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (plu *ProjectLogUpdate) Exec(ctx context.Context) error {
	_, err := plu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plu *ProjectLogUpdate) ExecX(ctx context.Context) {
	if err := plu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plu *ProjectLogUpdate) check() error {
	if _, ok := plu.mutation.ProjectID(); plu.mutation.ProjectCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project\"")
	}
	return nil
}

func (plu *ProjectLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectlog.Table,
			Columns: projectlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectlog.FieldID,
			},
		},
	}
	if ps := plu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := plu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldTitle,
		})
	}
	if value, ok := plu.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldAuthor,
		})
	}
	if value, ok := plu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldContent,
		})
	}
	if value, ok := plu.mutation.StartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldStartDate,
		})
	}
	if value, ok := plu.mutation.EndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldEndDate,
		})
	}
	if value, ok := plu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projectlog.FieldCreatedAt,
		})
	}
	if plu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlog.ProjectTable,
			Columns: []string{projectlog.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlog.ProjectTable,
			Columns: []string{projectlog.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if plu.mutation.ProjectLogParticipantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogParticipantTable,
			Columns: []string{projectlog.ProjectLogParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogparticipant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plu.mutation.RemovedProjectLogParticipantIDs(); len(nodes) > 0 && !plu.mutation.ProjectLogParticipantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogParticipantTable,
			Columns: []string{projectlog.ProjectLogParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogparticipant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plu.mutation.ProjectLogParticipantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogParticipantTable,
			Columns: []string{projectlog.ProjectLogParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogparticipant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if plu.mutation.ProjectLogFeedbackCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogFeedbackTable,
			Columns: []string{projectlog.ProjectLogFeedbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogfeedback.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plu.mutation.RemovedProjectLogFeedbackIDs(); len(nodes) > 0 && !plu.mutation.ProjectLogFeedbackCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogFeedbackTable,
			Columns: []string{projectlog.ProjectLogFeedbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogfeedback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plu.mutation.ProjectLogFeedbackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogFeedbackTable,
			Columns: []string{projectlog.ProjectLogFeedbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogfeedback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, plu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectLogUpdateOne is the builder for updating a single ProjectLog entity.
type ProjectLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectLogMutation
}

// SetTitle sets the "title" field.
func (pluo *ProjectLogUpdateOne) SetTitle(s string) *ProjectLogUpdateOne {
	pluo.mutation.SetTitle(s)
	return pluo
}

// SetAuthor sets the "author" field.
func (pluo *ProjectLogUpdateOne) SetAuthor(s string) *ProjectLogUpdateOne {
	pluo.mutation.SetAuthor(s)
	return pluo
}

// SetContent sets the "content" field.
func (pluo *ProjectLogUpdateOne) SetContent(s string) *ProjectLogUpdateOne {
	pluo.mutation.SetContent(s)
	return pluo
}

// SetStartDate sets the "start_date" field.
func (pluo *ProjectLogUpdateOne) SetStartDate(s string) *ProjectLogUpdateOne {
	pluo.mutation.SetStartDate(s)
	return pluo
}

// SetEndDate sets the "end_date" field.
func (pluo *ProjectLogUpdateOne) SetEndDate(s string) *ProjectLogUpdateOne {
	pluo.mutation.SetEndDate(s)
	return pluo
}

// SetCreatedAt sets the "created_at" field.
func (pluo *ProjectLogUpdateOne) SetCreatedAt(t time.Time) *ProjectLogUpdateOne {
	pluo.mutation.SetCreatedAt(t)
	return pluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pluo *ProjectLogUpdateOne) SetNillableCreatedAt(t *time.Time) *ProjectLogUpdateOne {
	if t != nil {
		pluo.SetCreatedAt(*t)
	}
	return pluo
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (pluo *ProjectLogUpdateOne) SetProjectID(id int) *ProjectLogUpdateOne {
	pluo.mutation.SetProjectID(id)
	return pluo
}

// SetProject sets the "project" edge to the Project entity.
func (pluo *ProjectLogUpdateOne) SetProject(p *Project) *ProjectLogUpdateOne {
	return pluo.SetProjectID(p.ID)
}

// AddProjectLogParticipantIDs adds the "project_log_participant" edge to the ProjectLogParticipant entity by IDs.
func (pluo *ProjectLogUpdateOne) AddProjectLogParticipantIDs(ids ...int) *ProjectLogUpdateOne {
	pluo.mutation.AddProjectLogParticipantIDs(ids...)
	return pluo
}

// AddProjectLogParticipant adds the "project_log_participant" edges to the ProjectLogParticipant entity.
func (pluo *ProjectLogUpdateOne) AddProjectLogParticipant(p ...*ProjectLogParticipant) *ProjectLogUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pluo.AddProjectLogParticipantIDs(ids...)
}

// AddProjectLogFeedbackIDs adds the "project_log_feedback" edge to the ProjectLogFeedback entity by IDs.
func (pluo *ProjectLogUpdateOne) AddProjectLogFeedbackIDs(ids ...int) *ProjectLogUpdateOne {
	pluo.mutation.AddProjectLogFeedbackIDs(ids...)
	return pluo
}

// AddProjectLogFeedback adds the "project_log_feedback" edges to the ProjectLogFeedback entity.
func (pluo *ProjectLogUpdateOne) AddProjectLogFeedback(p ...*ProjectLogFeedback) *ProjectLogUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pluo.AddProjectLogFeedbackIDs(ids...)
}

// Mutation returns the ProjectLogMutation object of the builder.
func (pluo *ProjectLogUpdateOne) Mutation() *ProjectLogMutation {
	return pluo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (pluo *ProjectLogUpdateOne) ClearProject() *ProjectLogUpdateOne {
	pluo.mutation.ClearProject()
	return pluo
}

// ClearProjectLogParticipant clears all "project_log_participant" edges to the ProjectLogParticipant entity.
func (pluo *ProjectLogUpdateOne) ClearProjectLogParticipant() *ProjectLogUpdateOne {
	pluo.mutation.ClearProjectLogParticipant()
	return pluo
}

// RemoveProjectLogParticipantIDs removes the "project_log_participant" edge to ProjectLogParticipant entities by IDs.
func (pluo *ProjectLogUpdateOne) RemoveProjectLogParticipantIDs(ids ...int) *ProjectLogUpdateOne {
	pluo.mutation.RemoveProjectLogParticipantIDs(ids...)
	return pluo
}

// RemoveProjectLogParticipant removes "project_log_participant" edges to ProjectLogParticipant entities.
func (pluo *ProjectLogUpdateOne) RemoveProjectLogParticipant(p ...*ProjectLogParticipant) *ProjectLogUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pluo.RemoveProjectLogParticipantIDs(ids...)
}

// ClearProjectLogFeedback clears all "project_log_feedback" edges to the ProjectLogFeedback entity.
func (pluo *ProjectLogUpdateOne) ClearProjectLogFeedback() *ProjectLogUpdateOne {
	pluo.mutation.ClearProjectLogFeedback()
	return pluo
}

// RemoveProjectLogFeedbackIDs removes the "project_log_feedback" edge to ProjectLogFeedback entities by IDs.
func (pluo *ProjectLogUpdateOne) RemoveProjectLogFeedbackIDs(ids ...int) *ProjectLogUpdateOne {
	pluo.mutation.RemoveProjectLogFeedbackIDs(ids...)
	return pluo
}

// RemoveProjectLogFeedback removes "project_log_feedback" edges to ProjectLogFeedback entities.
func (pluo *ProjectLogUpdateOne) RemoveProjectLogFeedback(p ...*ProjectLogFeedback) *ProjectLogUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pluo.RemoveProjectLogFeedbackIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pluo *ProjectLogUpdateOne) Select(field string, fields ...string) *ProjectLogUpdateOne {
	pluo.fields = append([]string{field}, fields...)
	return pluo
}

// Save executes the query and returns the updated ProjectLog entity.
func (pluo *ProjectLogUpdateOne) Save(ctx context.Context) (*ProjectLog, error) {
	var (
		err  error
		node *ProjectLog
	)
	if len(pluo.hooks) == 0 {
		if err = pluo.check(); err != nil {
			return nil, err
		}
		node, err = pluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pluo.check(); err != nil {
				return nil, err
			}
			pluo.mutation = mutation
			node, err = pluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pluo.hooks) - 1; i >= 0; i-- {
			if pluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pluo *ProjectLogUpdateOne) SaveX(ctx context.Context) *ProjectLog {
	node, err := pluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pluo *ProjectLogUpdateOne) Exec(ctx context.Context) error {
	_, err := pluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pluo *ProjectLogUpdateOne) ExecX(ctx context.Context) {
	if err := pluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pluo *ProjectLogUpdateOne) check() error {
	if _, ok := pluo.mutation.ProjectID(); pluo.mutation.ProjectCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project\"")
	}
	return nil
}

func (pluo *ProjectLogUpdateOne) sqlSave(ctx context.Context) (_node *ProjectLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectlog.Table,
			Columns: projectlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectlog.FieldID,
			},
		},
	}
	id, ok := pluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProjectLog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectlog.FieldID)
		for _, f := range fields {
			if !projectlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projectlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pluo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldTitle,
		})
	}
	if value, ok := pluo.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldAuthor,
		})
	}
	if value, ok := pluo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldContent,
		})
	}
	if value, ok := pluo.mutation.StartDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldStartDate,
		})
	}
	if value, ok := pluo.mutation.EndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlog.FieldEndDate,
		})
	}
	if value, ok := pluo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: projectlog.FieldCreatedAt,
		})
	}
	if pluo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlog.ProjectTable,
			Columns: []string{projectlog.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pluo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlog.ProjectTable,
			Columns: []string{projectlog.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pluo.mutation.ProjectLogParticipantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogParticipantTable,
			Columns: []string{projectlog.ProjectLogParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogparticipant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pluo.mutation.RemovedProjectLogParticipantIDs(); len(nodes) > 0 && !pluo.mutation.ProjectLogParticipantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogParticipantTable,
			Columns: []string{projectlog.ProjectLogParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogparticipant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pluo.mutation.ProjectLogParticipantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogParticipantTable,
			Columns: []string{projectlog.ProjectLogParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogparticipant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pluo.mutation.ProjectLogFeedbackCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogFeedbackTable,
			Columns: []string{projectlog.ProjectLogFeedbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogfeedback.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pluo.mutation.RemovedProjectLogFeedbackIDs(); len(nodes) > 0 && !pluo.mutation.ProjectLogFeedbackCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogFeedbackTable,
			Columns: []string{projectlog.ProjectLogFeedbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogfeedback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pluo.mutation.ProjectLogFeedbackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectlog.ProjectLogFeedbackTable,
			Columns: []string{projectlog.ProjectLogFeedbackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlogfeedback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectLog{config: pluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
