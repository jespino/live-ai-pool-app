package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jespino/pool-app/internal/database"
	"github.com/jespino/pool-app/internal/models"
)

type Handler struct {
	DB database.Storage
}

func NewHandler(db database.Storage) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) CreatePool(w http.ResponseWriter, r *http.Request) {
	var poolRequest struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Options     []string `json:"options"`
		ExpiresIn   int      `json:"expires_in"` // in hours
	}

	if err := json.NewDecoder(r.Body).Decode(&poolRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create pool options with unique IDs
	options := make([]models.Option, 0, len(poolRequest.Options))
	for _, optText := range poolRequest.Options {
		options = append(options, models.Option{
			ID:    uuid.New().String(),
			Text:  optText,
			Votes: 0,
		})
	}

	// Create a new pool
	pool := models.Pool{
		ID:          uuid.New().String(),
		Title:       poolRequest.Title,
		Description: poolRequest.Description,
		Options:     options,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(time.Duration(poolRequest.ExpiresIn) * time.Hour),
	}

	if err := h.DB.CreatePool(pool); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pool)
}

func (h *Handler) GetPool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolID := vars["id"]

	pool, err := h.DB.GetPool(poolID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pool)
}

func (h *Handler) ListPools(w http.ResponseWriter, r *http.Request) {
	pools := h.DB.ListPools()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pools)
}

func (h *Handler) Vote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolID := vars["id"]

	var voteRequest struct {
		UserID   string `json:"user_id"`
		OptionID string `json:"option_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&voteRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vote := models.Vote{
		ID:        uuid.New().String(),
		PoolID:    poolID,
		OptionID:  voteRequest.OptionID,
		UserID:    voteRequest.UserID,
		CreatedAt: time.Now(),
	}

	if err := h.DB.CreateVote(vote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the updated pool to return
	pool, err := h.DB.GetPool(poolID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pool)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest struct {
		Username string `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{
		ID:       uuid.New().String(),
		Username: userRequest.Username,
	}

	if err := h.DB.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}