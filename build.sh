#!/bin/sh
set -ux
for f in $(ls *.go); do 
    # build a static executable
    CGO_ENABLED=0 go build -a -ldflags -s $f
    # pack it small, pack it tight (--ultra-brute might be tighter, but not necessarily)
    upx -9 --brute ${f%.*}
done
