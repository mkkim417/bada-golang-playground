package main

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/status-418", func() (int, string) {
		return 418, "i'm a teapot" // HTTP 418 : "i'm a teapot"
	})

	m.Get("/res-req", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
		res.WriteHeader(200) // HTTP 200
	})

	//m.Get("/status-test", func() {
	//m.Patch("/status-test", func() {
	//m.Post("/status-test", func() {
	//m.Put("/status-test", func() {
	//m.Delete("/status-test", func() {
	//m.Options("/status-test", func() {

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	/*
		m.Group("/books", func(r martini.Router) {
			r.Get("/:id", GetBooks)
			r.Post("/new", NewBook)
			r.Put("/update/:id", UpdateBook)
			r.Delete("/delete/:id", DeleteBook)
		})
	*/
	/*
		static files
	*/
	m.Use(martini.Static("assets"))
	/*
		not found handler with static files
	*/
	static := martini.Static("assets", martini.StaticOptions{Fallback: "/index.html", Exclude: "/api/v"})
	m.NotFound(static, http.NotFound)

	/*
		middleware
	*/
	/* api key check */
	/*
		m.Use(func(res http.ResponseWriter, req *http.Request) {
			// validate an api key
			if req.Header.Get("X-API-KEY") != "secret123" {
				res.WriteHeader(http.StatusUnauthorized)
			}
		})
	*/

	m.Use(func(c martini.Context, log *log.Logger) {
		log.Println("before a request")

		c.Next()

		log.Println("after a request")
	})

	m.Run()
}
