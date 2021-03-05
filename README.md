# View_Pictures
良い景色の写真を撮って共有できるアプリです。

## URL
https://pictures-app.work/

## 使用技術
### フロントエンド
- Vue.js: 2.6.12
- Vue CLI: 3.5.0
- Buefy

### バックエンド
- Go: 1.12
- Echo: 3.3.10

### データベース
- MySQL: 8.0

## イメージ図
<img width="1090" alt="View Pictures" src="https://user-images.githubusercontent.com/44848412/110053689-3657e700-7d9d-11eb-874e-f5b17a307714.png">


## クラウドアーキテクチャ
![aws_architecture_diagram](https://user-images.githubusercontent.com/44848412/110054109-db72bf80-7d9d-11eb-89b8-be2e80e4489e.png)

### AWS
- ECS(Fargate)/ECR/RDS(MySQL)/ELB/CloudWatach/S3/CloudFront/Route53/ACMを使用
- CloudFormationでコード化
- ACMからSSL証明書を発行しHTTPS化

## ER図

![ER図](https://github.com/Junpei-Nakasone/pictures_app/blob/main/pictures_ERdiagram.png?raw=true)

