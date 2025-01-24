# Go BX Push Server

## Configuration
Application configuration can be done via command-line flags or environment variables.

| Flag |    Env    | Default value  | Description             |
|:----:|:---------:|----------------|-------------------------|
|  a   |  ADDRESS  | localhost:8080 | server endpoint address |
|  l   | LOG_LEVEL | info           | logging level           |


## Make targets

- `make build` — build all commands
- `make build-echo` — build test echo server
- `make lint` — analyse code for potential errors