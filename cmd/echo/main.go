package main

import (
    "context"
    "log"

    "github.com/dvs-crcr/go-bx-push-server/internal/echo"
)

func main() {
    config := NewDefaultConfig()
    config.parseVariables()

    // TODO: connect logger

    if err := execute(config); err != nil {
        log.Fatal(err)
    }
}

func execute(cfg *Config) error {
    ctx := context.Background()

    echoService := echo.NewEchoService(
        echo.WithAddress(cfg.Address),
    )

    return echoService.Start(ctx)
}
