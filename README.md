# Pkl Go Bananas

This project is a great place to learn how Pkl might work end-to-end.


## Generate Go Lang binding, run this in the root of the app.

```shell
pkl-gen-go --project-dir $(pwd)/pkl pkl/*
```

## Release Pkl Configuration

```shell
pushd pkl/
rm -rf .out
VERSION=0.0.1 pkl project package
popd
```
Then upload `./pkl/.out` to a Github Release
