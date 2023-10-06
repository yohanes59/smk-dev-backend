# Intro Web Server

This project contains an introduction to how to create a web server for student management on 
the Backend Developer learning path in the SMKDEV Bootcamp Scholarship program.

![Logo](https://www.smk.dev/wp-content/uploads/SMKDEV-Logo-Long-150x38.png)

## Run Locally

Clone the project

```bash
  git clone https://github.com/yohanes59/smk-dev-backend.git
```

Go to the project directory

```bash
  cd smk-dev-backend
```

Start the server

```bash
  go run main.go
```

Import the thunder-collection_Intro Web Server.json to vscode extension Thunder Client

![import collection thunder client](https://github.com/yohanes59/smk-dev-backend/assets/80000614/f5c981fb-3bcc-4baa-9eeb-accb8ff994b4)


## API Reference

#### Get all students data

```http
  GET http://localhost:8080/students
```

#### Create new student data

```http
  GET http://localhost:8080/students
```

Body - JSON Content Format
```json
{
  "name": "tes",
  "age": 12,
  "grade": "A"
}
```

#### Get student data by id

```http
  GET http://localhost:8080/student/${id}
```

| Parameter | Type      | Description                            |
| :-------- | :-------- | :------------------------------------- |
| `id`      | `integer` | **Required**. Id of student to updated |

#### Update student data by id

```http
  PUT http://localhost:8080/students/${id}
```

| Parameter | Type      | Description                          |
| :-------- | :-------- | :----------------------------------- |
| `id`      | `integer` | **Required**. Id of student to fetch |


Body - JSON Content Format
```json
{
  "name": "tes1",
  "age": 21,
  "grade": "B"
}
```

#### Delete student data by id

```http
  DELETE http://localhost:8080/student/${id}
```

| Parameter | Type      | Description                            |
| :-------- | :-------- | :------------------------------------- |
| `id`      | `integer` | **Required**. Id of student to deleted |

## Tech Stack

**Server:** Go, Echo
