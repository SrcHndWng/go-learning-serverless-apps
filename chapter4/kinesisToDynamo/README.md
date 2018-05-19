# About This

「4-2　AWS Lambdaを使ってストリーミングデータをAmazon DynanoDBへ保存する」をGolangで実装したもの。

# Serverless Template for Golang

This repository contains template for creating serverless services written in Golang.

## Usage

1. Compile function

```
cd kinesisToDynamo
GOOS=linux go build -o bin/main
```

2. Deploy!

```
serverless deploy --account 9999999999(Your AWS Account Number)
```
