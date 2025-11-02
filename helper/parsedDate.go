package helper

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		cd.Time = time.Time{}
		return nil
	}

	// hapus tanda kutip
	s = s[1 : len(s)-1]

	// Coba format RFC3339 dulu
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		cd.Time = t
		return nil
	}

	// Coba format DATE (YYYY-MM-DD)
	if t, err := time.Parse("2006-01-02", s); err == nil {
		cd.Time = t
		return nil
	}

	return fmt.Errorf("invalid date format: %s (use YYYY-MM-DD or RFC3339)", s)
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	if cd.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(cd.Format("2006-01-02"))
}
