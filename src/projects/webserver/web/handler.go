package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rexsimiloluwah/hello-golang/src/projects/webserver/entity"
)

func NewHandler(store entity.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	// sub routers
	h.Use(middleware.Logger)
	h.Get("/tweet", h.NewTweet())
	h.Post("/tweet", h.HandleCreateTweet())
	h.Post("/tweet/delete/{id}", h.HandleDeleteTweet())
	h.Get("/tweets", h.Tweets())

	return h
}

type Handler struct {
	*chi.Mux
	store entity.Store
}

// const tweetsListHTML = `

// `

func (h *Handler) Tweets() http.HandlerFunc {
	// The data to be sent to the HTML template
	type data struct {
		Tweets []entity.Tweet
	}

	tmpl, err := template.ParseFiles("./web/ui/html/tweets.html")
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tweets, err := h.store.GetAllTweets()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data{Tweets: tweets})
	}
}

func (h *Handler) NewTweet() http.HandlerFunc {
	tmpl, err := template.ParseFiles("./web/ui/html/create.html")
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) HandleCreateTweet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := r.FormValue("content")

		if err := h.store.CreateTweet(&entity.Tweet{
			Content: content,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/tweets", http.StatusFound)
	}
}

func (h *Handler) HandleDeleteTweet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.store.DeleteTweet(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/tweets", http.StatusFound)
	}
}
