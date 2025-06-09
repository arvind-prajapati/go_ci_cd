package main

import (
	"context"
	"log"

	"google.golang.org/api/idtoken"
)

func main() {
	_, err := idtoken.NewClient(context.Background(), "https://kyc-dev-pcdsimq4za-ey.a.run.app")
	if err != nil {
		log.Fatalf("failed to create ID token client: %v", err)
	}

}
