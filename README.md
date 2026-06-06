# Hibiki

Here's a demonstration: https://x.com/_vector15/status/2063297998654984465?s=20

Hibiki is a high-performance, synchronized ASCII video and audio playback engine designed for the terminal. Written in Go, it leverages `ffmpeg` for raw byte-stream video decoding and `mpv` for headless, Wayland-compatible audio execution.

## Features

* **Zero-Allocation Rendering:** Pre-allocates memory buffers mapped exactly to terminal dimensions, eliminating garbage collection overhead during playback.
* **Raw Byte Pipeline:** Bypasses standard image decoding libraries by piping uncompressed grayscale bytes directly from `ffmpeg`, preventing frame drops and header parsing panics.
* **Strict Frame Synchronization:** Utilizes a precise 30 FPS internal ticker to maintain perfect alignment between terminal visual output and the detached audio process.
* **Flicker-Free Display:** Implements ANSI escape sequences for cursor repositioning rather than full terminal clearing, ensuring smooth optical persistence.

## Prerequisites

Ensure the following system dependencies are installed and available in your system path:

* **Go**
* **FFmpeg**
* **mpv**

For installation via pacman:

```bash
sudo pacman -S go ffmpeg mpv
```

## Project Structure

The architecture is strictly modularized to isolate distinct execution domains:

* `video/`: Manages the `ffmpeg` subprocess, configuring the resolution and streaming raw grayscale byte chunks into memory.
* `render/`: Translates raw luminance byte values (0-255) into mapped ASCII characters within a pre-allocated byte slice.
* `audio/`: Manages the detached `mpv` subprocess for asynchronous background audio.
* `engine/`: The central orchestrator handling terminal state management, OS interrupt signals, and the strict blocking 30 FPS rendering loop.
* `cmd/example/`: The standard entry point for CLI execution.

## Usage

1. Initialize the project and ensure all Go files are properly saved.
2. Place your target video file (e.g., `bad_apple.mp4`) in the root directory. Update the file reference in `cmd/example/main.go` if utilizing a different media file.
3. Maximize your terminal for optimal resolution mapping.
4. Execute the engine from the project root:

```bash
go run cmd/example/main.go
```

The application will automatically suppress terminal cursor visibility during execution and safely restore terminal state upon standard completion or an explicit interrupt signal (Ctrl+C).
