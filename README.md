# TODOアプリケーション

Udemyコース【Go入門】Golang基礎入門 + 各種ライブラリ + 簡単なTodoWebアプリケーション開発(Go言語)の一環として作成したシンプルなTODOアプリケーションです。詳しくは、[こちらのコースページ](https://www.udemy.com/course/golang-webgosql)をご覧ください。

## 機能

- タスクの追加、表示、削除
- ログイン機能
- 標準Goライブラリを使用

## 使用方法

1. リポジトリをクローンします:
    ```bash
    git clone https://github.com/kkato/todo_app.git
    cd todo_app
    ```

2. 依存関係をインストールします:
    ```bash
    go mod tidy
    ```

3. アプリケーションを実行します:
    ```bash
    go run main.go
    ```

4. ブラウザで以下のURLにアクセスします:
    ```
    http://localhost:8080
    ```

## 使用ライブラリ

- **net/http**: Webサーバの作成に使用
- その他、Go標準ライブラリ
