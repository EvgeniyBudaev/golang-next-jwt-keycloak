Инициализация зависимостей
```
go mod init github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend
```

Удаление неиспользуемых зависимостей
```
go mod tidy -v
```

Переменные окружения
```
go get -u github.com/spf13/viper
```

Golang Keycloak API Package
https://github.com/Nerzal/gocloak
```
go get -u github.com/Nerzal/gocloak/v13
```

Validator
https://github.com/go-playground/validator
```
go get -u github.com/go-playground/validator/v10
```

Fiber
https://github.com/gofiber/fiber
```
go get -u github.com/gofiber/fiber/v2
```

JWT
https://github.com/golang-jwt/jwt
```
go get -u github.com/golang-jwt/jwt/v5
go get -u github.com/gofiber/contrib/jwt
```

Go Util
https://github.com/gookit/goutil
```
go get -u github.com/gookit/goutil
```