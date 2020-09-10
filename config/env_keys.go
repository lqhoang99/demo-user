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

	// gRPC addresses
	GRPCAddresses struct {
		User        string
		Transaction string
	}

	// gRPC ports
	GRPCPorts struct {
		User        string
		Transaction string
	}
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
