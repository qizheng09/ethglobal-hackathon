#!/usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)

export LOG_FIR="$CURDIR/log"

if [ ! -d "$CURDIR/log" ]; then
  mkdir -p "$CURDIR/log"
fi

exec nohup "$CURDIR/bin/hackathon" &