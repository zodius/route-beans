package model

type Gateway struct {
	Name string
	IP   string
}

type Route struct {
	Dst     string
	Gateway Gateway
}

type Profile struct {
	Name           string
	Gateways       []Gateway
	Routes         []Route
	DefaultGateway Gateway
}

type ProfileRepo interface {
	LoadProfileFile(profile_path string) (profile Profile, err error)
}
