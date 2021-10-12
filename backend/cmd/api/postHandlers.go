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

func (app *application) addPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	type jsonResp struct {
		OK bool   `json:"ok"`
		ID string `json:"_id"`
	}

	id, err := app.models.DB.PostPost(post)
	if err != nil {
		app.logger.Println("fuck this failed already")
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

func (app *application) getOnePost(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := primitive.ObjectIDFromHex(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	post, err := app.models.DB.GetOnePost(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, post, "post")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := app.models.DB.GetAllPosts()
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, posts, "posts")
	if err != nil {
		app.logger.Println(err)
	}

}

func (app *application) deletePost(w http.ResponseWriter, r *http.Request) {
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
