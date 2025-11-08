#!/bin/bash

# Set the library path for ObjectBox
export DYLD_LIBRARY_PATH="$(cd "$(dirname "$0")" && pwd)/objectboxlib/lib:$DYLD_LIBRARY_PATH"

# Run wails dev with all arguments passed through
wails dev "$@"

