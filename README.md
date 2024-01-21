go mod init echo-template
go get github.com/labstack/echo/v4

// nodemon --ext go,mjs --watch . --watch ../cmd index.mjs
// nodemon --ext go,mjs,html --watch ../cmd --watch ../views --watch . index.mjs
// nodemon --exec "go run ./cmd/main.go" --ext go,html --watch ./cmd --watch ./views --signal SIGTERM
