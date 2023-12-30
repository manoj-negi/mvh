package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type CostCalculation struct {
	ID                  int32            `json:"id"`
	Parkcost            pgtype.Int4      `json:"parkcost"`
	Assetcost           pgtype.Int4      `json:"assetcost"`
	Loancost            pgtype.Int4      `json:"loancost"`
	TaxCost             pgtype.Int4      `json:"tax_cost"`
	AssetValueCost      pgtype.Int4      `json:"asset_value_cost"`
	Revenue             pgtype.Int4      `json:"revenue"`
	TotalCost           pgtype.Int4      `json:"total_cost"`
	Result              pgtype.Int4      `json:"result"`
	ResultPerc          pgtype.Int4      `json:"result_perc"`
	SavingRevenueAmount pgtype.Int4      `json:"saving_revenue_amount"`
	SavingRevenuePerc   pgtype.Int4      `json:"saving_revenue_perc"`
	YieldDiff           pgtype.Int4      `json:"yield_diff"`
	IsDeleted           pgtype.Bool      `json:"is_deleted"`
	CreatedAt           pgtype.Timestamp `json:"created_at"`
	UpdatedAt           pgtype.Timestamp `json:"updated_at"`
}

func (server *Server) handlerCreateCostCalculation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	costcalculation := CostCalculation{}
	err := json.NewDecoder(r.Body).Decode(&costcalculation)

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
	err = validate.Struct(costcalculation)
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

	arg := db.CreateCostCalculationParams{
		Parkcost:            costcalculation.Parkcost,
		Assetcost:           costcalculation.Assetcost,
		Loancost:            costcalculation.Loancost,
		TaxCost:             costcalculation.TaxCost,
		AssetValueCost:      costcalculation.AssetValueCost,
		Revenue:             costcalculation.Revenue,
		TotalCost:           costcalculation.TotalCost,
		Result:              costcalculation.Result,
		ResultPerc:          costcalculation.ResultPerc,
		SavingRevenueAmount: costcalculation.SavingRevenueAmount,
		SavingRevenuePerc:   costcalculation.SavingRevenuePerc,
		YieldDiff:           costcalculation.YieldDiff,
		IsDeleted:           costcalculation.IsDeleted,
	}

	costCalculationInfo, err := server.store.CreateCostCalculation(ctx, arg)
	if err != nil {
		util.ErrorResponse(w, http.StatusInternalServerError, "Failed to create park")
		return
	}

	response := struct {
		Status  bool                 `json:"status"`
		Message string               `json:"message"`
		Data    []db.CostCalculation `json:"data"`
	}{
		Status:  true,
		Message: "Cost Calculation created successfully",
		Data:    []db.CostCalculation{costCalculationInfo},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (server *Server) handlerGetAllCostCalculation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		util.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
		return
	}
	ctx := r.Context()

	calculationCostInfo, err := server.store.GetAllCostCalculation(ctx)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "Failed to fetch cost calculation ",
			StatusCode: http.StatusInternalServerError,
		}
		util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
		return
	}

	response := struct {
		Status  bool                 `json:"status"`
		Message string               `json:"message"`
		Data    []db.CostCalculation `json:"data"`
	}{
		Status:  true,
		Message: "Cost Calculation retrieved successfully",
		Data:    calculationCostInfo,
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
