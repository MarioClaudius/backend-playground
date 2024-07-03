package handler

import (
	"backend/go/rest/model/request"
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

func (h *Handlers) GetBlogList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errResp := response.APIResponse{
			Data:   nil,
			Errors: "offset is not a number",
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errResp := response.APIResponse{
			Data:   nil,
			Errors: "limit is not a number",
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	filter := request.BlogListFilter{
		Title:  r.URL.Query().Get("title"),
		Offset: offset,
		Limit:  limit,
	}

	resp, statusCode, err := h.blogUC.GetBlogList(r.Context(), filter)
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

func (h *Handlers) AddBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := request.CreateBlogRequest{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResp := response.APIResponse{
			Data:   nil,
			Errors: "invalid request",
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	statusCode, err := h.blogUC.CreateBlog(r.Context(), req)
	if err != nil {
		w.WriteHeader(statusCode)
		errResp := response.APIResponse{
			Data:   nil,
			Errors: err.Error(),
		}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	apiResp := response.APIResponse{
		Data: "Blog created successfully",
	}
	json.NewEncoder(w).Encode(apiResp)
}
