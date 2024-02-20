package models

type Device struct {
	ID       string
	Name     string
	IP_addr  string
	Model    string
	MAC_addr string
}
type Search struct {
	Name     string
	IP_addr  string
	Model    string
	MAC_addr string
}

type Event struct {
	Source_addr string
	Dest_addr   string
	ID          string
	Event_code  string
	Description string
	Timestamp   string
}
