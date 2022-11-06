package routes

import (
	control "github.com/SOOA-swarch-2022ii/sooa-subjects/controllers"
	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/new-subject", control.NewSubject).Methods("POST")
	router.HandleFunc("/new-course", control.NewCourse).Methods("POST")

	router.HandleFunc("/subjects/name={name}", control.GetsbName).Methods("GET")
	router.HandleFunc("/subjects/id={id}", control.GetsbID).Methods("GET")
	router.HandleFunc("/subjects/code={code}", control.GetsbCode).Methods("GET")
	router.HandleFunc("/subjects/cam={campus}/fac={faculty}", control.GetsbCAMFA).Methods("GET")
	/*
		router.HandleFunc("/subjects/{campus}", control.GetsbCAM).Methods("GET")
		router.HandleFunc("/subjects/{campus}/{faculty}/{bau}", control.GetsbCAMFABAU).Methods("GET")
	*/
	router.HandleFunc("/courses/all", control.CoTodos).Methods("GET")
	router.HandleFunc("/subjects/all", control.SbTodos).Methods("GET")
	router.HandleFunc("/courses/id={id}", control.GetcoID).Methods("GET")
	router.HandleFunc("/courses/sb={subject}", control.GetcoSB).Methods("GET")
	router.HandleFunc("/courses/sb={subject}/sm={semester}", control.GetcoSBSemester).Methods("GET")
	router.HandleFunc("/courses/sb={subject}/sm={semester}/sch/d={day}", control.GetcoSbSmSchDay).Methods("GET")

	router.HandleFunc("/courses/semester={sm}/d={day}/ti={ti}/tf={tf}", control.GetcoSchedule).Methods("GET")

	router.HandleFunc("/courses/st={student}", control.GetcoStudent).Methods("GET")
	router.HandleFunc("/courses/st={student}/sm={semester}", control.GetcoStSm).Methods("GET")
	router.HandleFunc("/courses/profe={professor}", control.GetcoProfe).Methods("GET")
	router.HandleFunc("/courses/location/house={house}", control.GetcoLocation).Methods("GET")
	router.HandleFunc("/courses/profe={professor}/sm={semester}", control.GetcoProffSemester).Methods("GET")
	/*

		router.HandleFunc("/subject/{id}", control.UpdateSB).Methods("PUT")
		router.HandleFunc("/course/{id}", control.UpdateCO).Methods("PUT")

		router.HandleFunc("/subject/{id}", control.DeleteSB).Methods("DELETE")
		router.HandleFunc("/course/{id}", control.DeleteCO).Methods("DELETE")*/

	return router
}
