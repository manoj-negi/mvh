package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type CostType struct {
	ID          int32            `json:"id"`
	Description pgtype.Text      `json:"description" validate:"required"`
	IsDeleted   pgtype.Bool      `json:"is_deleted"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreateCostType(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	costtype := CostType{}

	err := json.NewDecoder(r.Body).Decode(&costtype)

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
	err = validate.Struct(costtype)
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

	arg := db.CreateCostTypeParams{
		Description:    costtype.Description.String,
		IsDeleted:  	costtype.IsDeleted,
	}

	costtypeInfo, err := server.store.CreateCostType(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to create cost type")
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.CostType 	`json:"data"`
	}{
		Status:  true,
		Message: "Cost Type created successfully",
		Data:    []db.CostType{costtypeInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func (server *Server) handlerGetAllCostType(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	costtypeInfo, err := server.store.GetAllCostType(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch cost type",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.CostType 	`json:"data"` 
	}{
		Status:  true,
		Message: "costtype retrieved successfully",
		Data:    costtypeInfo,
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