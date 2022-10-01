package main

import (
	"fmt"
	"github.com/SOOA-swarch-2022ii/sooa-subjects/routes"
	"net/http"
)

func main() {
	fmt.Println("Inicializando microservicio sooa_subjects_ms")
	//enrutador := mux.NewRouter()
	enrutador := routes.Routes()
	http.ListenAndServe(":6666", enrutador)
}
