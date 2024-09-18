package httphandlers

import (
	"context"
	"departments-organigram/internal/core/domain"
	"errors"
	"net/http"
)

func (h *httpHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	department := domain.Department{}

	if err := readJSON(r, &department); err != nil {
		writeError(w, err)
		return
	}

	err := h.departmentsService.CreateDepartment(context.Background(), department)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	return
}
