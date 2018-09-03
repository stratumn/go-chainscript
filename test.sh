#!/usr/bin/env bash

mkdir -p test_samples

# Make sure images are up-to-date:
docker pull stratumn/go-chainscript:latest
docker pull stratumn/js-chainscript:latest

# Mount a directory for reading/writing samples.
MOUNT=type=bind,source="$(pwd)"/test_samples,target=/samples

# Go container details.
GO_CONTAINER=stratumn/go-chainscript:latest
GO_CONTAINER_SAMPLE_FILE=/samples/go-samples.json

# Javascript container details.
JS_CONTAINER=stratumn/js-chainscript:latest
JS_CONTAINER_SAMPLE_FILE=/samples/js-samples.json

# Sample files (outside of the Docker context):
GO_SAMPLE_FILE=./test_samples/go-samples.json
JS_SAMPLE_FILE=./test_samples/js-samples.json

# Generate test samples from all repositories:
docker run --mount $MOUNT $GO_CONTAINER generate $GO_CONTAINER_SAMPLE_FILE
if [ ! $? -eq 0 ]; then
    exit 1
fi

docker run --mount $MOUNT $JS_CONTAINER generate $JS_CONTAINER_SAMPLE_FILE
if [ ! $? -eq 0 ]; then
    exit 1
fi

# Verify files:
if [ ! -e $GO_SAMPLE_FILE ]; then
    echo "go-samples.json is missing"
    exit 1
fi

if [ ! -e $JS_SAMPLE_FILE ]; then
    echo "js-samples.json is missing"
    exit 1
fi

# Validate test samples:
echo "---------- Javascript ----------"
docker run --mount $MOUNT $JS_CONTAINER validate $GO_CONTAINER_SAMPLE_FILE
if [ ! $? -eq 0 ]; then
    exit 1
fi
echo "--------------------"

echo "---------- Golang ----------"
docker run --mount $MOUNT $GO_CONTAINER validate $JS_CONTAINER_SAMPLE_FILE
if [ ! $? -eq 0 ]; then
    exit 1
fi
echo "--------------------"
