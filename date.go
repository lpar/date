package date

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// Date represents a SQL date column with no time or timezone information.
type Date time.Time

// NewDate constructs a new Date object for the given year, month and day
func NewDate(y, m, d int) Date {
	return Date(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC))
}

// UnmarshalJSON unmarshals a Date from JSON format. The date is expected
// to be in full-date format as per RFC 3339 -- that is, yyyy-mm-dd.
func (d *Date) UnmarshalJSON(b []byte) error {
	sd, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", sd)
	*d = Date(t)
	return err
}

// MarshalJSON marshals a Date into JSON format. The date is formatted
// in RFC 3339 full-date format -- that is, yyyy-mm-dd.
func (d *Date) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	ds := "\"" + t.Format("2006-01-02") + "\""
	return []byte(ds), nil
}

// Implement Stringer

// String returns the value of the Date in ISO-8601 format.
func (d *Date) String() string {
	return time.Time(*d).Format("2006-01-02")
}

// Implement Valuer

func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

// Implement Scanner

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf("unsupported NULL date.Date value")
	}
	t, ok := value.(time.Time)
	if ok {
		*d = Date(t)
		return nil
	}
	return fmt.Errorf("unable to convert Date")
}
