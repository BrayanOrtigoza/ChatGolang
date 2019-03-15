package models

type Channel struct {
	Id string `json:"id" gorethink:"id,omitempty"`
}