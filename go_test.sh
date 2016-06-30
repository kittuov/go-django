#!/usr/bin/env bash

PACKAGES=$(find ./ -type d -not -path '*/\.*' -not -path "*/\webapp*")
echo $PACKAGES
for pkg in $PACKAGES;do
go test -coverprofile=coverage-one.out -covermode=count $pkg
tail -n +2 coverage-one.out >> coverage.out
done