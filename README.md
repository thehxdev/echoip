# EchoIP
A small program for cheching public IP address.


## Build
First you must install `go` compiler from [go.dev](https://go.dev) website.
Then build `echoip` executable:

```bash
go build -ldflags='-s -d' .
```

## Usage
Running `echoip` is very simple:

```bash
# default value for `-addr` flag is `localhost:5000`
./echoip -addr ':5000'
```

If you run `echoip` with command above, it will listen on `0.0.0.0:5000`.
For TLS support you can use a reverse proxy like [Caddy](https://caddyserver.com).
