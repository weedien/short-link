package link

import (
	"database/sql/driver"
	"errors"
)

type Status string

const (
	StatusActive    Status = "active"    // 激活状态
	StatusExpired   Status = "expired"   // 已过期
	StatusDisabled  Status = "disabled"  // 已停用
	StatusForbidden Status = "forbidden" // 已禁用
	StatusReserved  Status = "reserved"  // 预留状态
	StatusDeleted   Status = "deleted"   // 已删除
)

// Scan implements the Scanner interface for database deserialization
func (s *Status) Scan(value interface{}) error {
	strVal, ok := value.(string)
	if !ok {
		return errors.New("invalid data type for Status")
	}
	*s = Status(strVal)
	return nil
}

// Value implements the Valuer interface for database serialization
func (s *Status) Value() (driver.Value, error) {
	return string(*s), nil
}

func (s *Status) String() string {
	return string(*s)
}

func (s *Status) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

type ValidType int

const (
	// ValidTypePermanent 永久有效
	ValidTypePermanent ValidType = iota
	// ValidTypeTemporary 临时有效
	ValidTypeTemporary
)

var validTypeMap = map[ValidType]string{
	ValidTypePermanent: "永久有效",
	ValidTypeTemporary: "临时有效",
}

// Scan implements the sql.Scanner interface
func (v *ValidType) Scan(value interface{}) error {
	intVal, ok := value.(int64)
	if !ok {
		return errors.New("invalid data type for ValidType")
	}
	*v = ValidType(intVal)
	return nil
}

// Value implements the driver.Valuer interface
func (v *ValidType) Value() (driver.Value, error) {
	return int(*v), nil
}

func (v *ValidType) String() string {
	return validTypeMap[*v]
}

func (v *ValidType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + v.String() + `"`), nil
}

type CreateType int

const (
	CreateByApi CreateType = iota
	CreateByConsole
)

var createTypeMap = map[CreateType]string{
	CreateByApi:     "Api",
	CreateByConsole: "Console",
}

// Scan implements the sql.Scanner interface
func (c *CreateType) Scan(value interface{}) error {
	intVal, ok := value.(int64)
	if !ok {
		return errors.New("invalid data type for CreateType")
	}
	*c = CreateType(intVal)
	return nil
}

// Value implements the driver.Valuer interface
func (c *CreateType) Value() (driver.Value, error) {
	return int(*c), nil
}

func (c *CreateType) String() string {
	return createTypeMap[*c]
}

func (c *CreateType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}
