#!/bin/sh

if pgrep -x "./core" > /dev/null
then
    exit 0
else
    exit 1
fi
