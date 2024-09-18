package httphandlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (h *httpHandler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	var id string
	if len(parts) >= 3 {
		id = parts[2]
	} else {
		writeError(w, errors.New("invalid url"), http.StatusBadRequest)
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id is not an integer", http.StatusBadRequest)
		return
	}

	err = h.departmentsService.DeleteDepartment(context.Background(), idInt)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	return
}
