// Code generated by ent, DO NOT EDIT.

package times

import (
	"service/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Times {
	return predicate.Times(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Times {
	return predicate.Times(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Times {
	return predicate.Times(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Times {
	return predicate.Times(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Times {
	return predicate.Times(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Times {
	return predicate.Times(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Times {
	return predicate.Times(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Times {
	return predicate.Times(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Times {
	return predicate.Times(sql.FieldLTE(FieldID, id))
}

// UserId applies equality check predicate on the "userId" field. It's identical to UserIdEQ.
func UserId(v int64) predicate.Times {
	return predicate.Times(sql.FieldEQ(FieldUserId, v))
}

// TimeSpace applies equality check predicate on the "timeSpace" field. It's identical to TimeSpaceEQ.
func TimeSpace(v string) predicate.Times {
	return predicate.Times(sql.FieldEQ(FieldTimeSpace, v))
}

// StartTime applies equality check predicate on the "startTime" field. It's identical to StartTimeEQ.
func StartTime(v string) predicate.Times {
	return predicate.Times(sql.FieldEQ(FieldStartTime, v))
}

// UserIdEQ applies the EQ predicate on the "userId" field.
func UserIdEQ(v int64) predicate.Times {
	return predicate.Times(sql.FieldEQ(FieldUserId, v))
}

// UserIdNEQ applies the NEQ predicate on the "userId" field.
func UserIdNEQ(v int64) predicate.Times {
	return predicate.Times(sql.FieldNEQ(FieldUserId, v))
}

// UserIdIn applies the In predicate on the "userId" field.
func UserIdIn(vs ...int64) predicate.Times {
	return predicate.Times(sql.FieldIn(FieldUserId, vs...))
}

// UserIdNotIn applies the NotIn predicate on the "userId" field.
func UserIdNotIn(vs ...int64) predicate.Times {
	return predicate.Times(sql.FieldNotIn(FieldUserId, vs...))
}

// UserIdGT applies the GT predicate on the "userId" field.
func UserIdGT(v int64) predicate.Times {
	return predicate.Times(sql.FieldGT(FieldUserId, v))
}

// UserIdGTE applies the GTE predicate on the "userId" field.
func UserIdGTE(v int64) predicate.Times {
	return predicate.Times(sql.FieldGTE(FieldUserId, v))
}

// UserIdLT applies the LT predicate on the "userId" field.
func UserIdLT(v int64) predicate.Times {
	return predicate.Times(sql.FieldLT(FieldUserId, v))
}

// UserIdLTE applies the LTE predicate on the "userId" field.
func UserIdLTE(v int64) predicate.Times {
	return predicate.Times(sql.FieldLTE(FieldUserId, v))
}

// TimeSpaceEQ applies the EQ predicate on the "timeSpace" field.
func TimeSpaceEQ(v string) predicate.Times {
	return predicate.Times(sql.FieldEQ(FieldTimeSpace, v))
}

// TimeSpaceNEQ applies the NEQ predicate on the "timeSpace" field.
func TimeSpaceNEQ(v string) predicate.Times {
	return predicate.Times(sql.FieldNEQ(FieldTimeSpace, v))
}

// TimeSpaceIn applies the In predicate on the "timeSpace" field.
func TimeSpaceIn(vs ...string) predicate.Times {
	return predicate.Times(sql.FieldIn(FieldTimeSpace, vs...))
}

// TimeSpaceNotIn applies the NotIn predicate on the "timeSpace" field.
func TimeSpaceNotIn(vs ...string) predicate.Times {
	return predicate.Times(sql.FieldNotIn(FieldTimeSpace, vs...))
}

// TimeSpaceGT applies the GT predicate on the "timeSpace" field.
func TimeSpaceGT(v string) predicate.Times {
	return predicate.Times(sql.FieldGT(FieldTimeSpace, v))
}

// TimeSpaceGTE applies the GTE predicate on the "timeSpace" field.
func TimeSpaceGTE(v string) predicate.Times {
	return predicate.Times(sql.FieldGTE(FieldTimeSpace, v))
}

// TimeSpaceLT applies the LT predicate on the "timeSpace" field.
func TimeSpaceLT(v string) predicate.Times {
	return predicate.Times(sql.FieldLT(FieldTimeSpace, v))
}

// TimeSpaceLTE applies the LTE predicate on the "timeSpace" field.
func TimeSpaceLTE(v string) predicate.Times {
	return predicate.Times(sql.FieldLTE(FieldTimeSpace, v))
}

// TimeSpaceContains applies the Contains predicate on the "timeSpace" field.
func TimeSpaceContains(v string) predicate.Times {
	return predicate.Times(sql.FieldContains(FieldTimeSpace, v))
}

// TimeSpaceHasPrefix applies the HasPrefix predicate on the "timeSpace" field.
func TimeSpaceHasPrefix(v string) predicate.Times {
	return predicate.Times(sql.FieldHasPrefix(FieldTimeSpace, v))
}

// TimeSpaceHasSuffix applies the HasSuffix predicate on the "timeSpace" field.
func TimeSpaceHasSuffix(v string) predicate.Times {
	return predicate.Times(sql.FieldHasSuffix(FieldTimeSpace, v))
}

// TimeSpaceEqualFold applies the EqualFold predicate on the "timeSpace" field.
func TimeSpaceEqualFold(v string) predicate.Times {
	return predicate.Times(sql.FieldEqualFold(FieldTimeSpace, v))
}

// TimeSpaceContainsFold applies the ContainsFold predicate on the "timeSpace" field.
func TimeSpaceContainsFold(v string) predicate.Times {
	return predicate.Times(sql.FieldContainsFold(FieldTimeSpace, v))
}

// StartTimeEQ applies the EQ predicate on the "startTime" field.
func StartTimeEQ(v string) predicate.Times {
	return predicate.Times(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "startTime" field.
func StartTimeNEQ(v string) predicate.Times {
	return predicate.Times(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "startTime" field.
func StartTimeIn(vs ...string) predicate.Times {
	return predicate.Times(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "startTime" field.
func StartTimeNotIn(vs ...string) predicate.Times {
	return predicate.Times(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "startTime" field.
func StartTimeGT(v string) predicate.Times {
	return predicate.Times(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "startTime" field.
func StartTimeGTE(v string) predicate.Times {
	return predicate.Times(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "startTime" field.
func StartTimeLT(v string) predicate.Times {
	return predicate.Times(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "startTime" field.
func StartTimeLTE(v string) predicate.Times {
	return predicate.Times(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeContains applies the Contains predicate on the "startTime" field.
func StartTimeContains(v string) predicate.Times {
	return predicate.Times(sql.FieldContains(FieldStartTime, v))
}

// StartTimeHasPrefix applies the HasPrefix predicate on the "startTime" field.
func StartTimeHasPrefix(v string) predicate.Times {
	return predicate.Times(sql.FieldHasPrefix(FieldStartTime, v))
}

// StartTimeHasSuffix applies the HasSuffix predicate on the "startTime" field.
func StartTimeHasSuffix(v string) predicate.Times {
	return predicate.Times(sql.FieldHasSuffix(FieldStartTime, v))
}

// StartTimeEqualFold applies the EqualFold predicate on the "startTime" field.
func StartTimeEqualFold(v string) predicate.Times {
	return predicate.Times(sql.FieldEqualFold(FieldStartTime, v))
}

// StartTimeContainsFold applies the ContainsFold predicate on the "startTime" field.
func StartTimeContainsFold(v string) predicate.Times {
	return predicate.Times(sql.FieldContainsFold(FieldStartTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Times) predicate.Times {
	return predicate.Times(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Times) predicate.Times {
	return predicate.Times(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Times) predicate.Times {
	return predicate.Times(sql.NotPredicates(p))
}
