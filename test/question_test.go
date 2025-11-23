package test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/go-chi/chi/v5"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "qa-service/internal/handlers"
    "qa-service/internal/models"
)

func setupTestDB(t *testing.T) *gorm.DB {
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        t.Fatalf("failed to init test DB: %v", err)
    }
    if err := db.AutoMigrate(&models.Question{}, &models.Answer{}); err != nil {
        t.Fatalf("failed migration: %v", err)
    }
    return db
}

func setupAnswerRouter(h *handlers.AnswerHandler) *chi.Mux {
    r := chi.NewRouter()
    r.Post("/questions/{id}/answers", h.Create)
    return r
}

func TestAnswerCreate(t *testing.T) {
    db := setupTestDB(t)

    q := models.Question{Text: "Что такое Go?"}
    db.Create(&q)

    handler := &handlers.AnswerHandler{DB: db}
    router := setupAnswerRouter(handler)

    body := []byte(`{"user_id":"550e8400-e29b-41d4-a716-446655440000","text":"Go - это язык от Google."}`)

    req := httptest.NewRequest(
        http.MethodPost,
        "/questions/1/answers",
        bytes.NewReader(body),
    )
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var a models.Answer
    err := json.NewDecoder(w.Body).Decode(&a)
    assert.NoError(t, err)

    assert.NotZero(t, a.ID)
    assert.Equal(t, q.ID, a.QuestionID)
    assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", a.UserID)
    assert.Equal(t, "Go - это язык от Google.", a.Text)

    var dbA models.Answer
    err = db.First(&dbA, a.ID).Error
    assert.NoError(t, err)

    assert.Equal(t, a.UserID, dbA.UserID)
    assert.Equal(t, a.Text, dbA.Text)
    assert.Equal(t, q.ID, dbA.QuestionID)
}
