package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/msc-privacy-grid-mpc-zkp/cloud-aggregator/internal/api"
	"github.com/msc-privacy-grid-mpc-zkp/cloud-aggregator/internal/zkp"
)

func main() {
	// 1. Definišemo parametre (zastavice) koje možemo proslediti iz terminala
	port := flag.String("port", "8080", "Port na kom server sluša")
	name := flag.String("name", "Server A", "Ime MPC čvora")
	flag.Parse()

	fmt.Printf("☁️  Starting MPC Cloud Aggregator: [%s]\n", *name)
	fmt.Printf("---------------------------------------------------------\n")

	// 2. Učitavamo ključ sa diska
	vk, err := zkp.LoadVerifyingKey("keys/verifying.key")
	if err != nil {
		log.Fatalf("[FATAL] Failed to load verifying key: %v", err)
	}
	fmt.Println("[SECURITY] ZKP Verifying Key loaded successfully!")

	// 3. Pravimo memoriju koja čeka 10 brojila
	store := api.NewMemoryStore(10)

	// 4. Startujemo server, ali sada koristimo dinamički port!
	address := ":" + *port
	api.StartServer(address, vk, store)
}
