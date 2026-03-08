package main

import (
	"fmt"
	"github.com/msc-privacy-grid-mpc-zkp/cloud-aggregator/internal/api"
)

func main() {
	fmt.Println("☁️  Starting MPC Cloud Aggregator...")
	fmt.Println("---------------------------------------------------------")

	api.StartServer(":8080")
}
