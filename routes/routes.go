package routes

import (
	control "github.com/SOOA-swarch-2022ii/sooa-subjects/controllers"
	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/sooa-sb-ms/new-subject", control.NewSubject).Methods("POST")
	router.HandleFunc("/sooa-sb-ms/new-course", control.NewCourse).Methods("POST")

	router.HandleFunc("/sooa-sb-ms/subjects/name={name}", control.GetsbName).Methods("GET")
	router.HandleFunc("/sooa-sb-ms/subjects/id={id}", control.GetsbID).Methods("GET")
	router.HandleFunc("/sooa-sb-ms/subjects/code={code}", control.GetsbCode).Methods("GET")
	/*
		router.HandleFunc("/sooa-sb-ms/subjects/{campus}", control.GetsbCAM).Methods("GET")
		router.HandleFunc("/sooa-sb-ms/subjects/{campus}/{faculty}", control.GetsbCAMFA).Methods("GET")
		router.HandleFunc("/sooa-sb-ms/subjects/{campus}/{faculty}/{bau}", control.GetsbCAMFABAU).Methods("GET")
	*/
	router.HandleFunc("/sooa-sb-ms/courses/all", control.CursosResponseTodos).Methods("GET")
	router.HandleFunc("/sooa-sb-ms/courses/id={id}", control.GetcoID).Methods("GET")
	router.HandleFunc("/sooa-sb-ms/courses/sb={subject}", control.GetcoSB).Methods("GET")
	/*
		router.HandleFunc("/sooa-sb-ms/courses/sb={subject}/sm={semester}", control.GetcoSBSemester).Methods("GET")
		router.HandleFunc("/sooa-sb-ms/courses/d={day}/ti={ti}/tf={tf}", control.GetcoSchedule).Methods("GET")

		router.HandleFunc("/sooa-sb-ms/courses/{student}", control.GetcoStudeent).Methods("GET")
		router.HandleFunc("/sooa-sb-ms/courses/{professor}", control.GetcoProff).Methods("GET")
		router.HandleFunc("/sooa-sb-ms/courses/{professor}/{semester}", control.GetcoProffSemester).Methods("GET")

		router.HandleFunc("/sooa-sb-ms/courses/{location}", control.GetcoLocation).Methods("GET")

		router.HandleFunc("/sooa-sb-ms/subject/{id}", control.UpdateSB).Methods("PUT")
		router.HandleFunc("/sooa-sb-ms/course/{id}", control.UpdateCO).Methods("PUT")

		router.HandleFunc("/sooa-sb-ms/subject/{id}", control.DeleteSB).Methods("DELETE")
		router.HandleFunc("/sooa-sb-ms/course/{id}", control.DeleteCO).Methods("DELETE")*/

	return router
}
