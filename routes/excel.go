package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mejiadev643/excelgo/DB"
	"github.com/mejiadev643/excelgo/models"
	"gorm.io/gorm/clause"
)

func SaveExcel(w http.ResponseWriter, r *http.Request) {
	var users []models.Usuario
	DB.DB.Find(&users)
	fmt.Println(&users)
	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(w).Encode(&users)
}

func GetUnidadesHandler(w http.ResponseWriter, r *http.Request) {
	var unidades []models.UnidadOrganizativa
	DB.DB.Where("nodo_padre=?", true).Where("activo=?", true).Preload(clause.Associations).Find(&unidades)
	if len(unidades) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Units not found"))
		return
	}

	json.NewEncoder(w).Encode(&unidades)
}
