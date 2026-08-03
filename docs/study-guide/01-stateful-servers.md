# Block 1 - Stateful Servers (4:04-5:26)

## Starting point

A single server that stores data inside itself. The guide illustrates this with a simple structure:

```go
type Server struct {
    files map[string]string // ID -> file location
}
```

## Problems this reveals

| Problem | Consequence |
|---|---|
| Data coupled to the server | If the server dies, the data is lost |
| Cannot scale | Each new instance would have different data |

## Rule to remember (from the source)

Data coupled to the machine means no scalability and no resilience.
