package video

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	"os"

	"github.com/joho/godotenv"
)

const AllowedCORSDomain = "http://localhost"

func Hola() {
	fmt.Println("test mod")
}

func ConectarDB(parConnectionString string) (*sql.DB, error) {
	return sql.Open("mysql", parConnectionString)
}

//===================================================================================================
// Funciones de CORS

func EnableCORS(parRouter *mux.Router) {

	parRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)

	parRouter.Use(MiddlewareCors)

}

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})

}

//===================================================================================================
// Funciones de respuesta

func RespondWithError(parError error, parWriter http.ResponseWriter) {

	parWriter.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(parWriter).Encode(parError.Error())

}

func RespondWithSuccess(parDato interface{}, parWriter http.ResponseWriter) {

	parWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(parWriter).Encode(parDato)

}

//===================================================================================================

func DefinirCadenaConexion(parCadenaConexion string) {

	var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"
	var (
		parCadenaConexion2 = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			os.Getenv("user"),
			os.Getenv("pass"),
			os.Getenv("host"),
			os.Getenv("port"),
			os.Getenv("db_name"))
	)
	parCadenaConexion = parCadenaConexion2
}
