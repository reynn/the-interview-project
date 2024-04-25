# The Interview Project

Technical interviews tend to be hand-wavy exercises focused on toy applications or over emphasize on lofty concepts and techniques. This project provides realistic scenarios for assessing a candidate's fundamental engineering skills: remote/asynchronous collaboration, code review, version control, and testing, to name a few.

Refer to the [contribution guide](/CONTRIBUTING.md) to get started.

## Architecture

Currently, the project consists of a Go microservice and a corresponding client application. They communicate using [gRPC](https://grpc.io/).

```mermaid
sequenceDiagram
    client->>authservice: send user/pass to authenticate
    activate authservice
    authservice-->>client: token or error
    client->>interviewservice: send request with retrieved token
    activate interviewservice
    interviewservice->>authservice: validate client provided jwt
    activate authservice
    authservice-->>interviewservice: response or error
    interviewservice-->>client: greeting or error
```

### Potential next steps

| Feature/Component      | Description                                                                                                  |
|------------------------|--------------------------------------------------------------------------------------------------------------|
| Authentication         | Authenticate and authorize API requests.                                                                     |
| Client SDK             | Manage client-side authentication.                                                                           |
| Configuration          | Inject config values from environment.                                                                       |
| Structured Logging     | -                                                                                                            |
| Testing                | Unit tests, integration tests.                                                                               |
| Continuous Integration | GitHub Actions for tests.                                                                                    |
| Front-end client       | A visual interface.                                                                                          |
| Setup scripts          | Automate onboarding.                                                                                         |
| Containerization       | Caution: Local Docker setup can be tricky so this should be a nice-to-have, not a requirement to contribute. |
