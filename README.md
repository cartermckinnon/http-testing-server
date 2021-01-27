# HTTP Testing Server

HTTP server for testing/experimentation.
This server will respond with a configured response code, and by default will listen and respond to any path sent to the configured port.

## Build
```sh
go build
```

## Usage
```sh
testing-server - HTTP testing server which responds to requests as configured.

  Flags: 
       --version         Displays the program version string.
    -h --help            Displays help with available flag, subcommand, and positional value parameters.
    -x --path            HTTP path pattern on which to listen. Pattern follows the same rules as the one documented in golang's http.ServceMux. (default: /)
    -p --port            Port on which to listen. (default: 8080)
    -r --response-code   HTTP response code for all requests. (default: 200)
    -c --tls-cert-file   Certificate for TLS. Requires -k/--key to also be provided.
    -k --tls-key-file    Key for TLS. Requires -c/--cert to also be provided.
```

## TLS

To generate a self-signed certificate, use `tls.sh` and provide the certificate and key arguments when starting the server.
