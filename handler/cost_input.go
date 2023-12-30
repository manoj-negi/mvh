package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type CostInput struct {
	ID        int32            `json:"id"`
	Type      pgtype.Text      `json:"type" validate:"required"`
	IsDeleted pgtype.Bool      `json:"is_deleted"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreateCostInput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}

	ctx := r.Context()
	costInput := CostInput{}

	err := json.NewDecoder(r.Body).Decode(&costInput)

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
	err = validate.Struct(costInput)
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

	arg := db.CreateCostInputParams{
		Type:    costInput.Type.String,
		IsDeleted:  costInput.IsDeleted,
	}

	costInputInfo, err := server.store.CreateCostInput(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to create cost input")
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.CostInput 	`json:"data"`
	}{
		Status:  true,
		Message: "cost input created successfully",
		Data:    []db.CostInput{costInputInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func (server *Server) handlerGetAllCostInput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	costInputInfo, err := server.store.GetAllCostInput(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch cost input",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.CostInput 	`json:"data"` 
	}{
		Status:  true,
		Message: "cost input retrieved successfully",
		Data:    costInputInfo,
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