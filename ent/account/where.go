// Code generated by entc, DO NOT EDIT.

package account

import (
	"entgo.io/ent/dialect/sql"
	"github.com/yonidavidson/cockroachent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBalance), v))
	})
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBalance), v))
	})
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBalance), v))
	})
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...int) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBalance), v...))
	})
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...int) predicate.Account {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Account(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBalance), v...))
	})
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBalance), v))
	})
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBalance), v))
	})
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBalance), v))
	})
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v int) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBalance), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}
