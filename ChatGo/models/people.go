package models

type People struct {
	Id       string  `json:"id" gorethink:"id,omitempty"`
	Name     string  `json:"name" gorethink:"name"`
	LastName string  `json:"last_name"  gorethink:"last_name"`
	Status   int     `json:"status" gorethink:"status"`
	IdUser   string  `json:"id_user" gorethink:"id_user"`
}
