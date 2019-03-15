package models

import "time"

type Message struct {
	Id          string    `json:"id" gorethink:"id,omitempty"`
	Message        string    `json:"message" gorethink:"message"`
	Id_channel  string    `json:"id_channel"  gorethink:"id_channel"`
	Id_people   string    `json:"id_people_message"  gorethink:"id_people_message"`
	Author      string    `json:"author" gorethink:"author"`
	CreatedAt   time.Time `json:"createdAt" gorethink:"createdAt"`
}