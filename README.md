# EUI

EUI is a minimal retained‑mode UI built with the [Ebiten](https://ebiten.org/) game engine.
It provides basic window management, flows and widgets built entirely with vector graphics.
The project is currently in an early development and APIs will likely change.

## Currently the API is not complete!

![screenshot](https://raw.githubusercontent.com/Distortions81/EUI/refs/heads/main/Screenshot.png)

## Getting Started

The demo and tests rely on a few system packages in addition to Go. A helper script is included to install them. **The script is written for Linux distributions that use the `apt` package manager such as Ubuntu or Debian.** It performs an `apt-get` update followed by installation of the libraries required by Ebiten (`libxrandr-dev`, `libxinerama-dev`, `libxcursor-dev`, `libxi-dev`, `libxxf86vm-dev`, `libgl1-mesa-dev` and others). Run it once before building. If you are on a different operating system install equivalent packages using your distribution's package manager before continuing:

```sh
./scripts/setup.sh
```

After dependencies are installed you can vet and build the project with the standard Go tools. Running these commands ensures the library compiles correctly before you start experimenting:

```sh
go vet ./...
go build ./...
```

Format modified Go files before committing:

```sh
gofmt -w <files>
```

## Running the Demo

The demonstration application lives under `cmd/demo`. You can run it directly using `go run` or build a binary:

```sh
go run ./cmd/demo             # launches the showcase window
# or
go build -o demo ./cmd/demo
./demo -debug                 # optional debug overlays
# dump cached images to ./debug after one frame then exit
./demo -dump
# pass -debug with go run to enable overlays
go run ./cmd/demo -debug
```

## Customization

The library loads the built in `AccentDark` color theme, `RoundHybrid` layout and a default font automatically. Additional examples live under [`eui/themes`](eui/themes). Use `eui.ListThemes()` and `eui.ListLayouts()` to see the names that are available. To try a different look enable `eui.AutoReload` and load files explicitly:

```go
// var ttf []byte
// _ = eui.EnsureFontSource(ttf)
// _ = eui.LoadTheme("MyTheme")
// _ = eui.LoadLayout("MyLayout")
```

See [themes/README.md](eui/themes/README.md) for a list of the bundled schemes and details on creating your own.

## Project Layout

- `eui` – reusable library code containing windows, flows and widgets
- `cmd/demo` – standalone example program wiring the library together
- `scripts/setup.sh` – helper for installing build dependencies

For a generated listing of all library functions see the [API reference](api.md).
