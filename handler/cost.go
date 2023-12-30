package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type Cost struct {
	ID          int32            `json:"id"`
	CostInputID pgtype.Int4      `json:"cost_input_id"`
	CostTypeID  pgtype.Int4      `json:"cost_type_id"`
	PeriodID    pgtype.Int4      `json:"period_id"`
	Title       string           `json:"title"`
	IsDeleted   pgtype.Bool      `json:"is_deleted"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreateCost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	cost := Cost{}

	err := json.NewDecoder(r.Body).Decode(&cost)

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
	err = validate.Struct(cost)
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

	arg := db.CreateCostParams{
		CostInputID:    cost.CostInputID,
		CostTypeID:     cost.CostTypeID,
		PeriodID:       cost.PeriodID,
		Title:          cost.Title,
		IsDeleted:  	cost.IsDeleted,
	}

	costInfo, err := server.store.CreateCost(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to create cost")
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.Cost		`json:"data"`
	}{
		Status:  true,
		Message: "Cost created successfully",
		Data:    []db.Cost{costInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func (server *Server) handlerGetAllCost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	costInfo, err := server.store.GetAllCost(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch cost",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.Cost 		`json:"data"` 
	}{
		Status:  true,
		Message: "cost retrieved successfully",
		Data:    costInfo,
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