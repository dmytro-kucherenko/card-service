# Card Service

## About

This is a repository for bank card verification service.

## Steps to run

### One-command

Build and run docker container:
```shell
make compose-up
```

### Development

Download dependencies:
```shell
go mod download
```

(Optional) Create and fill `.env` file using `.env.example` as a reference.

Run application:
```shell
make run
```

Or run with hot reload:
```shell
make watch
```

Run automated tests:
```shell
make test
```

Check `Makefile` file for more commands.

## Application

### RESTful API

Routes:

- Card Validation: /api/card/validate
- OpenAPI Specification: /swagger/index.html

### gRPC

Procedures:

- Card Validation: /card.Service/Validate

### Error Codes

Global:

- `001` - Internal server error.
- `002` - Validation error.

Card:

- `101` - Number is invalid.
- `102` - Number check digit is invalid.
- `103` - Date is expired.

## Instruments

### Postman

Use public [workspace](https://www.postman.com/solar-star-295145/workspace/card-service) either directly, as a reference or fork its collections.
