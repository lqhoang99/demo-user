package zookeeper

import (
	"fmt"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"

	"demo-user/config"
)

var conn *zk.Conn

// Connect ...
func Connect() {
	var (
		uri     = os.Getenv("ZOOKEEPER_URI")
		envVars = config.GetEnv()
	)

	// Connect
	conn, _, err := zk.Connect([]string{uri}, time.Second*30)
	if err != nil {
		fmt.Println("ZookeeperURI:", uri)
		panic(err)
	}
	fmt.Println("Zookeeper Connected to", uri)

	// Get env key
	// App port
	appUserPort, _, _ := conn.Get("/app/port/user")
	envVars.AppPort = string(appUserPort)

	// Database
	databaseURI, _, _ := conn.Get("/database/uri")
	envVars.Database.URI = string(databaseURI)
	databaseUserName, _, _ := conn.Get("/database/name/user")
	envVars.Database.Name = string(databaseUserName)
	databaseTestName, _, _ := conn.Get("/database/test/user")
	envVars.Database.TestName = string(databaseTestName)

	// gRPCAddresses
	grpcAddressUser, _, _ := conn.Get("/grpc/uri/user")
	envVars.GRPCAddresses.User = string(grpcAddressUser)
	grpcAddressTransaction, _, _ := conn.Get("/grpc/uri/transaction")
	envVars.GRPCAddresses.Transaction = string(grpcAddressTransaction)

	// gRPCPorts
	grpcPortUser, _, _ := conn.Get("/grpc/port/user")
	envVars.GRPCPorts.User = string(grpcPortUser)
	grpcPortTransaction, _, _ := conn.Get("/grpc/port/transaction")
	envVars.GRPCPorts.Transaction = string(grpcPortTransaction)
}
