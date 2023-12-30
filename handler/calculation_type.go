package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type CalculationType struct {
	ID          int32            `json:"id"`
	Description string           `json:"description" validate:"required"`
	IsDeleted   pgtype.Bool      `json:"is_deleted"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreateCalculationType(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	calculationtype := CalculationType{}
	err := json.NewDecoder(r.Body).Decode(&calculationtype)

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
	err = validate.Struct(calculationtype)
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

	arg := db.CreateCalculationTypeParams{
		Description: calculationtype.Description,
		IsDeleted:  calculationtype.IsDeleted,
	}

	calculationTypeInfo, err := server.store.CreateCalculationType(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to create calculation type")
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.CalculationType `json:"data"`
	}{
		Status:  true,
		Message: "Calculation Type created successfully",
		Data:    []db.CalculationType{calculationTypeInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerGetAllCalculationType(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	calculationTypeInfo, err := server.store.GetAllCalculationType(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch calculation type",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.CalculationType `json:"data"` // Use []db.BrandsLanguage
	}{
		Status:  true,
		Message: "Calculation Type retrieved successfully",
		Data:    calculationTypeInfo,
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