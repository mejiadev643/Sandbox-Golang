package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mejiadev643/excelgo/DB"
	jwtpersonal "github.com/mejiadev643/excelgo/jwt"
	"github.com/mejiadev643/excelgo/middleware"
	"github.com/mejiadev643/excelgo/routes"
)

func main() {
	DB.Connect()         //conexion a la base de datos
	r := mux.NewRouter() // enrutador

	//metodos de rutas
	r.HandleFunc("/excel", routes.SaveExcel).Methods("GET")
	r.Handle("/units", middleware.ControlMiddleware(http.HandlerFunc(routes.GetUnidadesHandler))).Methods("GET")
	r.HandleFunc("/token", jwtpersonal.CreateToken).Methods("GET")

	http.ListenAndServe(":8080", r) //escucha el puerto 8080
}
