package device

import (
	"devices-api/internal/types/dtos"
	"devices-api/internal/utils/helpers"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func CreateDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var device dtos.Device

		err := json.NewDecoder(r.Body).Decode(&device)

		if errors.Is(err, io.EOF) {
			helpers.WriteJsonResponse(w, http.StatusBadRequest, helpers.GeneralError(fmt.Errorf("empty body")))
			return
		}

		slog.Info("Creating device")

		helpers.WriteJsonResponse(w, http.StatusCreated, map[string]string{"msg": "created"})
	}
}
