// Code generated by entc, DO NOT EDIT.

package clubmember

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ClubID applies equality check predicate on the "club_id" field. It's identical to ClubIDEQ.
func ClubID(v int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClubID), v))
	})
}

// StudentID applies equality check predicate on the "student_id" field. It's identical to StudentIDEQ.
func StudentID(v int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudentID), v))
	})
}

// ClubIDEQ applies the EQ predicate on the "club_id" field.
func ClubIDEQ(v int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClubID), v))
	})
}

// ClubIDNEQ applies the NEQ predicate on the "club_id" field.
func ClubIDNEQ(v int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClubID), v))
	})
}

// ClubIDIn applies the In predicate on the "club_id" field.
func ClubIDIn(vs ...int) predicate.ClubMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClubMember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClubID), v...))
	})
}

// ClubIDNotIn applies the NotIn predicate on the "club_id" field.
func ClubIDNotIn(vs ...int) predicate.ClubMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClubMember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClubID), v...))
	})
}

// StudentIDEQ applies the EQ predicate on the "student_id" field.
func StudentIDEQ(v int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudentID), v))
	})
}

// StudentIDNEQ applies the NEQ predicate on the "student_id" field.
func StudentIDNEQ(v int) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStudentID), v))
	})
}

// StudentIDIn applies the In predicate on the "student_id" field.
func StudentIDIn(vs ...int) predicate.ClubMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClubMember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStudentID), v...))
	})
}

// StudentIDNotIn applies the NotIn predicate on the "student_id" field.
func StudentIDNotIn(vs ...int) predicate.ClubMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClubMember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStudentID), v...))
	})
}

// HasStudent applies the HasEdge predicate on the "student" edge.
func HasStudent() predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentWith applies the HasEdge predicate on the "student" edge with a given conditions (other predicates).
func HasStudentWith(preds ...predicate.Student) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClub applies the HasEdge predicate on the "club" edge.
func HasClub() predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClubTable, ClubColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClubWith applies the HasEdge predicate on the "club" edge with a given conditions (other predicates).
func HasClubWith(preds ...predicate.Club) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClubInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClubTable, ClubColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ClubMember) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ClubMember) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
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
func Not(p predicate.ClubMember) predicate.ClubMember {
	return predicate.ClubMember(func(s *sql.Selector) {
		p(s.Not())
	})
}