package models

type HousingResponse struct {
	Housings Housing `json: "housings"`
}

type Housing struct {
	Items []Parking `json: "items"`
}
