# README

## About

This is the official Wails Svelte-TS template.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

**Note for macOS with ObjectBox:** If you encounter library loading errors, use the provided wrapper scripts:

- `./dev.sh` - for development mode
- `./run.sh` - for other wails commands

Alternatively, set the environment variable manually:

```bash
export DYLD_LIBRARY_PATH="$(pwd)/objectboxlib/lib:$DYLD_LIBRARY_PATH"
```

## Building

To build a redistributable, production mode package, use `wails build`.

**Note for macOS with ObjectBox:** Make sure to set `DYLD_LIBRARY_PATH` before building, or use the wrapper script.
