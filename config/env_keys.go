package config

// ENV ...
type ENV struct {
	IsDev bool

	// ZookeeperURI
	ZookeeperURI     string

	// App port
	AppPort string

	// Database
	Database struct {
		URI            string
		TestName       string
		Name       string
	}

	GRPCUri string
}

var env ENV

// InitENV ...
func InitENV() {
	env = ENV{
		IsDev: true,
	}
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
