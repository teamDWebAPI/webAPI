# DogAPI
犬の画像検索webAPI
## 概要
![preview](./imagse/2024-03-11%2010.29.59%20localhost%20919dd51f5282.png)

## 説明

このAPIは、犬の画像検索WebAPIです。

犬の名前を選択することでランダムに該当する犬の画像を取得することができます。

## 使用技術

<table>
<tr>
  <th>カテゴリ</th>
  <th>技術スタック</th>
</tr>
<tr>
  <td rowspan=3>フロントエンド</td>
  <td>HTML</td>
</tr>
<tr>
  <td>CSS</td>
</tr>
<tr>
  <td>JavaScript</td>
</tr>
<tr>
  <td rowspan=2>バックエンド</td>
  <td>Go</td>
</tr>
</table>

## 使用方法
### クローン
このプロジェクトをあなたのPCで実行するために、クローンします。

下記手順でクローンしてください。

1. リポジトリをクローンする
```
git clone https://github.com/teamDWebAPI/webAPI.git
```

1. クローンしたリポジトリへ移動する
```
cd webAPI
```
### ローカルでサーバーを起動する
```
go run main.go
```
## エンドポイント
1. 犬のリスト表示`/api/list?filter=a&sort=ascend`
2. 犬の詳細表示`/api/item/{id} or /api/item/{breed}`
3. 画像の取得`/api/images?breed=hound&sub-breed=afghan?count=3`

## 今後の課題
### 本番環境へのデプロイ
- ローカル環境でしか動かないので、AWSのEC2を使って本番環境へデプロイする

### ログイン機能の作成
- データベースを導入してJWTを使って認証機能を作成する。

## チームメンバー
- [COCO](https://github.com/Taiga2022)
- [Blue](https://github.com/S-Taichiii)
