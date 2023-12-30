package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type Period struct {
	ID        int32            `json:"id"`
	Period    pgtype.Text      `json:"period" validate:"required"`
	IsDeleted pgtype.Bool      `json:"is_deleted"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
func (server *Server) handlerCreatePeriod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	period := Period{}

	err := json.NewDecoder(r.Body).Decode(&period)

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
	err = validate.Struct(period)
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

	arg := db.CreatePeriodParams{
		Period:    period.Period.String,
		IsDeleted:  period.IsDeleted,
	}

	periodInfo, err := server.store.CreatePeriod(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to create period")
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.Period 	`json:"data"`
	}{
		Status:  true,
		Message: "Period created successfully",
		Data:    []db.Period{periodInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}


func (server *Server) handlerGetAllPeriod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	periodInfo, err := server.store.GetAllPeriod(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch period",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool            `json:"status"`
		Message string          `json:"message"`
		Data    []db.Period 		`json:"data"` 
	}{
		Status:  true,
		Message: "period retrieved successfully",
		Data:    periodInfo,
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