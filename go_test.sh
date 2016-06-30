#!/usr/bin/env bash

PACKAGES=$(find ./ -type d -not -path '*/\.*' -not -path "*/\webapp*")
echo $PACKAGES
for pkg in $PACKAGES;do
if go test -coverprofile=coverage-one.out -covermode=atomic $pkg; then
tail -n +2 coverage-one.out >> coverage.txt
fi
done