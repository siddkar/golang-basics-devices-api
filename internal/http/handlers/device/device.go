package device

import (
	"devices-api/internal/storage"
	"devices-api/internal/types/dtos"
	"devices-api/internal/utils/helpers"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func CreateDevice(storage storage.Storage) http.HandlerFunc {
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

		fmt.Println("Data", device.Name, device.Manufacturer, device.Year)

		lastId, err := storage.CreateDevice(device.Name, device.Manufacturer, device.Year)

		slog.Info("Device created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			helpers.WriteJsonResponse(w, http.StatusInternalServerError, err)
			return
		}

		slog.Info("Creating device")
		helpers.WriteJsonResponse(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetDeviceById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a device", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			helpers.WriteJsonResponse(w, http.StatusBadRequest, helpers.GeneralError(err))
			return
		}

		rawDevice, err := storage.GetDeviceById(intId)

		if err != nil {
			slog.Error("error getting user", slog.String("id", id))
			helpers.WriteJsonResponse(w, http.StatusInternalServerError, helpers.GeneralError(err))
			return
		}

		var deviceDetails dtos.DeviceDetails = dtos.DeviceDetails(rawDevice)

		helpers.WriteJsonResponse(w, http.StatusOK, deviceDetails)

	}
}
