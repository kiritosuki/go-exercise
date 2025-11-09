package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	resRecord := make([]Record, 0)
	for _, record := range in {
		if predicate(record) {
			resRecord = append(resRecord, record)
		}
	}
	return resRecord
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	from := p.From
	to := p.To
	return func(record Record) bool {
		if record.Day >= from && record.Day <= to {
			return true
		} else {
			return false
		}
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		if record.Category == c {
			return true
		} else {
			return false
		}
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	from := p.From
	to := p.To
	total := 0.0
	for _, record := range in {
		if record.Day >= from && record.Day <= to {
			total += record.Amount
		}
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	err := errors.New("unknown category entertainment")
	total := 0.0
	for _, record := range in {
		if record.Category == c {
			err = nil
			if record.Day >= p.From && record.Day <= p.To {
				total += record.Amount
			}
		}
	}
	if err != nil {
		return 0, err
	}
	return total, nil
}
