package date

import (
	"time"
)

// Date is type conforming to RFC 3339
// see https://www.ietf.org/rfc/rfc3339.txt
type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Format("2006-01-02") + `"`), nil
}

func (d *Date) UnMarshalJSON(data []byte) error {
	var err error
	*d, err = Parse(`"2006-01-02"`, string(data))
	return err
}
