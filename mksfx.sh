#!/bin/bash

files="freeze hold hold2 rot x"

for name in $files; do
	sox ${name}.aif -t raw -c 1 -r 44100 -b 8 -e unsigned-integer --norm ${name}.raw
done

go run samples.go $(for n in $files; do echo -n "${n}.raw "; done) > sfx.tal 

