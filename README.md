# todo-app-go

Udemyコース【Go入門】Golang基礎入門 + 各種ライブラリ + 簡単なTodoWebアプリケーション開発(Go言語)の一部として作成したシンプルなTODOアプリケーションです。詳細については、[コースページ](https://www.udemy.com/course/golang-webgosql)を参照してください。

## 機能

- タスクの追加、表示、削除
- ログイン機能
- Go標準ライブラリを使用

## 使い方

1. リポジトリをクローン:
    ```bash
    git clone https://github.com/kkato/todo_app.git
    cd todo_app
    ```

2. 依存関係をインストール:
    ```bash
    go mod tidy
    ```

3. アプリケーションを実行:
    ```bash
    go run main.go
    ```

4. ブラウザで以下のURLにアクセス:
    ```
    http://localhost:8080
    ```

## 使用ライブラリ

- **net/http**: Webサーバーの作成に使用
- その他のGo標準ライブラリ
