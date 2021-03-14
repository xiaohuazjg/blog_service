package api

import (
	"expvar"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Expvar(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Printf(c.Writer, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Printf(c.Writer, "%q:%q", key, str)

		} else {
			fmt.Printf(c.Writer, "%q:%v", key, value)
		}

	}
	fmt.Printf(c.Writer, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Printf(c.Writer, "\n}\n")

}
