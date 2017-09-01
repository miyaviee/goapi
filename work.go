package main

import (
	"errors"

	"github.com/naoina/genmai"
)

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

func (w *Work) TableName() string {
	return "works"
}

func (w *Work) Validate() error {
	if err := w.validateMonth(); err != nil {
		return err
	}

	if err := w.validateDay(); err != nil {
		return err
	}

	return nil
}

func (w *Work) validateMonth() error {
	if 1 <= w.Month && w.Month <= 12 {
		return nil
	}

	return errors.New("invalid month.")
}

func (w *Work) validateDay() error {
	if 1 <= w.Day && w.Day <= 31 {
		return nil
	}

	return errors.New("invalid day.")
}
