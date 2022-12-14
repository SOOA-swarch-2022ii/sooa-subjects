package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SOOA-swarch-2022ii/sooa-subjects/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SbTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var materias []models.Subject
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := Asigs_handler.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en materias.Find":` + err.Error() + ` "}`))
		return
	}
	if err = cursor.All(ctx, &materias); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en materias.all":` + err.Error() + ` "}`))
		return
	}
	json.NewEncoder(response).Encode(materias)
}

func NewSubject(response http.ResponseWriter, request *http.Request) {
	var materia models.Subject
	response.Header().Set("content-type", "application/json")
	err := json.NewDecoder(request.Body).Decode(&materia)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"algo salio mal en el formato json de la asignatura:` + err.Error() + ` "}`))
		return
	}

	result, err := Asigs_handler.InsertOne(context.TODO(), materia)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en la base de datos mongo:` + err.Error() + ` "}`))
		return
	}
	res, _ := json.Marshal(result.InsertedID)

	fmt.Println(`materia inertada en:` + strings.Replace(string(res), `"`, ``, 2))

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"asignatura creada"}`))

}

func GetsbName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	//id, _ := primitive.ObjectIDFromHex(params["id"])
	var nombre string = params["name"]
	fmt.Println("buscando nombre " + nombre)
	var asignatura models.Subject
	filtro := bson.D{{"name", nombre}}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := Asigs_handler.FindOne(ctx, filtro).Decode(&asignatura)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en Asigsbyid: ` + err.Error() + ` "}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	//de json a struct golang
	json.NewEncoder(response).Encode(asignatura)
}

func GetsbID(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	var id = params["id"]
	objid, _ := primitive.ObjectIDFromHex(id)
	fmt.Println("buscando id: " + id)
	var asignatura models.Subject
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filtro := bson.D{{"_id", objid}}
	err := Asigs_handler.FindOne(ctx, filtro).Decode(&asignatura)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en Asigsbyid: ` + err.Error() + ` "}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	//de json a struct golang
	json.NewEncoder(response).Encode(asignatura)
}

func GetsbCode(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	var id = params["code"]
	fmt.Println("buscando c??digo: " + id)
	var asignatura models.Subject
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filtro := bson.D{{"code", id}}
	err := Asigs_handler.FindOne(ctx, filtro).Decode(&asignatura)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en Asigsbyid: ` + err.Error() + ` "}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	//de json a struct golang
	json.NewEncoder(response).Encode(asignatura)
}

func GetsbCAM(response http.ResponseWriter, request *http.Request)      {}
func GetsbCAMFA(response http.ResponseWriter, request *http.Request)    {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	var campus = params["campus"]
	var facultad = params["faculty"]
	var cursos []models.Subject

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filtro := bson.D{{"campus", campus},{"faculty",facultad}}
	cursor, err := Asigs_handler.Find(ctx, filtro)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en Getcostudent":` + err.Error() + ` "}`))
		return
	}
	if err = cursor.All(ctx, &cursos); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal al leer los datos getcostudent":` + err.Error() + ` "}`))
		return
	}
	fmt.Println("se encontraron: ", len(cursos), " asignaturas ofertadas en ", campus, " de la facultad ",facultad)
	json.NewEncoder(response).Encode(cursos)
}
func GetsbCAMFABAU(response http.ResponseWriter, request *http.Request) {}
func UpdateSB(response http.ResponseWriter, request *http.Request)      {}
func DeleteSB(response http.ResponseWriter, request *http.Request)      {}
