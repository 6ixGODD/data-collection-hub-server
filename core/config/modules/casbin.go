package modules

type CasbinConfig struct {
	ModelPath        string
	PolicyPath       string
	DbSpec           string
	EnableLog        bool
	EnableHttp       bool
	HttpPort         int
	EnableAutoLoad   bool
	AutoLoadInternal int
}
