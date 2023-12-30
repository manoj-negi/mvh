package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type Park struct {
	ID        int32            `json:"id"`
	BrandID   int32            `json:"brand_id" validate:"required"`
	Name      pgtype.Text      `json:"name" validate:"required"`
	CountryID pgtype.Int4      `json:"country_id" validate:"required"`
	Validated pgtype.Bool      `json:"validated"`
	IsDeleted pgtype.Bool      `json:"is_deleted"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreatePark(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	park := Park{}
	err := json.NewDecoder(r.Body).Decode(&park)

	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusNotAcceptable,
		}

		util.WriteJSONResponse(w, http.StatusNotAcceptable, jsonResponse)
		return
	}

	validate := validator.New()
	err = validate.Struct(park)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err != nil {
				jsonResponse := JsonResponse{
					Status:     false,
					Message:    "Invalid value for " + err.Field(),
					StatusCode: http.StatusNotAcceptable,
				}

				json.NewEncoder(w).Encode(jsonResponse)
				return

			}
		}
	}

	arg := db.CreateParkParams{
		BrandID:   park.BrandID,
		Name:      park.Name.String,
		CountryID: park.CountryID,
		Validated: park.Validated,
		IsDeleted: park.IsDeleted,
	}

	parkInfo, err := server.store.CreatePark(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to create park")
		return
	}

	response := struct {
		Status  bool      `json:"status"`
		Message string    `json:"message"`
		Data    []db.Park `json:"data"`
	}{
		Status:  true,
		Message: "Park created successfully",
		Data:    []db.Park{parkInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerGetAllPark(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	parkInfo, err := server.store.GetAllPark(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch park",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool      `json:"status"`
		Message string    `json:"message"`
		Data    []db.Park `json:"data"`
	}{
		Status:  true,
		Message: "park retrieved successfully",
		Data:    parkInfo,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to encode response",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}
}

func (server *Server) handlerGetParksByBrand(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	vars := mux.Vars(r)
	idParam, ok := vars["id"]
	if !ok {
		util.ErrorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		util.ErrorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
		return
	}

	parks, err := server.store.GetParksByBrand(ctx, int32(id))
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch parks",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool      `json:"status"`
		Message string    `json:"message"`
		Data    []db.Park `json:"data"`
	}{
		Status:  true,
		Message: "Parks retrieved successfully",
		Data:    parks,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to encode response",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}
}
