# GO Live Reload - Automatic Refresh

## Introduction

## How It Works

# Repo

Clone or fork this repo to a working directory

https://github.com/dmh2000/go-live-reload

## Dependencies

Install the following dependencies

- golang
- node.js
  - https://nodejs.org
  - download and install
- nodemon
  - https://nodemon.io/
  - npm install -g nodemon

## Setup

From root of cloned repo:

```bash
go mod tidy
cd ws
npm install
cd ..
```

## Startup

Nodemon options:

- --**ext** specifies what type of file extension to watch
- --**watch** specifies a directory to watch (can be multiple)
- --**exec** use to run a non-node program
- --**signal SIGTERM** forces the currently running executable to terminate (so it won't cause a duplicate port error)

Note : **cmd/main.go** must be run at the top level of the project, not in **cmd**. It needs to access the **views** directory to fetch the html template.

Run each of these commands in separate terminals

```bash
# restart the local websocket server so web app will detect disconnect and reload
cd ws
nodemon --ext go,mjs,html --watch ../cmd --watch ../views  index.mjs
```

```bash
# restart the web app to upload the changes
nodemon --exec "go run ./cmd/main.go" --ext go,html --watch ./cmd --watch ./views --signal SIGTERM
```
