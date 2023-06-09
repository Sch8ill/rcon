#!/usr/bin/env bash

target="cmd/cli/rcon-cli.go"
output_prefix=$1

platforms=("windows/amd64" "windows/arm" "linux/amd64" "linux/arm64" "linux/arm")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$output_prefix'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $target
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done