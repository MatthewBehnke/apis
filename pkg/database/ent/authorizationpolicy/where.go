// Code generated by ent, DO NOT EDIT.

package authorizationpolicy

import (
	"entgo.io/ent/dialect/sql"

	"github.com/MatthewBehnke/apis/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Ptype applies equality check predicate on the "Ptype" field. It's identical to PtypeEQ.
func Ptype(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPtype), v))
	})
}

// V0 applies equality check predicate on the "V0" field. It's identical to V0EQ.
func V0(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV0), v))
	})
}

// V1 applies equality check predicate on the "V1" field. It's identical to V1EQ.
func V1(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV1), v))
	})
}

// V2 applies equality check predicate on the "V2" field. It's identical to V2EQ.
func V2(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV2), v))
	})
}

// V3 applies equality check predicate on the "V3" field. It's identical to V3EQ.
func V3(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV3), v))
	})
}

// V4 applies equality check predicate on the "V4" field. It's identical to V4EQ.
func V4(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV4), v))
	})
}

// V5 applies equality check predicate on the "V5" field. It's identical to V5EQ.
func V5(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV5), v))
	})
}

// PtypeEQ applies the EQ predicate on the "Ptype" field.
func PtypeEQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPtype), v))
	})
}

// PtypeNEQ applies the NEQ predicate on the "Ptype" field.
func PtypeNEQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPtype), v))
	})
}

// PtypeIn applies the In predicate on the "Ptype" field.
func PtypeIn(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPtype), v...))
	})
}

// PtypeNotIn applies the NotIn predicate on the "Ptype" field.
func PtypeNotIn(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPtype), v...))
	})
}

// PtypeGT applies the GT predicate on the "Ptype" field.
func PtypeGT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPtype), v))
	})
}

// PtypeGTE applies the GTE predicate on the "Ptype" field.
func PtypeGTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPtype), v))
	})
}

// PtypeLT applies the LT predicate on the "Ptype" field.
func PtypeLT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPtype), v))
	})
}

// PtypeLTE applies the LTE predicate on the "Ptype" field.
func PtypeLTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPtype), v))
	})
}

// PtypeContains applies the Contains predicate on the "Ptype" field.
func PtypeContains(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPtype), v))
	})
}

// PtypeHasPrefix applies the HasPrefix predicate on the "Ptype" field.
func PtypeHasPrefix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPtype), v))
	})
}

// PtypeHasSuffix applies the HasSuffix predicate on the "Ptype" field.
func PtypeHasSuffix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPtype), v))
	})
}

// PtypeEqualFold applies the EqualFold predicate on the "Ptype" field.
func PtypeEqualFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPtype), v))
	})
}

// PtypeContainsFold applies the ContainsFold predicate on the "Ptype" field.
func PtypeContainsFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPtype), v))
	})
}

// V0EQ applies the EQ predicate on the "V0" field.
func V0EQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV0), v))
	})
}

// V0NEQ applies the NEQ predicate on the "V0" field.
func V0NEQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldV0), v))
	})
}

// V0In applies the In predicate on the "V0" field.
func V0In(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldV0), v...))
	})
}

// V0NotIn applies the NotIn predicate on the "V0" field.
func V0NotIn(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldV0), v...))
	})
}

// V0GT applies the GT predicate on the "V0" field.
func V0GT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldV0), v))
	})
}

// V0GTE applies the GTE predicate on the "V0" field.
func V0GTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldV0), v))
	})
}

// V0LT applies the LT predicate on the "V0" field.
func V0LT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldV0), v))
	})
}

// V0LTE applies the LTE predicate on the "V0" field.
func V0LTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldV0), v))
	})
}

// V0Contains applies the Contains predicate on the "V0" field.
func V0Contains(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldV0), v))
	})
}

// V0HasPrefix applies the HasPrefix predicate on the "V0" field.
func V0HasPrefix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldV0), v))
	})
}

// V0HasSuffix applies the HasSuffix predicate on the "V0" field.
func V0HasSuffix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldV0), v))
	})
}

