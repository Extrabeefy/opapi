package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/opAPIProgression/app/model"
)

func GetAllBosses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bosses := []model.Progression{}
	db.Find(&bosses)
	respondJSON(w, http.StatusOK, bosses)
}

func CreateBoss(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	boss := model.Progression{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&boss); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&boss).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, boss)
}

func GetBoss(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	boss := getBossOr404(db, name, w, r)
	if boss == nil {
		return
	}
	respondJSON(w, http.StatusOK, boss)
}

func UpdateBoss(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	boss := getBossOr404(db, name, w, r)
	if boss == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&boss); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&boss).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, boss)
}

func DeleteBoss(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	boss := getBossOr404(db, name, w, r)
	if boss == nil {
		return
	}
	if err := db.Delete(&boss).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func KillBoss(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	boss := getBossOr404(db, name, w, r)
	if boss == nil {
		return
	}
	boss.Killed()
	if err := db.Save(&boss).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, boss)
}

func ReviveBoss(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	boss := getBossOr404(db, name, w, r)
	if boss == nil {
		return
	}
	fmt.Println("test")
	boss.Revive()
	if err := db.Save(&boss).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, boss)
}

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getBossOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Progression {
	boss := model.Progression{}
	if err := db.First(&boss, model.Progression{Bossname: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &boss
}
