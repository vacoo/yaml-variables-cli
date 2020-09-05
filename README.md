# Утилита для замены переменных в yaml файле
Заменяет любую переменную например ${VAR} на значение переменной окружения VAR

### Скомпилировать под Linux
`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o yaml-variables-cli main.go`

### Запуск
`VAR=test yaml-variables-cli input.yml output.yml`