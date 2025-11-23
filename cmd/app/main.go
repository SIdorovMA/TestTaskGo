package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "qa-service/internal/db"
    "qa-service/internal/handlers"
)

func main() {
    database, err := db.Connect()
    if err != nil {
        log.Fatal(err)
    }

    qh := &handlers.QuestionHandler{DB: database}
    ah := &handlers.AnswerHandler{DB: database}

    r := chi.NewRouter()

    r.Get("/questions/", qh.GetAll)
    r.Post("/questions/", qh.Create)
    r.Get("/questions/{id}", qh.GetOne)
    r.Delete("/questions/{id}", qh.Delete)

    r.Post("/questions/{id}/answers/", ah.Create)
    r.Get("/answers/{id}", ah.GetOne)
    r.Delete("/answers/{id}", ah.Delete)

    log.Println("Server running at :8080")
    http.ListenAndServe(":8080", r)
}
