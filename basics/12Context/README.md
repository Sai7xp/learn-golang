### What is Context?

Context package is a standard library that helps manage the data flow of data and operation across different parts of a program.

It is useful in managing concurrent operations, handling timeouts, canceling tasks when they are no longer required.

#### Key Features

- **Context Values:** Allows us to store values that can be accessed across different parts of your program. It is useful for passing the data.
- **Cancelation Signals:** Provides a way to cancel tasks when they are no longer needed. This helps prevent tasks from running indefinitely and wasting resources.
- **Deadlines:** We can set deadlines for tasks using context.WithDeadline()
