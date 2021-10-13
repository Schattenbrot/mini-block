package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Schattenbrot/mini-blog/models"
	"github.com/julienschmidt/httprouter"
)

type PostPayload struct {
	ID        string `json:"_id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

// createPost is the handler for the InsertPost method.
func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	var payload PostPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var post models.Post
	post.Title = payload.Title
	post.Text = payload.Text

	type jsonResp struct {
		OK bool   `json:"ok"`
		ID string `json:"_id"`
	}

	id, err := app.models.DB.InsertPost(post)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
		ID: id.Hex(),
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// findOnePost is the handler for the findOnePost method.
func (app *application) findOnePost(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := primitive.ObjectIDFromHex(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	post, err := app.models.DB.FindOnePost(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, post, "post")
	if err != nil {
		app.logger.Println(err)
	}
}

// findAllPosts is the handler for the FindAllPosts method.
func (app *application) findAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := app.models.DB.FindAllPosts()
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, posts, "posts")
	if err != nil {
		app.logger.Println(err)
	}
}

// deleteOnePost is the handler for the DeleteOnePost method.
func (app *application) deleteOnePost(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := primitive.ObjectIDFromHex(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	post, err := app.models.DB.DeleteOnePost(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, post, "deletedCount")
	if err != nil {
		app.logger.Println(err)
	}
}
