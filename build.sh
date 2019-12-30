#!/usr/bin/env bash
mkdir -p output/bin
curdir=`cd $(dirname $0); pwd -P`
echo $curdir

go build -o bootstrap .
mv bootstrap output/bin
chmod +x output/bin/bootstrap