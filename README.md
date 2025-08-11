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

- **Window management** – draggable, resizable windows with optional pinning and fixed aspect ratios.
- **Main portal windows** – draw before other UI, black out the rest of the
  screen and leave the window background transparent for game content.
- **Flow layouts** – vertical or horizontal flows for composing widgets.
- **Overlay items** – keep controls always on screen.
- **Palette and style themes** – JSON files define colors and spacing. Switch
  them at runtime or reload automatically while iterating.
- **UI scaling** – call `eui.SetUIScale()` to adapt to any resolution or
  `eui.UIScale()` to read the current value. The scale is clamped between 0.5
  and 2.5. Windows using `AutoSize` adjust their dimensions automatically when
  the scale changes.
- **Image caching** – widgets cache their drawing for better performance.
  Enable `eui.DumpMode` to write the cached images to disk for inspection.
- **Tree dump** – press <kbd>Shift</kbd>+<kbd>`</kbd> or run the demo with
  `-tree` to save the window hierarchy as JSON.
- **Event system** – each widget returns an `EventHandler` that uses channels or
  callbacks so your code can react to clicks, slider movements and text input.
- **Debug overlays** – toggle with the `-debug` flag when running the demo.
- **Cross platform** – runs anywhere Ebiten does: desktop, web or mobile.
- **Basic touch support** – with two‑finger scrolling (drag up to scroll up).
  Mouse scrolling is clamped to +/-3 and rate-limited to 4 events per half-second on WebAssembly.
- **Image labels** – buttons, sliders, checkboxes, radios and dropdowns can combine image and text labels, with the image drawn before the text and optional custom sizing.
- **Vertical sliders** – sliders can be oriented vertically.
- **Logarithmic sliders** – sliders can map values on a logarithmic scale.
- **Hidden inputs** – text fields can mask their contents and reveal them while the eye icon is pressed.
- **Tooltips** – optional text hints appear when hovering over any item except flows.


## Running the Demo

The demonstration application lives under `cmd/demo`. You can run it directly using `go run` or build a binary:

```sh
go run ./cmd/demo             # launches the showcase window
# or
go build -o demo ./cmd/demo
./demo -debug                 # optional debug overlays and disposal logs
# dump cached images to ./debug after one frame then exit
./demo -dump
# dump the window tree to ./debug/tree.json then exit
./demo -tree
# pass -debug with go run to enable overlays
go run ./cmd/demo -debug
```

The demo now calculates the initial window size using the monitor's device
scale factor so the interface appears the same on standard and high‑DPI
displays.

## Building for WebAssembly

Compile the demo to `wasm` using the provided script. It passes size
optimization flags to `go build` and also creates a Brotli compressed
`demo.wasm.br` for the browser:

```sh
./scripts/build_wasm.sh
```

Open `web/index.html` in a browser to run it.

### WCAG adjustments

The `scripts/wcag_adjust.py` helper rewrites all palette and style files so text
and accents meet the WCAG 2 contrast recommendations. Run it whenever you tweak
the theme JSON files:

```sh
python3 scripts/wcag_adjust.py
```

## Customization

The library loads the built in `AccentDark` palette, `RoundHybrid` style and a default font automatically. Additional examples live under [`eui/themes`](eui/themes). Use `eui.ListThemes()` and `eui.ListStyles()` to see the names that are available. To try a different look enable `eui.AutoReload` and load files explicitly:

```go
// var ttf []byte
// _ = eui.EnsureFontSource(ttf)
// _ = eui.LoadTheme("MyTheme")
// _ = eui.LoadStyle("MyStyle")
```
See [themes/README.md](eui/themes/README.md) for a list of the bundled schemes and details on creating your own.

## Project Layout

- `eui` – reusable library code containing windows, flows and widgets
- `cmd/demo` – standalone example program wiring the library together

For a generated listing of all library functions see the [API reference](api.md).

## Testing

Some tests depend on a build tag. Run them with:

```sh
go test -tags test ./...
```
