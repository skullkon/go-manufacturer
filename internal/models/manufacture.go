package models

type Detail struct {
	NeedUpdate bool `json:"needUpdate" db:"needUpdate"`
}

type Manufacturer struct {
	Id      int64   `json:"id" db:"id"`
	Details *Detail `json:"details" db:"details"`
}
