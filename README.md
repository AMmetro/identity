# install gin
go get -u github.com/gin-gonic/gin

# run
go run main.go

# createEvent: 
POST http://localhost:8080/events

  "name": "samsung1",
  "Description": "1000",
  "Location": "sgdg",
  "DateTime": "2006-01-02T15:04:05Z"

# getEvent: 
GET http://localhost:8080/events