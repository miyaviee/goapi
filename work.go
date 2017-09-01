package main

import "github.com/naoina/genmai"

type Work struct {
	ID         int64          `db:"pk" json:"id"`
	EmployeeID int64          `json:"employee_id"`
	Year       int64          `json:"year"`
	Month      int64          `json:"month"`
	Day        int64          `json:"day"`
	StartTime  genmai.Float64 `json:"start_time"`
	EndTime    genmai.Float64 `json:"end_time"`
	BreakTime  genmai.Float64 `json:"break_time"`

	genmai.TimeStamp
}

// TableName work table name
func (w *Work) TableName() string {
	return "works"
}

// Validate work validate
func (w *Work) Validate() error {
	return nil
}
