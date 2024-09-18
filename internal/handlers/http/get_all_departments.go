package httphandlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *httpHandler) GetAllDepartments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	departments, err := h.departmentsService.GetAllDepartments(context.Background())
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(departments)
	return

}
