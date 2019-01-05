#!/bin/bash
for dir in */; do
  cd $dir
  echo Runing ${dir}main.go
  time go run main.go
  cd ..
  echo
done
