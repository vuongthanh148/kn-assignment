package infrastructure

import (
	"context"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
)

func NewSecretManagerClient(ctx context.Context) *secretmanager.Client {
	secretCli, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("unable to init secret manager: %v", err)
	}

	return secretCli
}
