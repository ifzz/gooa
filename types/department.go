package types

type Department struct {
	UUID                  string `json:"uuid"`
	ID                    int64  `json:"id"`
	Name                  string `json:"name"`
	ParentID              int64  `json:"parentid"`
	Order                 int64  `json:"order"`
	CreateDeptGroup       bool   `json:"createDeptGroup"`
	AutoAddUser           bool   `json:"autoAddUser"`
	DeptHiding            bool   `json:"deptHiding"`
	DeptPermits           string `json:"deptPermits"`
	UserPermits           string `json:"userPermits"`
	OuterDept             bool   `json:"outerDept"`
	OuterPermitDepts      string `json:"outerPermitDepts"`
	OuterPermitUsers      string `json:"outerPermitUsers"`
	OrgDeptOwner          string `json:"orgDeptOwner"`
	DeptManagerUseridList string `json:"deptManagerUseridList"`
	SourceIdentifier      string `json:"sourceIdentifier"`
	GroupContainSubDept   bool   `json:"groupContainSubDept"`
	Ext                   string `json:"ext"`
}

type DepartmentListDTO struct {
	DTOBase
	Department []Department `json:"department"`
}

type DepartmentDTO struct {
	DTOBase
	Department
}
