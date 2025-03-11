package models

type CheckRequest struct {
	CareGiverId     string `json:"cg"`
	CareReceipentId string `json:"cr"`
	Action          string `json:"action"`
	Resource        string `json:"resource"`
}
