package echo

import (
    "context"
)

func main() {
    config := NewDefaultConfig()
    config.parseVariables()

    // TODO: connect logger

    execute(config)
}

func execute(_ *Config) {
    _ = context.Background()

    // TODO: implement some methods...
}
