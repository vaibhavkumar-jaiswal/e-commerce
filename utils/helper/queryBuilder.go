package helper

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

// BuildQuery builds a GORM query based on the provided filter struct.
// It uses struct tags to determine the column names and query operators.
func BuildQuery(db *gorm.DB, filter any) *gorm.DB {
	v := reflect.ValueOf(filter)
	t := reflect.TypeOf(filter)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		valueField := v.Field(i)

		if !valueField.CanInterface() {
			continue
		}

		var value interface{}
		if valueField.Kind() == reflect.Ptr {
			if valueField.IsNil() {
				continue
			}
			value = valueField.Elem().Interface()
		} else {
			value = valueField.Interface()
		}

		column := field.Tag.Get("form")
		op := strings.ToUpper(field.Tag.Get("query"))

		if column == "" || isZero(value) {
			continue
		}

		switch op {
		case "LIKE":
			db = db.Where(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%%%v%%", value))
		case "ILIKE":
			db = db.Where(fmt.Sprintf("%s ILIKE ?", column), fmt.Sprintf("%%%v%%", value))
		case "GT":
			db = db.Where(fmt.Sprintf("%s > ?", column), value)
		case "GTE":
			db = db.Where(fmt.Sprintf("%s >= ?", column), value)
		case "LT":
			db = db.Where(fmt.Sprintf("%s < ?", column), value)
		case "LTE":
			db = db.Where(fmt.Sprintf("%s <= ?", column), value)
		case "IN":
			if str, ok := value.(string); ok {
				slice := strings.Split(str, ",")
				db = db.Where(fmt.Sprintf("%s IN ?", column), slice)
			}
		case "BETWEEN":
			if slice, ok := value.([]interface{}); ok && len(slice) == 2 {
				db = db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", column), slice[0], slice[1])
			}
		default:
			db = db.Where(fmt.Sprintf("%s = ?", column), value)
		}
	}

	return db
}

// isZero checks if the value is considered "zero" for its type.
func isZero(v any) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Slice, reflect.Map, reflect.Array:
		return rv.Len() == 0
	case reflect.Bool:
		return false // always include bool
	default:
		zero := reflect.Zero(rv.Type()).Interface()
		return reflect.DeepEqual(v, zero)
	}
}
