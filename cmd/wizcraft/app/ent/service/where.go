// Code generated by ent, DO NOT EDIT.

package service

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Service {
	return predicate.Service(sql.FieldLTE(FieldID, id))
}

// HasDatabases applies the HasEdge predicate on the "databases" edge.
func HasDatabases() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DatabasesTable, DatabasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDatabasesWith applies the HasEdge predicate on the "databases" edge with a given conditions (other predicates).
func HasDatabasesWith(preds ...predicate.Database) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newDatabasesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasApispec applies the HasEdge predicate on the "apispec" edge.
func HasApispec() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ApispecTable, ApispecColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApispecWith applies the HasEdge predicate on the "apispec" edge with a given conditions (other predicates).
func HasApispecWith(preds ...predicate.APISpec) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newApispecStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGeneralspec applies the HasEdge predicate on the "generalspec" edge.
func HasGeneralspec() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, GeneralspecTable, GeneralspecColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGeneralspecWith applies the HasEdge predicate on the "generalspec" edge with a given conditions (other predicates).
func HasGeneralspecWith(preds ...predicate.GeneralSpec) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := newGeneralspecStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Service) predicate.Service {
	return predicate.Service(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Service) predicate.Service {
	return predicate.Service(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Service) predicate.Service {
	return predicate.Service(sql.NotPredicates(p))
}
