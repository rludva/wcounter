#!/bin/bash

# Load common dependencies for bash scripts..
MY_PATH=$(dirname "$0")

OUTPUT_DIRECTORY="$MY_PATH/.."
if [ ! -e "$OUTPUT_DIRECTORY" ]; then
  mkdir $OUTPUT_DIRECTORY
fi

function build() {
  GOOS="$1"
  GOARCH="$2"

  echo "Building for $GOOS/$GOARCH.."

  OUTPUT="$OUTPUT_DIRECTORY/counter_${GOOS}_${GOARCH}"
  if [ -e "$OUTPUT" ]; then
    rm "$OUTPUT"
  fi

  env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$OUTPUT"
  chmod +x $OUTPUT
}

build linux amd64
#build windows amd64
#build darwin amd64 
ls -l "$OUTPUT_DIRECTORY"
