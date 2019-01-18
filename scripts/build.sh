#!/bin/sh

scriptsdir=$(dirname "$0")

docker build -t torbensky/golang-playground-postgres $scriptsdir/../