package builder

import (
	"strings"
)

func (b *SqlBuilder) Select(args ...string) *SqlBuilder {
	b.acts.mods = append(b.acts.mods, func(s *sql) {
		s.query = append(s.query, _select, strings.Join(args, ","))
	})
	return b
}

func (b *SqlBuilder) From(table string) *SqlBuilder {
	b.acts.mods = append(b.acts.mods, func(s *sql) {
		s.query = append(s.query, from, table)
	})
	return b
}

func (b *SqlBuilder) build() *sql {
	q := sql{}
	for _, a := range b.acts.mods {
		a(&q)
	}
	return &q
}
