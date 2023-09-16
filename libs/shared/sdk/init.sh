#!/bin/bash


# Remove go.mod and go.sum files
rm "openapi/go.mod" "openapi/go.sum"
rm -r "openapi/test"
echo "go.mod and go.sum removed successfully in openapi."
