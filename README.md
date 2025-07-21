# EUI

EUI is a minimal retained‑mode UI built with the [Ebiten](https://ebiten.org/) game engine.
It provides basic window management, flows and widgets built entirely with vector graphics.
The project is currently in early development and APIs will likely change.

## Currently the API is not complete!

![screenshot](https://raw.githubusercontent.com/Distortions81/EUI/refs/heads/main/Screenshot.png)

## Features

EUI aims to stay small and easy to use while still providing the essentials for
building a game UI. Highlights include:

- **Window management** – draggable, resizable windows with optional pinning.
- **Flow layouts** – vertical or horizontal flows for composing widgets.
- **Overlay items** – keep controls always on screen.
- **Palette and style themes** – JSON files define colors and spacing. Switch
  them at runtime or reload automatically while iterating.
- **UI scaling** – call `eui.SetUIScale()` to adapt to any resolution.
- **Image caching** – widgets cache their drawing for better performance.
  Enable `eui.DumpMode` to write the cached images to disk for inspection.
- **Event system** – each widget returns an `EventHandler` that uses channels or
  callbacks so your code can react to clicks, slider movements and text input.
- **Debug overlays** – toggle with the `-debug` flag when running the demo.
- **Cross platform** – runs anywhere Ebiten does: desktop, web or mobile.
- Touch controls are not yet implemented.

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
./demo -debug                 # optional debug overlays and disposal logs
# dump cached images to ./debug after one frame then exit
./demo -dump
# pass -debug with go run to enable overlays
go run ./cmd/demo -debug
```

## Customization

The library loads the built in `AccentDark` palette, `RoundHybrid` style and a default font automatically. Additional examples live under [`eui/themes`](eui/themes). Use `eui.ListThemes()` and `eui.ListLayouts()` to see the names that are available. To try a different look enable `eui.AutoReload` and load files explicitly:

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
