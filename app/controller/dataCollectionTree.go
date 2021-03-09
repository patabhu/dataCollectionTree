package controller

import (
	"context"
	"dataCollectionTree/app"
	"dataCollectionTree/models"
	"encoding/json"
	"io/ioutil"

	// "goTestServer/GoTestServer"
	"net/http"
)

type ToDoController struct {
	repo app.ToDoRepoInterface
}

func setJsonResponse(w http.ResponseWriter, r *http.Request, httpCode int, resp interface{}) {
	w.WriteHeader(httpCode)
	r.Header.Set("Content-Type", "application/json")
	jsonB, _ := json.MarshalIndent(resp, "", " ")
	w.Write(jsonB)
}

func (td *ToDoController) Insert(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			setJsonResponse(w, req, http.StatusInternalServerError, map[string]interface{}{"error": "internal server error"})
		}
	}()
	if req.Method == http.MethodPost {
		var t models.Data
		reqBody, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqBody, &t)
		ctx := req.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		resp, err := td.repo.Insert(ctx, &t)
		if err != nil {
			setJsonResponse(w, req, http.StatusBadRequest, resp)
		} else {
			setJsonResponse(w, req, http.StatusOK, resp)
		}
	} else {
		setJsonResponse(w, req, http.StatusBadRequest, map[string]interface{}{"error": "method not allowed"})
	}
}

func (td *ToDoController) Query(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			setJsonResponse(w, req, http.StatusInternalServerError, map[string]interface{}{"error": "internal server error"})
		}
	}()
	if req.Method == http.MethodPost {
		var t models.Data
		reqBody, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqBody, &t)
		ctx := req.Context()
		if ctx == nil {
			ctx = context.Background()
		}
		resp, err := td.repo.Query(ctx, &t)
		if err != nil {
			setJsonResponse(w, req, http.StatusBadRequest, resp)
		} else {
			setJsonResponse(w, req, http.StatusOK, resp)
		}
	} else {
		setJsonResponse(w, req, http.StatusBadRequest, map[string]interface{}{"error": "method not allowed"})
	}
}

func NewTodo(s *http.ServeMux, u app.ToDoRepoInterface) {
	hand := &ToDoController{
		repo: u,
	}
	s.HandleFunc("/v1/insert", hand.Insert)
	s.HandleFunc("/v1/query", hand.Query)
}
