# About This

「4-1　Amazon Kinesisを使ってTwittterのデータを受け取る」をGolangで実装したもの。

Serverless Frameworkは使用せず、CloudFormationにてKinesis、EC2とアプリのデプロイを行う。

# Quick Start

1. S3 Bucket

S3のBucketを用意し、ビルドして作成したバイナリを配置する。
Bucket名、バイナリ名はCloudFormation実行時に指定する。

バイナリは以下のコマンドでEC2で実行できる形式とすること。

```
$ GOOS=linux go build
```

2. CloudFormation

cfn.templateをテンプレートとしてCloudFormationを実行する。

# アプリについて

TwitterのAccessToken等はCloudFormationのパラメータに記述する。
パラメータに記述した値はEC2の環境変数に登録し、アプリは環境変数から値を取得する。
