package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"
    "qa-service/internal/models"
)

type AnswerHandler struct {
    DB *gorm.DB
}

func (h *AnswerHandler) Create(w http.ResponseWriter, r *http.Request) {
    qid := chi.URLParam(r, "id")

    var q models.Question
    if err := h.DB.First(&q, qid).Error; err != nil {
        http.Error(w, "question not found", http.StatusBadRequest)
        return
    }

    var a models.Answer
    json.NewDecoder(r.Body).Decode(&a)
    a.QuestionID = q.ID

    h.DB.Create(&a)
    json.NewEncoder(w).Encode(a)
}

func (h *AnswerHandler) GetOne(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var a models.Answer

    if err := h.DB.First(&a, id).Error; err != nil {
        http.NotFound(w, r)
        return
    }

    json.NewEncoder(w).Encode(a)
}

func (h *AnswerHandler) Delete(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    h.DB.Delete(&models.Answer{}, id)
    w.WriteHeader(http.StatusNoContent)
}
