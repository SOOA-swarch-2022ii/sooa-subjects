package controllers

import (
	"context"
	"encoding/json"

	//"fmt"
	"net/http"
	"time"

	"github.com/SOOA-swarch-2022ii/sooa-subjects/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Asigs_handler *mongo.Collection
var Curso_handler *mongo.Collection
var Cliente_mongo *mongo.Client

func AsigsResponseTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var asignaturas []models.Subject
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := Asigs_handler.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en Asigs_handler.find":` + err.Error() + ` "}`))
		return
	}
	if err = cursor.All(ctx, &asignaturas); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en cursor.all":` + err.Error() + ` "}`))
		return
	}
	json.NewEncoder(response).Encode(asignaturas)
}

func CursosResponseTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var cursos []models.Course
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := Curso_handler.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en Curso_handler.Find":` + err.Error() + ` "}`))
		return
	}
	if err = cursor.All(ctx, &cursos); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en cursor.all":` + err.Error() + ` "}`))
		return
	}
	json.NewEncoder(response).Encode(cursos)
}
