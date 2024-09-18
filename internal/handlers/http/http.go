package httphandlers

import (
	"departments-organigram/internal/core/ports"
	"encoding/json"
	"io"
	"net/http"
)

type httpHandler struct {
	departmentsService ports.DepartmentService
	userService        ports.UsersService
}

func NewHTTPHandler(departmentsSrv ports.DepartmentService, usersService ports.UsersService) *httpHandler {
	return &httpHandler{
		departmentsService: departmentsSrv,
		userService:        usersService,
	}
}

func writeError(w http.ResponseWriter, err error, status ...int) {
	responseStatus := http.StatusInternalServerError
	if len(status) > 0 {
		responseStatus = status[0]
	}

	http.Error(w, err.Error(), responseStatus)
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = w.Write(bytesData)
	if err != nil {
		return err
	}
	return nil
}

func readJSON(r *http.Request, data interface{}) error {
	byts, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(byts, data)
}
