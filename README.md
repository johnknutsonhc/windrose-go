# windrose-go

This runs an HTTP server that responds to requests with an SVG Windrose.

## Build & Run

Build the program:
```
go build .
```

Run the program:
```
./windrose-go
```

Alternatively, you can do it all at once:
```
go run .
```

## Usage

The server listens on port 8090. It will respond to the path `/windrose` and accepts a GET parameter, `angle`. For example:

```
curl http://localhost:8090/windrose?angle=270
```
