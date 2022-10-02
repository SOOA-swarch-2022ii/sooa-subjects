package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	db "github.com/SOOA-swarch-2022ii/sooa-subjects/db"
	"github.com/SOOA-swarch-2022ii/sooa-subjects/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Cliente_mongo = db.Dbconnect()
var Asigs_handler *mongo.Collection = Cliente_mongo.Database("SOOA_subjects_db").Collection("subjects")
var Curso_handler *mongo.Collection = Cliente_mongo.Database("SOOA_subjects_db").Collection("courses")

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

func NewCourse(response http.ResponseWriter, request *http.Request) {

	var curso models.Course
	response.Header().Set("content-type", "application/json")
	err := json.NewDecoder(request.Body).Decode(&curso)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"algo salio mal en el formato json de la asignatura:` + err.Error() + ` "}`))
		return
	}

	result, err := Curso_handler.InsertOne(context.TODO(), curso)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en mongo:` + err.Error() + ` "}`))
		return
	}
	res, _ := json.Marshal(result.InsertedID)

	fmt.Println(`curso insertado en id:` + strings.Replace(string(res), `"`, ``, 2))

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"curso creado"}`))

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
	fmt.Println("buscando código: " + id)
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

/*
func GetsbCAM(response http.ResponseWriter, request *http.Request)           {}
func GetsbCAMFA(response http.ResponseWriter, request *http.Request)         {}
func GetsbCAMFABAU(response http.ResponseWriter, request *http.Request)      {}6327497ce0e25fe0aaf4ebc4
*/
func GetcoID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	var id = params["id"]
	objid, _ := primitive.ObjectIDFromHex(id)
	fmt.Println("buscando id: " + id)
	var curso models.Course
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filtro := bson.D{{"_id", objid}}
	err := Curso_handler.FindOne(ctx, filtro).Decode(&curso)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en Asigsbyid: ` + err.Error() + ` "}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	//de json a struct golang
	json.NewEncoder(response).Encode(curso)
}

func GetcoSB(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	var sb = params["subject"]
	var cursos []models.Course
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filtro := bson.D{{"subject", sb}}
	cursor, err := Curso_handler.Find(ctx, filtro)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en GetcoSB":` + err.Error() + ` "}`))
		return
	}
	if err = cursor.All(ctx, &cursos); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en GetcoSB":` + err.Error() + ` "}`))
		return
	}
	fmt.Println("se encontraron: ", len(cursos), " cursos")
	json.NewEncoder(response).Encode(cursos)
}

func GetcoSBSemester(response http.ResponseWriter, request *http.Request)    {}
func GetcoSchedule(response http.ResponseWriter, request *http.Request)      {}
func GetcoStudeent(response http.ResponseWriter, request *http.Request)      {}
func GetcoProff(response http.ResponseWriter, request *http.Request)         {}
func GetcoProffSemester(response http.ResponseWriter, request *http.Request) {}
func GetcoLocation(response http.ResponseWriter, request *http.Request)      {}
func UpdateSB(response http.ResponseWriter, request *http.Request)           {}
func UpdateCO(response http.ResponseWriter, request *http.Request)           {}
func DeleteSB(response http.ResponseWriter, request *http.Request)           {}
func DeleteCO(response http.ResponseWriter, request *http.Request)           {}
