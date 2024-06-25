create a file in root folder of the project name it as `.env`
add the following contents into the .env file

```
API_SERVER_HOST=http://127.0.0.1
API_SERVER_PORT=8000
WS_SERVER_HOST=localhost
WS_SERVER_PORT=8080
```

open a termial and go the project folder and export the environmental varibles using
`export $(grep -v '^#' .env | xargs)`

Build the go application usign the command `go build`
and notice there is a executable file created in project folder now run the file `./scheduler` alternatively you can run the application without building it by running the command `go run main.go router.go`
