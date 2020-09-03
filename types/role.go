package types

type Role struct {
	UUID    string `json:"uuid"`
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	GroupId int64  `json:"groupId"`
}

type RoleGroup struct {
	UUID  string `json:"uuid"`
	ID    int64  `json:"groupId"`
	Roles []Role `json:"roles"`
	Name  string `json:"name"`
}

type RoleGroupResult struct {
	HasMore bool        `json:"hasMore"`
	List    []RoleGroup `json:"list"`
}

type RoleGroupResultDTO struct {
	DTOBase
	Result RoleGroupResult `json:"result"`
}

type RoleUserListResult struct {
	HasMore bool   `json:"hasMore"`
	List    []User `json:"list"`
}

type RoleUserListResultDTO struct {
	DTOBase
	Result RoleUserListResult `json:"result"`
}

type UserRole struct {
	UUID     string `json:"uuid"`
	UserId   string `json:"user_id"`
	RoleId   int64  `json:"role_id"`
	UserName string `json:"user_name"`
	RoleName string `json:"role_name"`
}
