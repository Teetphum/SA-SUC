// Code generated by entc, DO NOT EDIT.

package foodmenu

import (
	"github.com/Teeth/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FoodmenuName applies equality check predicate on the "foodmenu_name" field. It's identical to FoodmenuNameEQ.
func FoodmenuName(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuType applies equality check predicate on the "foodmenu_type" field. It's identical to FoodmenuTypeEQ.
func FoodmenuType(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuNameEQ applies the EQ predicate on the "foodmenu_name" field.
func FoodmenuNameEQ(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameNEQ applies the NEQ predicate on the "foodmenu_name" field.
func FoodmenuNameNEQ(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameIn applies the In predicate on the "foodmenu_name" field.
func FoodmenuNameIn(vs ...string) predicate.Foodmenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Foodmenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFoodmenuName), v...))
	})
}

// FoodmenuNameNotIn applies the NotIn predicate on the "foodmenu_name" field.
func FoodmenuNameNotIn(vs ...string) predicate.Foodmenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Foodmenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFoodmenuName), v...))
	})
}

// FoodmenuNameGT applies the GT predicate on the "foodmenu_name" field.
func FoodmenuNameGT(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameGTE applies the GTE predicate on the "foodmenu_name" field.
func FoodmenuNameGTE(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameLT applies the LT predicate on the "foodmenu_name" field.
func FoodmenuNameLT(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameLTE applies the LTE predicate on the "foodmenu_name" field.
func FoodmenuNameLTE(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameContains applies the Contains predicate on the "foodmenu_name" field.
func FoodmenuNameContains(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameHasPrefix applies the HasPrefix predicate on the "foodmenu_name" field.
func FoodmenuNameHasPrefix(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameHasSuffix applies the HasSuffix predicate on the "foodmenu_name" field.
func FoodmenuNameHasSuffix(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameEqualFold applies the EqualFold predicate on the "foodmenu_name" field.
func FoodmenuNameEqualFold(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuNameContainsFold applies the ContainsFold predicate on the "foodmenu_name" field.
func FoodmenuNameContainsFold(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFoodmenuName), v))
	})
}

// FoodmenuTypeEQ applies the EQ predicate on the "foodmenu_type" field.
func FoodmenuTypeEQ(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeNEQ applies the NEQ predicate on the "foodmenu_type" field.
func FoodmenuTypeNEQ(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeIn applies the In predicate on the "foodmenu_type" field.
func FoodmenuTypeIn(vs ...string) predicate.Foodmenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Foodmenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFoodmenuType), v...))
	})
}

// FoodmenuTypeNotIn applies the NotIn predicate on the "foodmenu_type" field.
func FoodmenuTypeNotIn(vs ...string) predicate.Foodmenu {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Foodmenu(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFoodmenuType), v...))
	})
}

// FoodmenuTypeGT applies the GT predicate on the "foodmenu_type" field.
func FoodmenuTypeGT(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeGTE applies the GTE predicate on the "foodmenu_type" field.
func FoodmenuTypeGTE(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeLT applies the LT predicate on the "foodmenu_type" field.
func FoodmenuTypeLT(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeLTE applies the LTE predicate on the "foodmenu_type" field.
func FoodmenuTypeLTE(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeContains applies the Contains predicate on the "foodmenu_type" field.
func FoodmenuTypeContains(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeHasPrefix applies the HasPrefix predicate on the "foodmenu_type" field.
func FoodmenuTypeHasPrefix(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeHasSuffix applies the HasSuffix predicate on the "foodmenu_type" field.
func FoodmenuTypeHasSuffix(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeEqualFold applies the EqualFold predicate on the "foodmenu_type" field.
func FoodmenuTypeEqualFold(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFoodmenuType), v))
	})
}

// FoodmenuTypeContainsFold applies the ContainsFold predicate on the "foodmenu_type" field.
func FoodmenuTypeContainsFold(v string) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFoodmenuType), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEatinghistory applies the HasEdge predicate on the "eatinghistory" edge.
func HasEatinghistory() predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EatinghistoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EatinghistoryTable, EatinghistoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEatinghistoryWith applies the HasEdge predicate on the "eatinghistory" edge with a given conditions (other predicates).
func HasEatinghistoryWith(preds ...predicate.Eatinghistory) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EatinghistoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EatinghistoryTable, EatinghistoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Foodmenu) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Foodmenu) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
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
func Not(p predicate.Foodmenu) predicate.Foodmenu {
	return predicate.Foodmenu(func(s *sql.Selector) {
		p(s.Not())
	})
}