// V0EqualFold applies the EqualFold predicate on the "V0" field.
func V0EqualFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldV0), v))
	})
}

// V0ContainsFold applies the ContainsFold predicate on the "V0" field.
func V0ContainsFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldV0), v))
	})
}

// V1EQ applies the EQ predicate on the "V1" field.
func V1EQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV1), v))
	})
}

// V1NEQ applies the NEQ predicate on the "V1" field.
func V1NEQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldV1), v))
	})
}

// V1In applies the In predicate on the "V1" field.
func V1In(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldV1), v...))
	})
}

// V1NotIn applies the NotIn predicate on the "V1" field.
func V1NotIn(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldV1), v...))
	})
}

// V1GT applies the GT predicate on the "V1" field.
func V1GT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldV1), v))
	})
}

// V1GTE applies the GTE predicate on the "V1" field.
func V1GTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldV1), v))
	})
}

// V1LT applies the LT predicate on the "V1" field.
func V1LT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldV1), v))
	})
}

// V1LTE applies the LTE predicate on the "V1" field.
func V1LTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldV1), v))
	})
}

// V1Contains applies the Contains predicate on the "V1" field.
func V1Contains(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldV1), v))
	})
}

// V1HasPrefix applies the HasPrefix predicate on the "V1" field.
func V1HasPrefix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldV1), v))
	})
}

// V1HasSuffix applies the HasSuffix predicate on the "V1" field.
func V1HasSuffix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldV1), v))
	})
}

// V1EqualFold applies the EqualFold predicate on the "V1" field.
func V1EqualFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldV1), v))
	})
}

// V1ContainsFold applies the ContainsFold predicate on the "V1" field.
func V1ContainsFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldV1), v))
	})
}

// V2EQ applies the EQ predicate on the "V2" field.
func V2EQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV2), v))
	})
}

// V2NEQ applies the NEQ predicate on the "V2" field.
func V2NEQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldV2), v))
	})
}

// V2In applies the In predicate on the "V2" field.
func V2In(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldV2), v...))
	})
}

// V2NotIn applies the NotIn predicate on the "V2" field.
func V2NotIn(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldV2), v...))
	})
}

// V2GT applies the GT predicate on the "V2" field.
func V2GT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldV2), v))
	})
}

// V2GTE applies the GTE predicate on the "V2" field.
func V2GTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldV2), v))
	})
}

// V2LT applies the LT predicate on the "V2" field.
func V2LT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldV2), v))
	})
}

// V2LTE applies the LTE predicate on the "V2" field.
func V2LTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldV2), v))
	})
}

// V2Contains applies the Contains predicate on the "V2" field.
func V2Contains(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldV2), v))
	})
}

// V2HasPrefix applies the HasPrefix predicate on the "V2" field.
func V2HasPrefix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldV2), v))
	})
}

// V2HasSuffix applies the HasSuffix predicate on the "V2" field.
func V2HasSuffix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldV2), v))
	})
}

// V2EqualFold applies the EqualFold predicate on the "V2" field.
func V2EqualFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldV2), v))
	})
}

// V2ContainsFold applies the ContainsFold predicate on the "V2" field.
func V2ContainsFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldV2), v))
	})
}

// V3EQ applies the EQ predicate on the "V3" field.
func V3EQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV3), v))
	})
}

// V3NEQ applies the NEQ predicate on the "V3" field.
func V3NEQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldV3), v))
	})
}

// V3In applies the In predicate on the "V3" field.
func V3In(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldV3), v...))
	})
}

// V3NotIn applies the NotIn predicate on the "V3" field.
func V3NotIn(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldV3), v...))
	})
}

// V3GT applies the GT predicate on the "V3" field.
func V3GT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldV3), v))
	})
}

// V3GTE applies the GTE predicate on the "V3" field.
func V3GTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldV3), v))
	})
}

// V3LT applies the LT predicate on the "V3" field.
func V3LT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldV3), v))
	})
}

