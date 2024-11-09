package types

import (
	"database/sql/driver"
	"github.com/bytedance/sonic"
	"shortlink/internal/base/errno"
	"time"
)

// JsonTime is a custom time type for JSON serialization
type JsonTime time.Time

func (jt *JsonTime) ToTime() *time.Time {
	return (*time.Time)(jt)
}

// MarshalJSON formats the time as a JSON string
func (jt *JsonTime) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(time.Time(*jt).Format("2006-01-02 15:04:05"))
}

// UnmarshalJSON parses the time from a JSON string
func (jt *JsonTime) UnmarshalJSON(data []byte) error {
	// Check if the input is an empty string
	if string(data) == `""` {
		*jt = JsonTime(time.Time{})
		return nil
	}

	// Parse the time from the JSON string
	var err error
	t, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return errno.NewRequestError("invalid time format")
	}
	*jt = JsonTime(t)
	return nil
}

// Scan implements the Scanner interface for database deserialization
func (jt *JsonTime) Scan(value interface{}) error {
	if value == nil {
		*jt = JsonTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*jt = JsonTime(v)
		return nil
	default:
		return errno.NewRequestError("invalid time format")
	}
}

// Value implements the Valuer interface for database serialization
func (jt *JsonTime) Value() (driver.Value, error) {
	return time.Time(*jt), nil
}
