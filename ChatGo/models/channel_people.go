package models

type ChannelPeople struct {
	Id             string  `json:"id" gorethink:"id,omitempty"`
	Id_people      string  `json:"id_people" gorethink:"id_people"`
	Id_channel     string  `json:"id_channel" gorethink:"id_channel"`
}
