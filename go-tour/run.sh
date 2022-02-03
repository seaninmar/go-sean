#!/bin/sh
for f in $(ls *.go); do 
    e=${f%.*}
    echo
    echo == $e ==
    ./$e
done
