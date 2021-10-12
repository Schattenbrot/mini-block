package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/posts", app.getAllPosts)
	router.HandlerFunc(http.MethodGet, "/v1/posts/:id", app.getOnePost)
	router.HandlerFunc(http.MethodPost, "/v1/posts", app.addPost)
	router.HandlerFunc(http.MethodDelete, "/v1/posts/:id", app.deletePost)

	return router
}
