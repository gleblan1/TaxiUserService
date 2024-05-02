package models

type Trip struct {
	Id       int
	TaxiType int
	From     string
	To       string
	Rate     int
	User     User
}

type TripHistory struct {
	Trips []Trip
	User  User
}
