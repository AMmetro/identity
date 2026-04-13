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

# Routes:
    - routes:    all routs 
        - events:    events controlers
        - users:     users controlers
    
    - models:        repositories 
    - db:            initial DB
    - utils:         utils

/cmd            // точка входа в приложение, отдельно для каждого сервиса или билда
/internal       // внутренние пакеты, не доступные вне модуля
/pkg            // публичные библиотеки, используемые другими проектами
/config         // конфигурационные файлы и схемы
/apis           // определения API, например Protobuf или OpenAPI
/migrations     // миграции базы данных
/pkg/logger     // логирование
/pkg/middleware // промежуточное ПО (middleware)
/services       // бизнес-логика, связанные с конкретной domain
/handlers       // HTTP handlers / controllers
/models         // модели данных
/utils          // вспомогательные функции и утилиты
/tests          // тесты
Dockerfile     // описание контейнера
Makefile       // сборочные команды