// V3LTE applies the LTE predicate on the "V3" field.
func V3LTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldV3), v))
	})
}

// V3Contains applies the Contains predicate on the "V3" field.
func V3Contains(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldV3), v))
	})
}

// V3HasPrefix applies the HasPrefix predicate on the "V3" field.
func V3HasPrefix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldV3), v))
	})
}

// V3HasSuffix applies the HasSuffix predicate on the "V3" field.
func V3HasSuffix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldV3), v))
	})
}

// V3EqualFold applies the EqualFold predicate on the "V3" field.
func V3EqualFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldV3), v))
	})
}

// V3ContainsFold applies the ContainsFold predicate on the "V3" field.
func V3ContainsFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldV3), v))
	})
}

// V4EQ applies the EQ predicate on the "V4" field.
func V4EQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV4), v))
	})
}

// V4NEQ applies the NEQ predicate on the "V4" field.
func V4NEQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldV4), v))
	})
}

// V4In applies the In predicate on the "V4" field.
func V4In(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldV4), v...))
	})
}

// V4NotIn applies the NotIn predicate on the "V4" field.
func V4NotIn(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldV4), v...))
	})
}

// V4GT applies the GT predicate on the "V4" field.
func V4GT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldV4), v))
	})
}

// V4GTE applies the GTE predicate on the "V4" field.
func V4GTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldV4), v))
	})
}

// V4LT applies the LT predicate on the "V4" field.
func V4LT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldV4), v))
	})
}

// V4LTE applies the LTE predicate on the "V4" field.
func V4LTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldV4), v))
	})
}

// V4Contains applies the Contains predicate on the "V4" field.
func V4Contains(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldV4), v))
	})
}

// V4HasPrefix applies the HasPrefix predicate on the "V4" field.
func V4HasPrefix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldV4), v))
	})
}

// V4HasSuffix applies the HasSuffix predicate on the "V4" field.
func V4HasSuffix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldV4), v))
	})
}

// V4EqualFold applies the EqualFold predicate on the "V4" field.
func V4EqualFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldV4), v))
	})
}

// V4ContainsFold applies the ContainsFold predicate on the "V4" field.
func V4ContainsFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldV4), v))
	})
}

// V5EQ applies the EQ predicate on the "V5" field.
func V5EQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldV5), v))
	})
}

// V5NEQ applies the NEQ predicate on the "V5" field.
func V5NEQ(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldV5), v))
	})
}

// V5In applies the In predicate on the "V5" field.
func V5In(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldV5), v...))
	})
}

// V5NotIn applies the NotIn predicate on the "V5" field.
func V5NotIn(vs ...string) predicate.AuthorizationPolicy {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldV5), v...))
	})
}

// V5GT applies the GT predicate on the "V5" field.
func V5GT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldV5), v))
	})
}

// V5GTE applies the GTE predicate on the "V5" field.
func V5GTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldV5), v))
	})
}

// V5LT applies the LT predicate on the "V5" field.
func V5LT(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldV5), v))
	})
}

// V5LTE applies the LTE predicate on the "V5" field.
func V5LTE(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldV5), v))
	})
}

// V5Contains applies the Contains predicate on the "V5" field.
func V5Contains(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldV5), v))
	})
}

// V5HasPrefix applies the HasPrefix predicate on the "V5" field.
func V5HasPrefix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldV5), v))
	})
}

// V5HasSuffix applies the HasSuffix predicate on the "V5" field.
func V5HasSuffix(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldV5), v))
	})
}

// V5EqualFold applies the EqualFold predicate on the "V5" field.
func V5EqualFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldV5), v))
	})
}

// V5ContainsFold applies the ContainsFold predicate on the "V5" field.
func V5ContainsFold(v string) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldV5), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AuthorizationPolicy) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AuthorizationPolicy) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
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
func Not(p predicate.AuthorizationPolicy) predicate.AuthorizationPolicy {
	return predicate.AuthorizationPolicy(func(s *sql.Selector) {
		p(s.Not())
	})
}
