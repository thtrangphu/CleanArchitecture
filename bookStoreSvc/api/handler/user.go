package handler

import (
	"encoding/json"
	"log"

	//"github.com/gorilla/mux"
	"net/http"

	"github.com/thtrangphu/bookStoreSvc/usercase/user"

	"github.com/thtrangphu/bookStoreSvc/api/presenter"

	"github.com/thtrangphu/bookStoreSvc/entity"
	//"github.com/urfave/negroni"
	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/mux"
)

//

func listUsers(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error list users"
		var data []*entity.User
		var err error
		name := r.URL.Query().Get("name")
		if name == "" {
			data, err = service.ListUsers()
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.User
		for _, d := range data {
			toJ = append(toJ, &presenter.User{
				ID:       d.ID,
				Email:    d.Email,
				FullName: d.FullName,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createUser(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			FullName string `json:"full_name"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateUser(input.Email, input.Password, input.FullName)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.User{
			ID:       id,
			Email:    input.Email,
			FullName: input.FullName,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//func getUser(service user.UseCase) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		errorMessage := "Error reading user"
//		vars := mux.Vars(r)
//		id, err := entity.StringToID(vars["id"])
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(errorMessage))
//			return
//		}
//		data, err := service.GetUser(id)
//		w.Header().Set("Content-Type", "application/json")
//		if err != nil && err != entity.ErrNotFound {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(errorMessage))
//			return
//		}
//
//		if data == nil {
//			w.WriteHeader(http.StatusNotFound)
//			w.Write([]byte(errorMessage))
//			return
//		}
//		toJ := &presenter.User{
//			ID:        data.ID,
//			Email:     data.Email,
//			FirstName: data.FirstName,
//			LastName:  data.LastName,
//		}
//		if err := json.NewEncoder(w).Encode(toJ); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(errorMessage))
//		}
//	})
//}
//
//func deleteUser(service user.UseCase) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		errorMessage := "Error removing user"
//		vars := mux.Vars(r)
//		id, err := entity.StringToID(vars["id"])
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(errorMessage))
//			return
//		}
//		err = service.DeleteUser(id)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			w.Write([]byte(errorMessage))
//			return
//		}
//	})
//}
//
//// MakeUserHandlers make url handlers
//func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
//	r.Handle("/v1/user", n.With(
//		negroni.Wrap(listUsers(service)),
//	)).Methods("GET", "OPTIONS").Name("listUsers")
//
//	r.Handle("/v1/user", n.With(
//		negroni.Wrap(createUser(service)),
//	)).Methods("POST", "OPTIONS").Name("createUser")
//
//	r.Handle("/v1/user/{id}", n.With(
//		negroni.Wrap(getUser(service)),
//	)).Methods("GET", "OPTIONS").Name("getUser")
//
//	r.Handle("/v1/user/{id}", n.With(
//		negroni.Wrap(deleteUser(service)),
//	)).Methods("DELETE", "OPTIONS").Name("deleteUser")
//}
