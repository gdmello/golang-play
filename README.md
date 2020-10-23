# Templates
```
$ export CGO_ENABLED=0
$ go get "github.com/Masterminds/sprig"
$ go run templates.go
```

# Flags
## Build
```
$ go build flags.go
```

## Run
```
$ ./flags -name Fred -age 30
name:  Fred
age:  21
```

