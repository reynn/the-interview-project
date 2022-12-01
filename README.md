# The Interview Repo

Technical interviews tend to either be a hand-wavy exercise of toy applications or are over-focused on lofty concepts and techniques. This project provides realistic scenarios for assessing a candidate's fundamental engineering skills: remote/asynchronous collaboration, code review, version control, testing, to name a few.

The monorepo serves as a long-lived project that candidates of all skill/career levels can build upon to demonstrate their technical knowledge and abilities. Since the project has no specific functional objectives beyond being the foundation for the interview - build anything you can dream of!

## Architecture

Currently, the project consists of a gRPC microservice implemented Go. Next steps should be to implement a client application to consume the service's API, perhaps with an SDK layer in-between.

```
---------------------            ---------------------
| Client (Go)       |            | Server (Go)       |
| -------           |            |                   |
| | SDK |           | <--gRPC--> |                   |
| -------           |            |                   |
|                   |            |                   |
---------------------            ---------------------
```

## Setup

### Requirements

- [Go 1.18+](https://go.dev/doc/install)
- [gRPC prerequisites](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

Run the setup script to generate protobuf stubs.

```sh
./setup.sh
```
