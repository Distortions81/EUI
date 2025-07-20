# Agent Instructions

Use the setup script once to install required system packages:

```sh
./scripts/setup.sh
```

Before committing new changes run the following checks from the repository root:

```sh
go vet ./...
go build ./...
```

Run `gofmt -w` on all Go source files you modify to keep formatting consistent.

Update any relevant documentation (README.md, api.md, etc.) when behavior,
commands or public APIs change so docs stay in sync with the code.

