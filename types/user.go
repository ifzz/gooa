package types

import "time"

type User struct {
	UUID            string     `json:"uuid"`
	UserId          string     `json:"userid"`
	UnionId         string     `json:"unionid"`
	Name            string     `json:"name"`
	Tel             string     `json:"tel"`
	WorkPlace       string     `json:"workPlace"`
	Remark          string     `json:"remark"`
	Mobile          string     `json:"mobile"`
	Email           string     `json:"email"`
	OrgEmail        string     `json:"orgEmail"`
	Active          bool       `json:"active"`
	OrderInDepts    bool       `json:"orderInDepts"`
	IsAdmin         bool       `json:"isAdmin"`
	IsBoss          bool       `json:"isBoss"`
	IsLeaderInDepts bool       `json:"isLeaderInDepts"`
	IsHide          bool       `json:"isHide"`
	Department      []int64    `json:"department"`
	Position        string     `json:"position"`
	Avatar          string     `json:"avatar"`
	HiredDate       int64      `json:"hiredDate"`
	JobNumber       string     `json:"jobnumber"`
	Extattr         string     `json:"extattr"`
	IsSenior        bool       `json:"isSenior"`
	StateCode       string     `json:"stateCode"`
	RealAuthed      bool       `json:"realAuthed"`
	CreateTime      *time.Time `json:"create_time"`
}

type UserListDTO struct {
	DTOBase
	UserList []User `json:"userlist"`
	HasMore  bool   `json:"hasMore"`
}

type UserDepartment struct {
	UUID         string     `json:"uuid"`
	UserID       string     `json:"user_id"`
	DepartmentId int64      `json:"department_id"`
	CreateTime   *time.Time `json:"create_time"`
}

type UserInfo struct {
	Nick    string `json:"nick"`
	UnionID string `json:"unionid"`
	OpenID  string `json:"openid"`
	DingID  string `json:"dingId"`
}

type UserInfoDTO struct {
	DTOBase
	UserInfo UserInfo `json:"user_info"`
}
