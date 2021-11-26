// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Wanted-Linx/linx-backend/api/ent/club"
	"github.com/Wanted-Linx/linx-backend/api/ent/project"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectclub"
)

// ProjectClub is the model entity for the ProjectClub schema.
type ProjectClub struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ClubID holds the value of the "club_id" field.
	ClubID int `json:"club_id,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID int `json:"project_id,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate string `json:"start_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectClubQuery when eager-loading is set.
	Edges ProjectClubEdges `json:"edges"`
}

// ProjectClubEdges holds the relations/edges for other nodes in the graph.
type ProjectClubEdges struct {
	// Club holds the value of the club edge.
	Club *Club `json:"club,omitempty"`
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// ProjectLog holds the value of the project_log edge.
	ProjectLog []*ProjectLog `json:"project_log,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ClubOrErr returns the Club value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectClubEdges) ClubOrErr() (*Club, error) {
	if e.loadedTypes[0] {
		if e.Club == nil {
			// The edge club was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: club.Label}
		}
		return e.Club, nil
	}
	return nil, &NotLoadedError{edge: "club"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectClubEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[1] {
		if e.Project == nil {
			// The edge project was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// ProjectLogOrErr returns the ProjectLog value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectClubEdges) ProjectLogOrErr() ([]*ProjectLog, error) {
	if e.loadedTypes[2] {
		return e.ProjectLog, nil
	}
	return nil, &NotLoadedError{edge: "project_log"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProjectClub) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case projectclub.FieldID, projectclub.FieldClubID, projectclub.FieldProjectID:
			values[i] = new(sql.NullInt64)
		case projectclub.FieldStartDate:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProjectClub", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProjectClub fields.
func (pc *ProjectClub) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case projectclub.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int(value.Int64)
		case projectclub.FieldClubID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field club_id", values[i])
			} else if value.Valid {
				pc.ClubID = int(value.Int64)
			}
		case projectclub.FieldProjectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value.Valid {
				pc.ProjectID = int(value.Int64)
			}
		case projectclub.FieldStartDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				pc.StartDate = value.String
			}
		}
	}
	return nil
}

// QueryClub queries the "club" edge of the ProjectClub entity.
func (pc *ProjectClub) QueryClub() *ClubQuery {
	return (&ProjectClubClient{config: pc.config}).QueryClub(pc)
}

// QueryProject queries the "project" edge of the ProjectClub entity.
func (pc *ProjectClub) QueryProject() *ProjectQuery {
	return (&ProjectClubClient{config: pc.config}).QueryProject(pc)
}

// QueryProjectLog queries the "project_log" edge of the ProjectClub entity.
func (pc *ProjectClub) QueryProjectLog() *ProjectLogQuery {
	return (&ProjectClubClient{config: pc.config}).QueryProjectLog(pc)
}

// Update returns a builder for updating this ProjectClub.
// Note that you need to call ProjectClub.Unwrap() before calling this method if this ProjectClub
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *ProjectClub) Update() *ProjectClubUpdateOne {
	return (&ProjectClubClient{config: pc.config}).UpdateOne(pc)
}

// Unwrap unwraps the ProjectClub entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *ProjectClub) Unwrap() *ProjectClub {
	tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProjectClub is not a transactional entity")
	}
	pc.config.driver = tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *ProjectClub) String() string {
	var builder strings.Builder
	builder.WriteString("ProjectClub(")
	builder.WriteString(fmt.Sprintf("id=%v", pc.ID))
	builder.WriteString(", club_id=")
	builder.WriteString(fmt.Sprintf("%v", pc.ClubID))
	builder.WriteString(", project_id=")
	builder.WriteString(fmt.Sprintf("%v", pc.ProjectID))
	builder.WriteString(", start_date=")
	builder.WriteString(pc.StartDate)
	builder.WriteByte(')')
	return builder.String()
}

// ProjectClubs is a parsable slice of ProjectClub.
type ProjectClubs []*ProjectClub

func (pc ProjectClubs) config(cfg config) {
	for _i := range pc {
		pc[_i].config = cfg
	}
}
