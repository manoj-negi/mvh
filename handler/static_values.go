package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)


type StaticValue struct {
	ID           int32            `json:"id"`
	CostInputID  int32            `json:"cost_input_id" validate:"required"`
	Description  string           `json:"description" validate:"required"`
	Value        float64          `json:"value" validate:"required"`
	IsPercentage pgtype.Bool      `json:"is_percentage"`
	IsDeleted    pgtype.Bool      `json:"is_deleted"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}


func (server *Server) handlerCreateStaticValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	staticvalue := StaticValue{}
	err := json.NewDecoder(r.Body).Decode(&staticvalue)

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
	err = validate.Struct(staticvalue)
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

	arg := db.CreateStaticValuesParams{
		CostInputID:  staticvalue.CostInputID,
		Description:  staticvalue.Description,
		Value:        staticvalue.Value,
		IsDeleted:  staticvalue.IsDeleted,
	}

	staticvalueInfo, err := server.store.CreateStaticValues(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to create static value")
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.StaticValue `json:"data"`
	}{
		Status:  true,
		Message: "Static Value created successfully",
		Data:    []db.StaticValue{staticvalueInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerGetAllStaticValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	staticvalueInfo, err := server.store.GetAllStaticValues(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch static value",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.StaticValue `json:"data"` 
	}{
		Status:  true,
		Message: "Static Value retrieved successfully",
		Data:    staticvalueInfo,
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