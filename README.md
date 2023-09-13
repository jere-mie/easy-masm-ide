# easy-masm-ide

A simple online IDE for MASM32 programs, built with Go

## Setup

To run easy-masm-ide locally, you'll need:

- A Linux based system (Windows support will come soon)
- [Wine](https://www.winehq.org/)
- A modern version of [Go](https://go.dev/)

To run the application, simply run the command:

```sh
go run main.go
```

And the server will start on port 8080.

### Running on a Different Port

To change the port, simply modify the `EASY_MASM_IDE_PORT` environment variable.

## Future Plans

- [ ] Containerize with Docker
- [ ] Add Windows support
- [ ] Add support for specifying input
