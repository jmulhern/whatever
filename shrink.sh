#!/usr/bin/env bash

name=${1}
filename=$(basename -- "$name")
filename="${filename%.*}"
cwebp -q 50 public/${name} -o public/${filename}.webp