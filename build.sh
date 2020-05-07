#!/usr/bin/env bash

echo "Start build Radio Simulator...."
rm -rf bin/simulator
go build -o bin/simulator -x src/simulator.go
