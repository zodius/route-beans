package model

type RouteBeanService interface {
	ListProfiles() (err error)
	ApplyProfile(profileName string) (err error)
	Reset() (err error)
}
