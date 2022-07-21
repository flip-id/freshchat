# Freshchat
[![Go Reference](https://pkg.go.dev/badge/github.com/flip-id/freshchat.svg)](https://pkg.go.dev/github.com/flip-id/freshchat)
[![Go Report Card](https://goreportcard.com/badge/github.com/flip-id/freshchat)](https://goreportcard.com/report/github.com/flip-id/freshchat)

Freshchat API client library using Go, https://developers.freshchat.com/api/#whatsapp.

# How to Test

To run the integration tests, we need to do the following:
1. Make a new copy of `.env.example` to `.env` by running this command below:
```bash
cat .env.example > .env
```
2. Fill the new .env with the parameter that we already prepared.
3. Run the tests by running this command:
```bash
go test -v -race -tags=integration -covermode=atomic ./...
```