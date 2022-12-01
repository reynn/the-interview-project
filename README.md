# The Interview Project

Technical interviews tend to be a hand-wavy exercise of toy applications or are overly focused on lofty concepts and techniques. This project provides realistic scenarios for assessing a candidate's fundamental engineering skills: remote/asynchronous collaboration, code review, version control, testing, to name a few.

The monorepo is a long-lived project that candidates of all skill/career levels can build upon to demonstrate their technical knowledge and abilities. Since the project has no specific functional objectives beyond being the foundation for interviews - build anything you can dream of!

## Architecture

Currently, the project consists of a gRPC microservice implemented Go. Next steps should be to implement a client application to consume the service's API, perhaps with an SDK layer in-between.

```
------------------            ------------------
| Client         |            | Server         |
| -------        |            |                |
| | SDK |        | <--gRPC--> |                |
| -------        |            |                |
------------------            ------------------
```

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

3. Once approved, use the "Squash and rebase"<sup>2</sup> feature to commit your work to `main`.

## Debug

### Inspect service with [grpcurl](https://github.com/fullstorydev/grpcurl)

```sh
# describe the service API
grpcurl -plaintext localhost:8080 describe InterviewService

# make an API call
grpcurl -plaintext -d '{"Name": "Obi-Wan"}' localhost:8080 InterviewService/HelloWorld
```

## Appendix

### 1. Style guide

#### Branch names

Prefer `kebab-case` for branch names.

#### Linting and formatting

- [gopls]()

### 2. Why "Squash and rebase"?

Squashing commits into a single commit keeps the main code history focused on significant chunks of contributions while preserving as much detail about commits as desired by the author.

Using the rebase strategy to fold changes into the main code history (instead of using merge) maintains a chronological history best suited for easy code review and reversion when necessary.

As a rule of thumb, PR authors should manage the squashing and rebasing of their own branches. However, since this project is an interview format, it is likely that you will leave a PR up and not return to the repo. In this case, the interviewer should complete the PR on the author's behalf.
