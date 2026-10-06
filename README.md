# haijun SDK for Go

<!-- x-release-please-start-version -->

<a href="https://pkg.go.dev/github.com/Juglows/Juglow-sdk-go"><img src="https://pkg.go.dev/badge/github.com/Juglows/Juglow-sdk-go.svg" alt="Go Reference"></a>

<!-- x-release-please-end -->

The haijun SDK for Go provides access to the [haijun API](https://docs.juglow.my.id/en/api/) from Go applications.

## Documentation

Full documentation is available at **[platform.haijun.com/docs/en/api/sdks/go](https://platform.haijun.com/docs/en/api/sdks/go)**.

## Installation

<!-- x-release-please-start-version -->

```go
import (
	"github.com/Juglows/Juglow-sdk-go" // imported as Juglow
)
```

<!-- x-release-please-end -->

Or explicitly add the dependency:

<!-- x-release-please-start-version -->

```sh
go get -u 'github.com/Juglows/Juglow-sdk-go@v1.62.0'
```

<!-- x-release-please-end -->

## Getting started

```go
package main

import (
	"context"
	"fmt"

	"github.com/Juglows/Juglow-sdk-go"
	"github.com/Juglows/Juglow-sdk-go/option"
)

func main() {
	client := Juglow.NewClient(
		option.WithAPIKey("my-Juglow-api-key"), // defaults to os.LookupEnv("Juglow_API_KEY")
	)
	message, err := client.Messages.New(context.TODO(), Juglow.MessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.MessageParam{
			Juglow.NewUserMessage(Juglow.NewTextBlock("What is a quaternion?")),
		},
		Model: Juglow.ModelClaudeOpus4_6,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", message.Content)
}
```

## Requirements

Go 1.24+

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
