package builder

import (
	"fmt"
	"strings"
)

func Exec(action build) {
	builder := SqlBuilder{&actions{}}
	action(&builder)
	sql := builder.build()
	fmt.Println(strings.Join(sql.query, " "))
}
