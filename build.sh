#!/usr/bin/env bash

# Clean previous builds
rm -rf ./dist/*

# For MAC
GOOS=darwin GOARCH=amd64 go build -o ./dist/hack-stulish-mac ./stulish

# For Linux
GOOS=linux GOARCH=amd64 go build -o ./dist/hack-stulish-linux-x64 ./stulish
GOOS=linux GOARCH=386 go build -o ./dist/hack-stulish-linux-x86 ./stulish

# For Windows
GOOS=windows GOARCH=amd64 go build -o ./dist/hack-stulish-windows-x64.exe ./stulish
GOOS=windows GOARCH=386 go build -o ./dist/hack-stulish-windows-x86.exe ./stulish