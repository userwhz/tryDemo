package main

import (
	"net/http"
	"tryDemo/project/seven/gee/day7/gee"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello www\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"www"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
