package handler

import (
	"backend/go/rest/model/response"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handlers) GetBlogByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errResp := response.APIResponse{
			Data:   nil,
			Errors: "id is not a number",
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	ctx := context.Background()
	resp, statusCode, err := h.blogUC.GetBlogByID(ctx, id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		errResp := response.APIResponse{
			Data:   nil,
			Errors: err.Error(),
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	apiResp := response.APIResponse{
		Data: resp,
	}
	json.NewEncoder(w).Encode(apiResp)
}
