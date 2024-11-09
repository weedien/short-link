package resp

import (
	"encoding/json"
)

type UserLoginResp struct {
	Token string `json:"token"`
}

type UserResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	RealName string `json:"real_name"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
}

func (r UserResp) MarshalJSON() ([]byte, error) {
	type Alias UserResp
	phone := []rune(r.Phone)
	if len(phone) == 11 {
		phone = append(phone[:3], append([]rune("****"), phone[7:]...)...)
	}
	return json.Marshal(&struct {
		Phone string `json:"phone"`
		*Alias
	}{
		Phone: string(phone),
		Alias: (*Alias)(&r),
	})
}

type UserActualResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	RealName string `json:"real_name"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
}

type LinkGroupResp []LinkGroupDTO

type LinkGroupDTO struct {
	Gid       string `json:"gid"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
	LinkCount int    `json:"short_link_count"`
}
