# The Interview Project

Technical interviews tend to be hand-wavy exercises focused on toy applications or over emphasize on lofty concepts and techniques. This project provides realistic scenarios for assessing a candidate's fundamental engineering skills: remote/asynchronous collaboration, code review, version control, and testing, to name a few.

This monorepo is an open-ended project that candidates of all skill/career levels can build upon to demonstrate their technical knowledge and abilities. Since the project has no specific functional objectives beyond being the foundation for interviews - build anything you can dream of!

## Architecture

Currently, the project consists of a Go microservice and a corresponding client application communicating over [gRPC](https://grpc.io/).

```
-------------------            -------------------
| Client          |            | Server          |
| --------        |            |                 |
| | SDK* |        | <--gRPC--> |                 |
| --------        |            |                 |
-------------------            -------------------

* Unimplemented
```

### Potential next steps

| Feature        | Description                            |
| -------------- | -------------------------------------- |
| Authentication | Authenticate API requests.             |
| Client SDK     | Manage client-side authentication.     |
| Configuration  | Inject config values from environment. |
| Logging        | gRPC middleware, log levels.           |

## Contribute

### Requirements

- [Go 1.18+](https://go.dev/doc/install)
- [gRPC prerequisites](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

### Setup

Run the setup script to generate protobuf stubs.

```sh
./setup.sh
```

### Workflow

Practice [Scaled Trunk-Based Development](https://trunkbaseddevelopment.com/#scaled-trunk-based-development) to reduce long-running branches in favor of **short-lived feature branches**.

The trunk is called `main`.

![scaled trunk-based development workflow](https://trunkbaseddevelopment.com/trunk1c.png)

1. Create a branch with a succinct name describing the intention for that piece of work (e.g. `dark-mode-toggle`<sup>1</sup>). Please do not contribute directly to `main`.

2. Create a [Pull Request](https://docs.github.com/en/pull-requests) (PR) when you are ready for feedback. Include a reasonably detailed description of the changes to help reviewers contextualize their feedback.

3. Once approved, use the [Squash and rebase](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/configuring-pull-request-merges/about-merge-methods-on-github#squashing-your-merge-commits)<sup>2</sup> feature to commit your work to `main`.

## Debug

### Inspect service with [grpcurl](https://github.com/fullstorydev/grpcurl)

```sh
# describe the service API
grpcurl -plaintext localhost:8080 describe InterviewService

# make an API call
grpcurl -plaintext -d '{"Name": "Obi-Wan"}' localhost:8080 InterviewService/HelloWorld
```

### Go runtime error: _package your_pkg is not in GOROOT_

A Go [workspace](https://go.dev/ref/mod#workspaces) is a collection of [modules](https://go.dev/ref/mod#modules-overview) defined by the `go.work` file. For a module to be properly indexed by `gopls` or the go compiler, it must be referenced in the workspace file.

```
use ./your-module
```

---

## Appendix

### 1. Style guide

#### Branch names

Prefer `kebab-case` for branch names.

#### Linting and formatting

- [gopls](https://pkg.go.dev/golang.org/x/tools/gopls)

### 2. Why "Squash and rebase"?

Squashing commits into a single commit keeps the main code history focused on significant chunks of contributions while preserving as much detail about commits as desired by the author.

Using the rebase strategy to fold changes into the main code history (instead of using merge) maintains a chronological history best suited for easy code review and reversion when necessary.

As a rule of thumb, PR authors should manage the squashing and rebasing of their own branches. However, since this project is an interview format, it is likely that you will leave a PR up and not return to the repo. In this case, the interviewer should complete the PR on the author's behalf.
