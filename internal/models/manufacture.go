package models

type detail struct {
	NeedUpdate bool `json:"needUpdate"`
}

type Manufacturer struct {
	Id      int64   `json:"id"`
	Details *detail `json:"details"`
}
