# About This

「4-1　Amazon Kinesisを使ってTwittterのデータを受け取る」をGolangで実装したもの。

Serverless Frameworkは使用せず、CloudFormationにてKinesis、EC2とアプリのデプロイを行う。

# Quick Start

1. CloudFormation

cfn.templateをテンプレートとしてCloudFormationを実行する。

# アプリについて

TwitterのAccessToken等はCloudFormationのパラメータに記述する。
パラメータに記述した値はEC2の環境変数に登録し、アプリは環境変数から値を取得する。
