### What is Context?

Context package is used for carrying cancelation signals, deadlines.

Context is a way for us to add a Timeout or Cancellation to a goroutine.

It is useful in managing concurrent operations, handling timeouts, canceling tasks when they are no longer required.

#### Key Features

**Context Values:** Allows us to store values that can be accessed across different parts of your program. It is useful for passing the data.

```go
ctx := context.WithValue(context.Background(), "key", "example value")

value := ctx.Value("key")
fmt.Println(value) // Output: example value
```

**Cancelation Signals:** Provides a way to cancel tasks when they are no longer needed. This helps prevent tasks from running indefinitely and wasting resources.

**Deadlines:** We can set deadlines for tasks using context.WithDeadline()

```go
deadline := time.Now().Add(1 * time.Minute)
ctx, cancel := context.WithDeadline(context.Background(), deadline)

// timeout
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
```

#### context.Background() vs context.TODO()

Both returns an empty context(context.Context object). They are often used to create root contexts.
`context.Background()`: Use when there is no existing context, and you want to start a new one. Typically used in top-level functions like main() and init()

`context.TODO()`: When context usage is undecided or TBD

#### Where Is It Useful?

- HTTP Handlers → Cancel requests if the client closes the connection.
- Database Queries → Abort queries if they take too long.
- Microservices → Propagate request timeouts across service calls.
- Pass data from middlewares to the main http handlers using context - check [`web-auth-go`][def] jwt token example

[def]: https://github.com/Sai7xp/web-auth-go/blob/main/internal/05_jwt_auth_api.go
