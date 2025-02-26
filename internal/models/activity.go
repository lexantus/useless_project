package models

type Activity struct {
	ID         int    `json:"id"`
	TimespanMs int    `json:"timespan_ms"`
	Desc       string `json:"description"`
}
