package admin

import (
	"github.com/gorilla/mux"
	"net/http"
)

/*
Use this method to register router.

*/
func RegRouter(prefix string, Router *mux.Router) {
	s := Router.PathPrefix(prefix).Subrouter()
	s.PathPrefix("/").Handler(http.StripPrefix(prefix, http.FileServer(http.Dir("./src/github.com/kittuov/go-django/admin/webapp/static"))))
}
