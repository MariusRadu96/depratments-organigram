package httphandlers

import (
	"context"
	"departments-organigram/internal/core/domain"
	"errors"
	"net/http"
)

func (h *httpHandler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	department := domain.Department{}

	if err := readJSON(r, &department); err != nil {
		writeError(w, err)
		return
	}

	err := h.departmentsService.UpdateDepartment(context.Background(), department)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	return
}
