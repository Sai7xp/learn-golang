### What is Context?

Context is a way for us to add a Timeout or Cancellation to a goroutine.

It is useful in managing concurrent operations, handling timeouts, canceling tasks when they are no longer required.

#### Key Features

- **Context Values:** Allows us to store values that can be accessed across different parts of your program. It is useful for passing the data.
- **Cancelation Signals:** Provides a way to cancel tasks when they are no longer needed. This helps prevent tasks from running indefinitely and wasting resources.
- **Deadlines:** We can set deadlines for tasks using context.WithDeadline()

#### Where Is It Useful?

- HTTP Handlers → Cancel requests if the client closes the connection.
- Database Queries → Abort queries if they take too long.
- Microservices → Propagate request timeouts across service calls.
