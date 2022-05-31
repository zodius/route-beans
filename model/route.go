package model

type RouteService interface {
	FlushRoutingTable() (err error)
	ApplyProfile(profilePath string) (err error)
}

type RouteTableRepo interface {
	FlushRoutingTable() (err error)
	AddRouting(dst string, gateway string) (err error)
}
