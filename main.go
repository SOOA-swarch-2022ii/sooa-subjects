package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	cont "github.com/SOOA-swarch-2022ii/sooa-subjects/controllers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Inicializando microservicio")
	Cliente_mongo, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://sooa_mongo_admin:CbfRdzY1dULYKIiE@sooa-mongo-cluster.lrlq0px.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Cliente_mongo.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer Cliente_mongo.Disconnect(ctx)

	cont.Asigs_handler = Cliente_mongo.Database("SOOA_subjects_db").Collection("subjects")
	cont.Curso_handler = Cliente_mongo.Database("SOOA_subjects_db").Collection("courses")

	enrutador := mux.NewRouter()
	enrutador.HandleFunc("/subjects", cont.AsigsResponseTodos).Methods("GET")
	//enrutador.HandleFunc("/subjects/{name}", controllers.AsigsNombre).Methods("GET")
	enrutador.HandleFunc("/courses", cont.CursosResponseTodos).Methods("GET")
	http.ListenAndServe(":6666", enrutador)

}
