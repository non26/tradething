#!/bin/bash
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap cmd/app/aws_lambda/main.go
zip tradething.zip bootstrap