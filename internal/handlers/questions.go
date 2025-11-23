package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"

    "qa-service/internal/models"
)

type QuestionHandler struct {
    DB *gorm.DB
}

func (h *QuestionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    var qs []models.Question
    h.DB.Find(&qs)
    json.NewEncoder(w).Encode(qs)
}

func (h *QuestionHandler) Create(w http.ResponseWriter, r *http.Request) {
    var q models.Question
    json.NewDecoder(r.Body).Decode(&q)
    h.DB.Create(&q)
    json.NewEncoder(w).Encode(q)
}

func (h *QuestionHandler) GetOne(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var q models.Question

    if err := h.DB.Preload("Answers").First(&q, id).Error; err != nil {
        http.NotFound(w, r)
        return
    }

    json.NewEncoder(w).Encode(q)
}

func (h *QuestionHandler) Delete(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    h.DB.Delete(&models.Question{}, id)
    w.WriteHeader(http.StatusNoContent)
}
