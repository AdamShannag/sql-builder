package builder

func (b *SqlBuilder) Where() *SqlConditionBuilder {
	b.acts.mods = append(b.acts.mods, func(s *sql) {
		s.query = append(s.query, where)
	})
	return &SqlConditionBuilder{*b}
}

func (b *SqlConditionBuilder) Is(left, right string, is int) *SqlConditionBuilder {
	b.acts.mods = append(b.acts.mods, func(s *sql) {
		s.query = append(s.query, left, operationAsString(is), right)
	})
	return b
}

func (b *SqlConditionBuilder) And() *SqlConditionBuilder {
	b.acts.mods = append(b.acts.mods, func(s *sql) {
		s.query = append(s.query, and)
	})
	return b
}

func (b *SqlConditionBuilder) Or() *SqlConditionBuilder {
	b.acts.mods = append(b.acts.mods, func(s *sql) {
		s.query = append(s.query, or)
	})
	return b
}
