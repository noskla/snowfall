#!/bin/bash
if [ ! go version &> /dev/null ]; then
    printf "Go has been not found. Install golang using your package manager or by"
    printf "following the instructions for your OS at https://golang.org/\n"
    exit 1
fi

if [ ! -d "dist/" ]; then
    mkdir dist
else
    rm -r dist/*
fi

printf "Compiling... "
export GOARCH="amd64"
for os in 'windows' 'linux' 'freebsd'; do
    filepath="dist/snowfall.$os"
    if [ "$os" = "windows" ]; then
        filepath="$filepath.exe"
    fi
    GOOS="$os" go build -o "$filepath"
    if [ $? -eq 0 ]; then
        printf "$os ok,"
    else
        printf "$os failed,"
    fi
done
printf "Done\n"

printf "Copying assets..."
cp -r templates dist/
cp -r static dist/
printf "Done\n"

printf "Compiling SASS..."
if [ ! -d "dist/static/css" ]; then
    mkdir dist/static/css
fi
for file in sass/*.sass; do
    outfile=$(basename "$file" .sass)
    sass --no-source-map --style compressed "$file" "dist/static/css/$outfile.css"
done
printf "OK\n"

printf "Binaries are ready in \"dist\" directory.\n"
