package main

import (
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

func (w *Work) Validate() *Error {
	if err := w.validateEmployeeID(); err != nil {
		return err
	}

	if err := w.validateYear(); err != nil {
		return err
	}

	if err := w.validateMonth(); err != nil {
		return err
	}

	if err := w.validateDay(); err != nil {
		return err
	}

	if err := w.validateStartTime(); err != nil {
		return err
	}

	if err := w.validateEndTime(); err != nil {
		return err
	}

	if err := w.validateBreakTime(); err != nil {
		return err
	}

	return nil
}

func (w *Work) validateEmployeeID() *Error {
	if w.EmployeeID != 0 {
		return nil
	}

	return NewError(400, "invalid employee_id.")
}

func (w *Work) validateYear() *Error {
	if 2000 <= w.Year && w.Year <= 2100 {
		return nil
	}

	return NewError(400, "invalid year.")
}

func (w *Work) validateMonth() *Error {
	if 1 <= w.Month && w.Month <= 12 {
		return nil
	}

	return NewError(400, "invalid month.")
}

func (w *Work) validateDay() *Error {
	if 1 <= w.Day && w.Day <= 31 {
		return nil
	}

	return NewError(400, "invalid day.")
}

func (w *Work) validateStartTime() *Error {
	if w.StartTime <= 24 {
		return nil
	}

	return NewError(400, "invalid start_time.")
}

func (w *Work) validateEndTime() *Error {
	if w.EndTime <= 24 {
		return nil
	}

	return NewError(400, "invalid end_time.")
}

func (w *Work) validateBreakTime() *Error {
	if w.BreakTime <= 24 {
		return nil
	}

	return NewError(400, "invalid break_time.")
}
