Me playing with `net/rpc` package in Go

## app

main app, connects to both micro services

## math-service

service to sum ints

## time-service

service get current time

## Usage

Start `math-service`, `time-service` and `app` separately with `go run`.

```bash
go run math-service
```

```bash
go run time-service
```

```bash
go run app
```
