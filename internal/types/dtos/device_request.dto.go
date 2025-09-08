package dtos

type CreateDevice struct {
	Name         string `json:"name" validate:"required"`
	Manufacturer string `json:"manufacturer" validate:"required"`
	Year         int    `json:"year" validate:"required"`
}
