package scripts

import (
	"fmt"

	"github.com/AliyevH/exportDB/schemas"
)

// GenerateSQL to generate query
func GenerateSQL(jd *schemas.JSONData) string {
	query := fmt.Sprintf(
		"select %s from %s",
		jd.Columns,
		jd.TableName,
	)
	if len(jd.Condition) > 0 {
		query += fmt.Sprintf(" where %s", jd.Condition)
	}

	return query
}
