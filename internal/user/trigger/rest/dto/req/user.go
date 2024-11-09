package req

type RecycleBinRecoverReq struct {
	Gid          string `json:"gid"`
	FullShortUrl string `json:"full_short_url"`
}

type RecycleBinRemoveReq struct {
	Gid          string `json:"gid"`
	FullShortUrl string `json:"full_short_url"`
}

type RecycleBinSaveReq struct {
	Gid          string `json:"gid"`
	FullShortUrl string `json:"full_short_url"`
}

type LinkGroupSaveReq struct {
	Name string `json:"name"`
}

type LinkGroupSortReq []LinkGroupSortDto

type LinkGroupSortDto struct {
	Gid       string `json:"gid"`
	SortOrder int    `json:"sort_order"`
}

type LinkGroupUpdateReq struct {
	Gid  string `json:"gid"`
	Name string `json:"name"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RealName string `json:"real_name"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
}

type UserUpdateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RealName string `json:"real_name"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
}
