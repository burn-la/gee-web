package main

import (
	"gee"
	"net/http"
	"time"
)

//func indexHandler(w http.ResponseWriter, req *http.Request){
//	fmt.Fprintf(w, "url = %q",req.URL.Path)
//}
//
//func helloHandler(w http.ResponseWriter, req *http.Request){
//	for k,v := range req.Header{
//		fmt.Fprintf(w, "header[%q] = %q",k,v)
//	}
//}

//type Engine struct {
//
//}
//
//func (e Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
//	switch req.URL.Path {
//	case "/":
//		fmt.Fprintf(w, "url = %q",req.URL.Path)
//	case "/hello":
//		for k,v := range req.Header{
//					fmt.Fprintf(w, "header[%q] = %q",k,v)
//				}
//	default:
//		fmt.Fprintf(w,"404 not found %q", req.URL)
//
//	}
//}


func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.String(http.StatusOK, "this is second %s\n", c.Path)
		c.Next()
		// Calculate resolution time
		c.String(c.StatusCode,"%s in %v for group v2\n ",  c.Req.RequestURI, time.Since(t))
	}
}

func onlyForV3() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.String(http.StatusOK, "this is third %s\n", c.Path)
		c.Next()
		// Calculate resolution time
		c.String(c.StatusCode,"%s in %v for group v3\n ",  c.Req.RequestURI, time.Since(t))
	}
}

func main(){
	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/hello", helloHandler)
	//log.Fatal(http.ListenAndServe(":9999",nil))

	//engine := new(Engine)
	//log.Fatal(http.ListenAndServe(":9999",engine))

	//engine := gee.New()
	//engine.Get("/", func (c *gee.Context){
	//		c.String(http.StatusOK, "string %s", c.Req.URL)
	//	})
	//engine.Post("/hello", func (c *gee.Context){
	//		c.Json(http.StatusOK, gee.H{
	//			"name" : c.PostForm("name"),
	//			"age" : c.PostForm("age"),
	//		})
	//	})
	//engine.Get("/html", func(c *gee.Context) {
	//	c.Html(http.StatusOK, "<h1>Hello Gee</h1>")
	//})
	//
	//engine.Get("/hello/:name", func(c *gee.Context) {
	//	// expect /hello/geektutu
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})
	//
	//engine.Get("/assets/*filepath", func(c *gee.Context) {
	//	c.Json(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	//})
	//
	//engine.Run(":9999",engine)

	//r := gee.New()
	//r.User(gee.Logger())
	//r.Get("/index", func(c *gee.Context) {
	//	c.Html(http.StatusOK, "<h1>Index Page</h1>")
	//})
	//v1 := r.Group("/v1")
	//{
	//	v1.Get("/", func(c *gee.Context) {
	//		c.Html(http.StatusOK, "<h1>Hello Gee</h1>")
	//	})
	//
	//	v1.Get("/hello", func(c *gee.Context) {
	//		// expect /hello?name=geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//	})
	//}
	//v2 := r.Group("/v2")
	//v2.User(onlyForV2(),onlyForV3())
	//{
	//	v2.Get("/hello/:name", func(c *gee.Context) {
	//		// expect /hello/geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//	})
	//	v2.Post("/login", func(c *gee.Context) {
	//		c.Json(http.StatusOK, gee.H{
	//			"username": c.PostForm("username"),
	//			"password": c.PostForm("password"),
	//		})
	//	})
	//v3 := r.Group("/v3")
	//v3.User(onlyForV3())
	//v3.Get("/hhh", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hhhhhh")
	//})
	//
	//}
	//
	//r.Run(":9999")

	r := gee.New()
	r.User(gee.Logger(), gee.Recover())
	r.Get("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.Get("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}




