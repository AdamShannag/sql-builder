package main

import (
	b "sql-builder/builder"
)

func main() {
	b.Exec(func(sb *b.SqlBuilder) {
		sb.Select("name", "age").
			From("users").
			Where().
			Is("name", "adam", b.EQUAL_TO).
			And().
			Is("age", "15", b.GREATER_THAN_OR_EQUAL)
	})

	// output SELECT name,age FROM users WHERE name = adam AND age >= 15
}
