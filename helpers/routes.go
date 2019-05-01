package helpers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type RestReader interface {
	Read(w http.ResponseWriter, r *http.Request)
}
type RestManyReader interface {
	ReadMany(w http.ResponseWriter, r *http.Request)
}
type RestCreator interface {
	Create(w http.ResponseWriter, r *http.Request)
}
type RestUpdater interface {
	Update(w http.ResponseWriter, r *http.Request)
}
type RestReplacer interface {
	Replace(w http.ResponseWriter, r *http.Request)
}
type RestDeleter interface {
	Delete(w http.ResponseWriter, r *http.Request)
}

func Route(path string, controller interface{}, router *mux.Router) {
	if c, ok := controller.(RestReader); ok {
		router.HandleFunc(path+"/{id}", c.Read).Methods("GET")
	}
	if c, ok := controller.(RestManyReader); ok {
		router.HandleFunc(path, c.ReadMany).Methods("GET")
	}
	if c, ok := controller.(RestCreator); ok {
		router.HandleFunc(path, c.Create).Methods("POST")
	}

	if c, ok := controller.(RestDeleter); ok {
		router.HandleFunc(path+"/{id}", c.Delete).Methods("DELETE")
	}
	if c, ok := controller.(RestReplacer); ok {
		router.HandleFunc(path+"/{id}", c.Replace).Methods("PUT")
	}
	if c, ok := controller.(RestUpdater); ok {
		router.HandleFunc(path+"/{id}", c.Update).Methods("PATH")
	}
	router.Handle(path, http.NotFoundHandler()).Methods("*")
}
