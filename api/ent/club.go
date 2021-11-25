// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Wanted-Linx/linx-backend/api/ent/club"
	"github.com/Wanted-Linx/linx-backend/api/ent/student"
)

// Club is the model entity for the Club schema.
type Club struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Organization holds the value of the "organization" field.
	Organization string `json:"organization,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ProfileImage holds the value of the "profile_image" field.
	ProfileImage *string `json:"profile_image,omitempty"`
	// ProfileLink holds the value of the "profile_link" field.
	ProfileLink *string `json:"profile_link,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClubQuery when eager-loading is set.
	Edges        ClubEdges `json:"edges"`
	student_club *int
}

// ClubEdges holds the relations/edges for other nodes in the graph.
type ClubEdges struct {
	// Leader holds the value of the leader edge.
	Leader *Student `json:"leader,omitempty"`
	// ClubMember holds the value of the club_member edge.
	ClubMember []*ClubMember `json:"club_member,omitempty"`
	// Project holds the value of the project edge.
	Project []*Project `json:"project,omitempty"`
	// ProjectClub holds the value of the project_club edge.
	ProjectClub []*ProjectClub `json:"project_club,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// LeaderOrErr returns the Leader value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClubEdges) LeaderOrErr() (*Student, error) {
	if e.loadedTypes[0] {
		if e.Leader == nil {
			// The edge leader was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: student.Label}
		}
		return e.Leader, nil
	}
	return nil, &NotLoadedError{edge: "leader"}
}

// ClubMemberOrErr returns the ClubMember value or an error if the edge
// was not loaded in eager-loading.
func (e ClubEdges) ClubMemberOrErr() ([]*ClubMember, error) {
	if e.loadedTypes[1] {
		return e.ClubMember, nil
	}
	return nil, &NotLoadedError{edge: "club_member"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading.
func (e ClubEdges) ProjectOrErr() ([]*Project, error) {
	if e.loadedTypes[2] {
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// ProjectClubOrErr returns the ProjectClub value or an error if the edge
// was not loaded in eager-loading.
func (e ClubEdges) ProjectClubOrErr() ([]*ProjectClub, error) {
	if e.loadedTypes[3] {
		return e.ProjectClub, nil
	}
	return nil, &NotLoadedError{edge: "project_club"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Club) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case club.FieldID:
			values[i] = new(sql.NullInt64)
		case club.FieldName, club.FieldOrganization, club.FieldDescription, club.FieldProfileImage, club.FieldProfileLink:
			values[i] = new(sql.NullString)
		case club.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case club.ForeignKeys[0]: // student_club
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Club", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Club fields.
func (c *Club) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case club.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case club.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case club.FieldOrganization:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization", values[i])
			} else if value.Valid {
				c.Organization = value.String
			}
		case club.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case club.FieldProfileImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_image", values[i])
			} else if value.Valid {
				c.ProfileImage = new(string)
				*c.ProfileImage = value.String
			}
		case club.FieldProfileLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_link", values[i])
			} else if value.Valid {
				c.ProfileLink = new(string)
				*c.ProfileLink = value.String
			}
		case club.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case club.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field student_club", value)
			} else if value.Valid {
				c.student_club = new(int)
				*c.student_club = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryLeader queries the "leader" edge of the Club entity.
func (c *Club) QueryLeader() *StudentQuery {
	return (&ClubClient{config: c.config}).QueryLeader(c)
}

// QueryClubMember queries the "club_member" edge of the Club entity.
func (c *Club) QueryClubMember() *ClubMemberQuery {
	return (&ClubClient{config: c.config}).QueryClubMember(c)
}

// QueryProject queries the "project" edge of the Club entity.
func (c *Club) QueryProject() *ProjectQuery {
	return (&ClubClient{config: c.config}).QueryProject(c)
}

// QueryProjectClub queries the "project_club" edge of the Club entity.
func (c *Club) QueryProjectClub() *ProjectClubQuery {
	return (&ClubClient{config: c.config}).QueryProjectClub(c)
}

// Update returns a builder for updating this Club.
// Note that you need to call Club.Unwrap() before calling this method if this Club
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Club) Update() *ClubUpdateOne {
	return (&ClubClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Club entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Club) Unwrap() *Club {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Club is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Club) String() string {
	var builder strings.Builder
	builder.WriteString("Club(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", organization=")
	builder.WriteString(c.Organization)
	builder.WriteString(", description=")
	builder.WriteString(c.Description)
	if v := c.ProfileImage; v != nil {
		builder.WriteString(", profile_image=")
		builder.WriteString(*v)
	}
	if v := c.ProfileLink; v != nil {
		builder.WriteString(", profile_link=")
		builder.WriteString(*v)
	}
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Clubs is a parsable slice of Club.
type Clubs []*Club

func (c Clubs) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
