# Project stereo-server

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### CLI

Run:

```bash
make run-cli ARGS="process -f {{ path to json file }}"
```

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the server application
```bash
make run
```

Run the cli application
```bash
make run-cli
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
