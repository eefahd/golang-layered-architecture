package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"golang/internal/models"
	"golang/internal/service"
)

// Server handles HTTP requests (Presentation Layer)
type Server struct {
	router  *chi.Mux
	service *service.Service
}

func NewServer(svc *service.Service) *Server {
	s := &Server{
		router:  chi.NewRouter(),
		service: svc,
	}

	// Middleware
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	// Routes
	s.router.Get("/health", s.handleHealth)
	s.router.Get("/contacts", s.handleGetAll)
	s.router.Get("/contacts/{id}", s.handleGetByID)
	s.router.Post("/contacts", s.handleCreate)
	s.router.Put("/contacts/{id}", s.handleUpdate)
	s.router.Delete("/contacts/{id}", s.handleDelete)

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}

func (s *Server) handleGetAll(w http.ResponseWriter, r *http.Request) {
	contacts, err := s.service.ContactService.GetAll(r.Context())
	if err != nil {
		log.Printf("Error: %v", err)
		respondError(w, http.StatusInternalServerError, "Failed to fetch contacts")
		return
	}
	respondJSON(w, http.StatusOK, contacts)
}

func (s *Server) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	contact, err := s.service.ContactService.GetByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Contact not found")
		return
	}
	respondJSON(w, http.StatusOK, contact)
}

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	created, err := s.service.ContactService.Create(r.Context(), contact)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create contact")
		return
	}

	respondJSON(w, http.StatusCreated, created)
}

func (s *Server) handleUpdate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	contact.ID = id

	if err := s.service.ContactService.UpdateAndNotify(r.Context(), contact); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Contact updated"})
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if err := s.service.ContactService.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete contact")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Contact deleted"})
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
