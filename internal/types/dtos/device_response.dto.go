package dtos

type DeviceDetails struct {
	Id           int64  `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Manufacturer string `json:"manufacturer" validate:"required"`
	Year         int    `json:"year" validate:"required"`
}
