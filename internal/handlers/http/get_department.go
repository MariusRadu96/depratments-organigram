package httphandlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (h *httpHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, errors.New("id not provided"), http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id is not an integer", http.StatusBadRequest)
		return
	}

	department, err := h.departmentsService.GetDepartmentByID(context.Background(), idInt)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(department)
	return
}
