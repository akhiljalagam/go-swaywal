#!/bin/bash
set -x
env -i bash -c "HOME=$HOME \
	SWAYSOCK=/run/user/1000/sway-ipc.1000.$(pgrep -x sway).sock \
  UNSPLASH_ACCESS_KEY=$UNSPLASH_ACCESS_KEY \
  SWAY_WIDTH=3840 \
  SWAY_HEIGHT=2160 \
  UNSPLASH_SEARCH=nature \
  go run main.go"
