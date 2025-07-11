# Agent Instructions

To compile and test this project, run the setup script once to install system dependencies:

```sh
./scripts/setup.sh
```

Then run the following checks before committing any changes:

```sh
go vet ./...
go build ./...
```

The project has no tests, but building should succeed.
