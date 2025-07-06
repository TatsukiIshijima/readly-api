# ガイドライン

# アプリケーション概要

readly-apiは読書履歴を管理するアプリケーションのためのAPIです。ユーザーは読書履歴を登録、更新、削除することができます。

# 使用言語・ライブラリ

Go

| ライブラリ                                                          | 用途                     |
|----------------------------------------------------------------|------------------------|
| [jwt-go](https://github.com/golang-jwt/jwt)                    | JSON Web Token         |
| [uuid](https://github.com/google/uuid)                         | UUID                   |
| [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) | gPRCからJSONへのプロキシジェネレータ |
| [pq](https://github.com/lib/pq)                                | postgresドライバー          |
| [paseto](https://github.com/o1egl/paseto)                      | PASETO Token           |
| [viper](https://github.com/spf13/viper)                        | 設定ファイル読み込みなど           |
| [crypt](https://github.com/golang/crypto)                      | 暗号化                    |
| [go-genproto](https://github.com/googleapis/go-genproto)       |                        |
| [grpc-go](https://github.com/grpc/grpc-go)                     | gRPC                   |
| [protobuf](https://github.com/protocolbuffers/protobuf)        | Protocol Buffers       |

Go、及び各種ライブラリのバージョンについては `go.mod` を参照してください。

# ディレクトリ構成
readly-apiのディレクトリ構成は [katzien/app-structure-examples](https://github.com/katzien/app-structure-examples)を参考にしています。
原則インデントレベルを2としています。

```
.
├── cmd/
│   └── main.go              // アプリケーションのエントリーポイント
├── db/                      // DB操作関連/
│   ├── migration/           // マイグレーション操作のSQLクエリ
│   ├── query/               // CRUD操作のSQLクエリ
│   ├── sqlc/                // sqlcによりSQLクエリから自動生成されたGoコード
│   └── transaction/         // トランザクション管理 
├── configs/                 // 設定関連
│   └── env/                 // 環境変数
├── features/                // 機能ごとのディレクトリ
│   ├── book/                // 書籍管理機能
│   │   ├── domain/          // ドメインモデル
│   │   ├── repository/      // リポジトリ。DBやAPIへのアクセスの隠蔽。及びデータのキャッシュなどを行う
│   │   ├── server/          // HTTP/gRPCサーバーのエントリーポイントとハンドラ実装
│   │   └── usecase/         // ロジック・主要な処理を行う
│   ├── image/               // 画像アップロード機能
│   ├── user/                // ユーザ管理機能
├── middleware/              // ミドルウェア。各機能で共通で使用するものもここに定義
│   ├── auth/                // 認証ミドルウェア
│   └── image/               // 画像バリデーションミドルウェア
├── pb/                      // protoc-gen-goによりprotoファイルから自動生成されたGoコード/
│   └── readly/
│       └── v1/
├── proto/                   // protoファイル/
│   └── readly/
│       └── v1/
├── testdata/                // テストデータ
├── tools/                   // ツール
└── util/                    // 非常に小さい単位の共通処理
```

# アーキテクチャ

Clean Architectureを採用しています。

## db

- DB操作を行う場所です。SQLクエリは `db/query` に記述し、`db/sqlc` により自動生成されたコードを使用します。

## domain

- ドメインモデルを定義する場所です。protoへの変換などもここに記述します。

## repository

- リポジトリは、DBやAPIへのアクセスを隠蔽し、データのキャッシュなどを行います。

## server

- gRPCサーバーの実装を行います。gRPCのサービス定義は `pb/readly/v1` にあります。
- 主にusecaseを呼び出し、レスポンスを返します。

## usecase

- ロジックを定義する場所です。

# エラーハンドリング

- repositoryで発生したエラーは基本そのまま呼び出し元に返してください。呼び出し元でエラーハンドリングしてください。
    - 例：ある repository の関数を usecase で呼び出した場合は、usecase 側でエラーハンドリング
- repository内で独自エラーが必要な場合は `repository/error.go` に追記してください。
- usecaseで発生したエラーは、`usecase/error.go` に定義されたエラーを使用してください。
-usecase 内で独自エラーが必要な場合は `repository/error.go` に新たなエラーコードを追記してください。
+usecase 内で独自エラーが必要な場合は `usecase/error.go` に新たなエラーコードを追記してください。
- serverで発生したエラーは、`server/error.go` に定義された `gRPCStatusError` を用いて gRPCのステータスコードを返してください。

# テスト
- 同じ機能(Feature)内のテストはFakeを推奨します。他 Feature の機能が必要なテストの場合、その機能はMockを使用してください。
- ServerのテストはIntegration Testとして本物のDBを使用してください。