# EUI

EUI is a minimal retained‑mode UI built with the [Ebiten](https://ebiten.org/) game engine.
It provides basic window management, flows and widgets built entirely with vector graphics.
The project is currently in early development and APIs will likely change.

[Live demo here](https://m45sci.xyz/u/dist/eui/)

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
- **UI scaling** – call `eui.SetUIScale()` to adapt to any resolution or
  `eui.UIScale()` to read the current value. Windows using `AutoSize` adjust
  their dimensions automatically when the scale changes.
- **HiDPI support** – automatically adjusts the UI scale and screen resolution
  to match the device scale factor. Disable by setting `eui.AutoHiDPI = false`.
- **Image caching** – widgets cache their drawing for better performance.
  Enable `eui.DumpMode` to write the cached images to disk for inspection.
- **Event system** – each widget returns an `EventHandler` that uses channels or
  callbacks so your code can react to clicks, slider movements and text input.
- **Debug overlays** – toggle with the `-debug` flag when running the demo.
- **Cross platform** – runs anywhere Ebiten does: desktop, web or mobile.
- Touch controls are not yet implemented.

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

## Building for WebAssembly

Compile the demo to `wasm` using the provided script. It passes size
optimization flags to `go build` and also creates a Brotli compressed
`demo.wasm.br` for the browser:

```sh
./scripts/build_wasm.sh
```

Open `web/index.html` in a browser to run it.

## Customization

The library loads the built in `AccentDark` palette, `RoundHybrid` style and a default font automatically. Additional examples live under [`eui/themes`](eui/themes). Use `eui.ListThemes()` and `eui.ListStyles()` to see the names that are available. To try a different look enable `eui.AutoReload` and load files explicitly:

```go
// var ttf []byte
// _ = eui.EnsureFontSource(ttf)
// _ = eui.LoadTheme("MyTheme")
// _ = eui.LoadStyle("MyStyle")
```
HiDPI scaling is enabled by default so the UI keeps the same size on screen when switching between standard and high‑DPI displays. `eui.Layout` automatically scales the game screen to the device resolution. Set `eui.AutoHiDPI = false` to disable this behavior.

See [themes/README.md](eui/themes/README.md) for a list of the bundled schemes and details on creating your own.

## Project Layout

- `eui` – reusable library code containing windows, flows and widgets
- `cmd/demo` – standalone example program wiring the library together

For a generated listing of all library functions see the [API reference](api.md).
