// Code generated by ent, DO NOT EDIT.

package reserve

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the reserve type in the database.
	Label = "reserve"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldOrderedID holds the string denoting the ordered_id field in the database.
	FieldOrderedID = "ordered_id"
	// FieldDay holds the string denoting the day field in the database.
	FieldDay = "day"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// Table holds the table name of the reserve in the database.
	Table = "reserves"
)

// Columns holds all SQL columns for reserve fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
	FieldStartTime,
	FieldOrderedID,
	FieldDay,
	FieldDetail,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Reserve queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByOrderedID orders the results by the ordered_id field.
func ByOrderedID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderedID, opts...).ToFunc()
}

// ByDay orders the results by the day field.
func ByDay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDay, opts...).ToFunc()
}

// ByDetail orders the results by the detail field.
func ByDetail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDetail, opts...).ToFunc()
}
