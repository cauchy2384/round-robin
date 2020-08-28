# Simple round-robin
Thread safe round-robin example implementation using atomics

## Installation
```
go get github.com/cauchy2384/round-robin
```

## Testing
```
go test -race .
```


## Usage
```
import (
    roundrobin "github.com/cauchy2384/round-robin" 
)

list := []string{"Frodo", "Samwise", "Meriadoc", "Peregrin"}

rr, err := roundrobin.New(list)
// handle error

fmt.Println(rr.Next())  // Frodo
fmt.Println(rr.Next())  // Samwise
fmt.Println(rr.Next())  // Meriadoc
fmt.Println(rr.Next())  // Peregrin
fmt.Println(rr.Next())  // Frodo
```
