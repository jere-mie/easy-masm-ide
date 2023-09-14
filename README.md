# easy-masm-ide

A simple online IDE for MASM32 programs, built with Go

## Setup

To run easy-masm-ide locally, you'll need:

- A modern version of [Go](https://go.dev/)
- If you're a Linux user, you'll also need [Wine](https://www.winehq.org/)

To run the application, simply run the command:

```sh
go run main.go
```

And the server will start on port [8080](localhost:8080).

### Running on a Different Port

To change the port, simply modify the `EASY_MASM_IDE_PORT` environment variable.

## Future Plans

- [x] Add Windows support
- [ ] Containerize with Docker
- [ ] Add support for specifying input

## Known Bugs

- MASM source files on Windows are not deleted after being run
