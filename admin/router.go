package admin

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"os"
)

func RegRouter(prefix string, Router *mux.Router) {
	s := Router.PathPrefix(prefix).Subrouter()
	s.PathPrefix("/").Handler(http.StripPrefix(prefix,http.FileServer(http.Dir("./src/github.com/kittuov/go-django/admin/webapp/static"))))
	s.HandleFunc("/shit/", testing)
}

func testing (w http.ResponseWriter, r *http.Request){
	path,_ := os.Getwd()

	fmt.Fprint(w, path )
	fmt.Fprintln(w, "hello world")
}
