package models

import (
	"time"
)

type Apartment struct {
	// ID        int       `gorm:"primarykey" json:"neid,omitempty"`
	CreatedAt time.Time `json:",omitempty"`

	ID                  string `gorm:"primarykey" json: "id"`
	Name                string `json: "name"`
	Object              string `json: "object"`
	Object_id           string `json: "object_id"`
	Object_code         string `json: "object_code"`
	Number              string `json: "number"`
	Rooms               string `json: "rooms"`
	Floor               string `json: "floor"`
	Block               string `json: "block"`
	Area                string `json: "area"`
	Price               string `json: "price"`
	Price_m             string `json: "price_m"`
	Plan_s              string `json: "plan_s"`
	Plan                string `json: "plan"`
	App_type            string `json: "type"`
	Term_of_application string `json: "term_of_application"`
	Open_sale           int    `json: "open_sale"`
	Y2_sell             int    `json: "y2_sell"`
	For_sell            int    `json: "for_sell"`
	Num_on_floor        string `json: "num_on_floor"`
	Property            string `json: "property"`
	Requested           int    `json: "requested"`
	Block_name          string `json: "block_name"`
	Priority_date_end   string `json: "priority_date_end"`
}
