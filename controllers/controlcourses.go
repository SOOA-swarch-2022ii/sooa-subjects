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

func CoTodos(response http.ResponseWriter, request *http.Request) {
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


func GetcoStudent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	var estudiante = params["student"]
	var cursos []models.Course

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filtro := bson.D{{"students_record.student", estudiante}}
	cursor, err := Curso_handler.Find(ctx, filtro)
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
	fmt.Println("se encontraron: ", len(cursos), " cursos para ", estudiante)
	json.NewEncoder(response).Encode(cursos)
}

func GetcoProfe(response http.ResponseWriter, request *http.Request)         {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	var profe = params["professor"]
	var cursos []models.Course

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filtro := bson.D{{"professors", profe}}
	cursor, err := Curso_handler.Find(ctx, filtro)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal en Getcoprofe":` + err.Error() + ` "}`))
		return
	}
	if err = cursor.All(ctx, &cursos); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"algo salio mal al leer los datos getcoprofe":` + err.Error() + ` "}`))
		return
	}
	fmt.Println("se encontraron: ", len(cursos), " cursos para el profe ", profe)
	json.NewEncoder(response).Encode(cursos)
}


func GetcoSBSemester(response http.ResponseWriter, request *http.Request) {}
func GetcoSchedule(response http.ResponseWriter, request *http.Request)   {}

func GetcoProffSemester(response http.ResponseWriter, request *http.Request) {}
func GetcoLocation(response http.ResponseWriter, request *http.Request)      {}

func UpdateCO(response http.ResponseWriter, request *http.Request)           {}

func DeleteCO(response http.ResponseWriter, request *http.Request)           {}

//Y también porfa un servicio que actualice el valor de los cupos 🥺
//func DeleteCO(response http.ResponseWriter, request *http.Request)           {}
