package model

type Route struct {
}

type RouteService interface {
	FlushRoutingTable() (err error)
	LoadFromProfile(profilePath string) (err error)
}

type RouteTableRepo interface {
	FlushRoutingTable() (err error)
	// DeleteRouting() (err error)
	// AddRouting() (err error)
}

type ProfileRepo interface {
}
