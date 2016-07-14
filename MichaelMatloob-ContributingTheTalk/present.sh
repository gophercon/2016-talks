#!/bin/sh

cd $(dirname $0)
awk  'BEGIN{FS="[()]"} /^!/{print $2}' README.md \
    | xargs open /Applications/Preview.app


