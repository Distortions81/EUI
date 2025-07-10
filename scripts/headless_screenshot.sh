#!/bin/bash
set -e

# Ensure system dependencies are installed
./scripts/setup.sh
sudo apt-get install -y xvfb xdotool

# Download Go module dependencies
go mod download

# Start virtual frame buffer
Xvfb :99 -screen 0 1024x768x24 &
XVFB_PID=$!
sleep 2

# Run the demo with a screenshot key
DISPLAY=:99 EBITENGINE_SCREENSHOT_KEY=q go run . &
APP_PID=$!

# Give the window time to initialize
sleep 5

# Trigger screenshot
xdotool search --name "EUI Prototype" key q

# Wait a moment for the image to be written
sleep 2

# Clean up
kill $APP_PID || true
wait $APP_PID 2>/dev/null || true
kill $XVFB_PID || true
wait $XVFB_PID 2>/dev/null || true

echo "Screenshot saved as screenshot_<timestamp>.png"
