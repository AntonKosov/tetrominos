#!/bin/bash

go build -gcflags="all=-N -l" -o ./build/tetrominos.debug

dlv --listen=:2345 --headless=true --api-version=2 exec ./build/tetrominos.debug