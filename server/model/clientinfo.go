package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Clientinfo struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"type:uuid"` 
	Quantity              string    `json:"quantity"`
	Unit                  string    `json:"unit"`
	StartDate             time.Time `json:"start_date"`
	UUID                  uuid.UUID `json:"uuid"`
	Count                 int       `json:"count"`
	EndDate               time.Time `json:"end_date"`
	Value                 float64   `json:"value"`
	SourceName            string    `json:"source_name"`
	OperatingSystemVersion string    `json:"operating_system_version"`
	SourceBundleIdentifier string    `json:"source_bundle_identifier"`
	QuantityType          string    `json:"quantity_type"`
	ClientID              string    `json:"client_id"`
	SourceProductType     string    `json:"source_product_type"`
	QuantityTypeOther     string    `json:"quantity_type_other"`
}