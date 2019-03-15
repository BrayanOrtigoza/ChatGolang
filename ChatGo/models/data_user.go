package models

type DataUser struct {
	Id             string  `json:"id" gorethink:"id,omitempty"`
	Username       string  `json:"username" gorethink:"username"`
	Password       string  `json:"password" gorethink:"password"`
	ResetPassword  string  `json:"reset_password" gorethink:"reset_password"`
	Status         int     `json:"status" gorethink:"status"`
}


