# ToDo app

## Description

- Simple ToDo management app that was created based on 
[Build a Todo App in Golang, MongoDB, and React](https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6).

- This app is used [echo](https://github.com/labstack/echo) instead of [gorilla/mux](https://github.com/gorilla/mux).

## Usage

1. Make a clone
```
git clone https://github.com/seicode/goToDo.git
cd goToDo
```

2. Set up your mongoDB atlas and replace .env variables with what you have.

3. Build the server with go on localhost:1323.
```
go run server/main.go
```

4. Build the frontend with npm on localhost:3000.
```
cd client
npm install
npm start
```
Automatically, the app will show up on your browser.





