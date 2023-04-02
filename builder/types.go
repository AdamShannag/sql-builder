package builder

type sql struct{ query []string }
type sqlMods func(*sql)
type actions struct{ mods []sqlMods }
type build func(*SqlBuilder)
type SqlBuilder struct{ acts *actions }
type SqlConditionBuilder struct{ SqlBuilder }
