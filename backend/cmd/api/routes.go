package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes handles all the URL-routing.
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/posts", app.findAllPosts)
	router.HandlerFunc(http.MethodGet, "/v1/posts/:id", app.findOnePost)
	router.HandlerFunc(http.MethodPost, "/v1/posts", app.createPost)
	router.HandlerFunc(http.MethodDelete, "/v1/posts/:id", app.deleteOnePost)

	return router
}
