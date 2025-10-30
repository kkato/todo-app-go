# go-todo-app

This is a simple TODO application created as part of the Udemy course【Go入門】Golang基礎入門 + 各種ライブラリ + 簡単なTodoWebアプリケーション開発(Go言語). For more details, please refer to the [course page](https://www.udemy.com/course/golang-webgosql).

## Features

- Add, view, and delete tasks
- Login functionality
- Uses standard Go libraries

## How to Use

1. Clone the repository:
    ```bash
    git clone https://github.com/kkato/todo_app.git
    cd todo_app
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

4. Access the following URL in your browser:
    ```
    http://localhost:8080
    ```

## Libraries Used

- **net/http**: Used to create the web server
- Other standard Go libraries
