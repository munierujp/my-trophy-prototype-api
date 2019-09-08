package date

import (
	"database/sql/driver"
	"time"
)

const (
	ISO8601 = "2006-01-02"
	RFC3339 = "2006-01-02"
)

func (d Date) String() string {
	return d.Format(RFC3339)
}

func (d Date) Format(layout string) string {
	return d.Time.Format(layout)
}

func Parse(layout, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}
	return Date{t}, nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

func (d *Date) Scan(value interface{}) error {
	d.Time = value.(time.Time)
	return nil
}
