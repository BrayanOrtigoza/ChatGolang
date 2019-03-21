package models

type Channel struct {
	Id string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
	Type string `json:"type" gorethink:"type"`
}