# Events

## ⚡️ Quickstart

```go
package main

import "fmt"
import "github.com/ichbinbekir/events"

func main() {
	eventEmitter := events.NewEventEmitter()

	eventEmitter.On("message", func(args ...any) {
		fmt.Println(args)
	})

	eventEmitter.Emit("message", "hello") 
}
```
