#!/bin/sh
for f in $(ls *.go); do 
    echo
    echo == $f ==
    ./${f%.*}
done
