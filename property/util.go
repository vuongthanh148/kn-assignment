package property

import (
	"context"
	"log"

	"github.com/kelseyhightower/envconfig"
)

func InitProperty(ctx context.Context) {
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}

}
