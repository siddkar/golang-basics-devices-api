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

	"github.com/go-playground/validator/v10"
)

func CreateDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var device dtos.CreateDevice

		err := json.NewDecoder(r.Body).Decode(&device)

		if errors.Is(err, io.EOF) {
			helpers.WriteJsonResponse(w, http.StatusBadRequest, helpers.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			helpers.WriteJsonResponse(w, http.StatusBadRequest, helpers.GeneralError(err))
			return
		}

		// request validation
		if err := validator.New().Struct(device); err != nil {
			validationErrors := err.(validator.ValidationErrors) // Typecasting
			helpers.WriteJsonResponse(w, http.StatusBadRequest, helpers.ValidationError(validationErrors))
			return
		}

		slog.Info("Creating device")
		helpers.WriteJsonResponse(w, http.StatusCreated, map[string]string{"msg": "created"})
	}
}
