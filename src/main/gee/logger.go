package gee

import (
	"time"
)

func Logger() HandlerFunc{
	return func(c *Context) {
		t := time.Now()
		c.Next()
		c.String(c.StatusCode,"%s in %v for logger\n ",  c.Req.RequestURI, time.Since(t))
	}
}
