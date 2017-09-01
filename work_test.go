package main

import "testing"

func getWork() Work {
	return Work{
		Year:      2017,
		Month:     1,
		Day:       1,
		StartTime: 24,
		EndTime:   24,
		BreakTime: 24,
	}
}

func TestTableName(t *testing.T) {
	w := getWork()
	if w.TableName() != "works" {
		t.Fatalf("table name test failed")
	}
}

func TestValidateMonth(t *testing.T) {
	w := getWork()
	if err := w.Validate(); err != nil {
		t.Fatalf("validate month test failed")
	}

	w.Month = 12
	if err := w.Validate(); err != nil {
		t.Fatalf("validate month test failed")
	}

	w.Month = 0
	if err := w.Validate(); err == nil {
		t.Fatalf("validate month test failed")
	}

	w.Month = 13
	if err := w.Validate(); err == nil {
		t.Fatalf("validate month test failed")
	}
}

func TestValidateDay(t *testing.T) {
	w := getWork()
	if err := w.Validate(); err != nil {
		t.Fatalf("validate day test failed")
	}

	w.Day = 31
	if err := w.Validate(); err != nil {
		t.Fatalf("validate day test failed")
	}

	w.Day = 0
	if err := w.Validate(); err == nil {
		t.Fatalf("validate day test failed")
	}

	w.Day = 32
	if err := w.Validate(); err == nil {
		t.Fatalf("validate day test failed")
	}
}

func TestValidateStartTime(t *testing.T) {
	w := getWork()
	w.StartTime = 24.01
	if err := w.Validate(); err == nil {
		t.Fatalf("validate start time test failed")
	}
}

func TestValidateEndTime(t *testing.T) {
	w := getWork()
	w.EndTime = 24.01
	if err := w.Validate(); err == nil {
		t.Fatalf("validate end time test failed")
	}
}

func TestValidateBreakTime(t *testing.T) {
	w := getWork()
	w.BreakTime = 24.01
	if err := w.Validate(); err == nil {
		t.Fatalf("validate break time test failed")
	}
}
