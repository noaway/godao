package godao

import (
	"testing"
	"time"
)

func TestOrmLogger_Print(t *testing.T) {
	file := "/Users/wzy/dev/zaoshu/go/src/github.com/zaoshu/tasuku-service/model/tasuku_task_statistics.go:96"
	duration := 2 * time.Millisecond
	sql := "SELECT \"total_url\", \"done_url\", \"failed_url\", \"record_num\", \"network_traffic\" FROM \"tasuku_task_statistics\"  WHERE (id = $1)"
	vars := []interface{}{1}

	l := ormLogger{}
	l.Print("sql", file, duration, sql, vars)
}
