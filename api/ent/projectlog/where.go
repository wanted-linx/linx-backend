// Code generated by entc, DO NOT EDIT.

package projectlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndDate), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthor), v))
	})
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAuthor), v...))
	})
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAuthor), v...))
	})
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthor), v))
	})
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthor), v))
	})
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthor), v))
	})
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthor), v))
	})
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthor), v))
	})
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthor), v))
	})
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthor), v))
	})
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthor), v))
	})
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthor), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartDate), v))
	})
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartDate), v...))
	})
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartDate), v...))
	})
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartDate), v))
	})
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartDate), v))
	})
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartDate), v))
	})
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartDate), v))
	})
}

// StartDateContains applies the Contains predicate on the "start_date" field.
func StartDateContains(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStartDate), v))
	})
}

// StartDateHasPrefix applies the HasPrefix predicate on the "start_date" field.
func StartDateHasPrefix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStartDate), v))
	})
}

// StartDateHasSuffix applies the HasSuffix predicate on the "start_date" field.
func StartDateHasSuffix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStartDate), v))
	})
}

// StartDateEqualFold applies the EqualFold predicate on the "start_date" field.
func StartDateEqualFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStartDate), v))
	})
}

// StartDateContainsFold applies the ContainsFold predicate on the "start_date" field.
func StartDateContainsFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStartDate), v))
	})
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndDate), v))
	})
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndDate), v))
	})
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndDate), v...))
	})
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...string) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndDate), v...))
	})
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndDate), v))
	})
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndDate), v))
	})
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndDate), v))
	})
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndDate), v))
	})
}

// EndDateContains applies the Contains predicate on the "end_date" field.
func EndDateContains(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEndDate), v))
	})
}

// EndDateHasPrefix applies the HasPrefix predicate on the "end_date" field.
func EndDateHasPrefix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEndDate), v))
	})
}

// EndDateHasSuffix applies the HasSuffix predicate on the "end_date" field.
func EndDateHasSuffix(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEndDate), v))
	})
}

// EndDateEqualFold applies the EqualFold predicate on the "end_date" field.
func EndDateEqualFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEndDate), v))
	})
}

// EndDateContainsFold applies the ContainsFold predicate on the "end_date" field.
func EndDateContainsFold(v string) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEndDate), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProjectLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProjectLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjectLogParticipant applies the HasEdge predicate on the "project_log_participant" edge.
func HasProjectLogParticipant() predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectLogParticipantTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectLogParticipantTable, ProjectLogParticipantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectLogParticipantWith applies the HasEdge predicate on the "project_log_participant" edge with a given conditions (other predicates).
func HasProjectLogParticipantWith(preds ...predicate.ProjectLogParticipant) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectLogParticipantInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectLogParticipantTable, ProjectLogParticipantColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjectLogFeedback applies the HasEdge predicate on the "project_log_feedback" edge.
func HasProjectLogFeedback() predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectLogFeedbackTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectLogFeedbackTable, ProjectLogFeedbackColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectLogFeedbackWith applies the HasEdge predicate on the "project_log_feedback" edge with a given conditions (other predicates).
func HasProjectLogFeedbackWith(preds ...predicate.ProjectLogFeedback) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectLogFeedbackInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectLogFeedbackTable, ProjectLogFeedbackColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProjectLog) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProjectLog) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProjectLog) predicate.ProjectLog {
	return predicate.ProjectLog(func(s *sql.Selector) {
		p(s.Not())
	})
}