package types

type OpenExtContact struct {
	UUID           string   `json:"uuid"`
	Title          string   `json:"title"`
	ShareDeptIds   []int64  `json:"share_dept_ids"`
	LabelIds       []int64  `json:"label_ids"`
	Remark         string   `json:"remark"`
	Address        string   `json:"address"`
	Name           string   `json:"name"`
	FollowerUserId string   `json:"follower_user_id"`
	StateCode      string   `json:"state_code"`
	CompanyName    string   `json:"company_name"`
	ShareUserIds   []string `json:"share_user_ids"`
	Mobile         string   `json:"mobile"`
	UserId         string   `json:"userid"`
	Email          string   `json:"email"`
}

type OpenExtContactDTO struct {
	DTOBase
	Results []OpenExtContact `json:"results"`
}
