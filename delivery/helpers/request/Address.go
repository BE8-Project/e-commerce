package request

type InsertAddress struct {
	Address string `json:"address" validate:"required"`
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required"`
	ZipCode int    `json:"zip_code" validate:"required"`
}