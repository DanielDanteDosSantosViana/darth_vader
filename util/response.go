package util

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrorNoRedirecionamento = errors.New("Error ao tentar redirencionar")

func ResponseOK(w http.ResponseWriter, a ...interface{}) {
	if a != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", a)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func ResponseInternalError(w http.ResponseWriter, a ...interface{}) {
	if a != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", a)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ResponseCreate(w http.ResponseWriter, a ...interface{}) {
	if a != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", a)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func ResponseBadRequest(w http.ResponseWriter, a ...interface{}) {
	if a != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", a)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Redirect(urlAcesso string) {

}
