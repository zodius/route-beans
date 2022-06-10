package model

type RouteRepo interface {
	FlushRoutingTable() (err error)
	AddRouting(dst string, gateway string) (err error)
}
