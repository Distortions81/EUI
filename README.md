## A Minimal vector-based Ebitenengine retained-mode UI library
### This library is nowhere near complete (PRE-ALPHA), and breaking changes will come very often.
*For quicker development, the code is not composed as a library at this time*  
Drawn using vectors and supports floating-point UI scaling.

![screenshot](https://raw.githubusercontent.com/Distortions81/EUI/refs/heads/main/Screenshot.png)

An additional showcase window demonstrating all widget types can be found in
`showcase.go`.

# Flow and item pinning:
Flows and items can have pinning type set. The default is PIN_TOP_LEFT.
The pin changes the point of reference for the item/flow position.

Other options are:  
PIN_TOP_LEFT  
PIN_TOP_CENTER  
PIN_TOP_RIGHT

PIN_MID_LEFT  
PIN_MID_CENTER  
PIN_MID_RIGHT

PIN_BOTTOM_LEFT  
PIN_BOTTOM_CENTER  
PIN_BOTTOM_RIGHT

## Rendering a screenshot on a headless system
Run `scripts/headless_screenshot.sh` to install dependencies and launch the demo under `Xvfb`. The script sends a key press to trigger Ebitengine's built-in screenshot function. The resulting PNG is saved in the current directory with a name such as `screenshot_YYYYMMDDHHMMSS.png`.

## Themes
Theme settings are defined as JSON files inside the `themes` directory. The file `DefaultDark.json` reflects the built-in style and is loaded at startup. Theme files contain widget colors, margins and other styling values which can be modified or replaced to create custom appearances.
