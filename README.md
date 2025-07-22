# go‑kata

> **Учебный репозиторий с набором «katas» — маленьких тренировочных задач на Go.**  
> Каждая задача оформлена отдельным пакетом, а общий CLI‑раннер (`cmd/runner`) позволяет
> запускать их по имени.

## Требования
- **Go 1.24+** (см. `go.mod`)
- *(опционально)* `golangci-lint` для статического анализа


## Список задач

| Имя | Описание | Условие |
|-----|----------|---------|
| `actionpurpose` | Парсер «A: B is (not) x!» | [`internal/actionpurpose`](internal/actionpurpose) |


## Команды

| Действие | Команда |
|----------|---------|
| Собрать CLI | `go build ./cmd/runner` |
| Запустить задачу | `go run ./cmd/runner --task <name>` |
| Показать все задачи | `go run ./cmd/runner --list` |
| Все тесты | `go test ./...` |
| Тесты конкретной задачи | `go test ./internal/<task>` |
| Один тест‑кейс | `go test ./internal/<task> -run '^TestCase$'` |


## Как запустить программу:
``` bash
# 1. Получить список доступных katas
go run ./cmd/runner --list
# → Доступные задачи:
#     actionpurpose

# 2. Запустить задачу и передать вход с клавиатуры/файла
go run ./cmd/runner --task actionpurpose < testdata/sample.txt
```

## Как запустить тесты
``` bash
# Все тесты во всех пакетах
go test ./...

# Только Action Purpose Parser
go test ./internal/actionpurpose

# Один конкретный тест‑кейс
go test ./internal/actionpurpose -run '^TestParseLine$'
```

## Как добавить новую задачу
1. Создать пакет `internal/<task‑name>/`.
2. Реализовать интерфейс `task.Runner` в файле `task.go`.
3. В `init()` вызвать `task.Register(Task{})`.
4. Написать юнит‑тесты `<task>_test.go`.
5. Добавь `blank‑import` в `cmd/runner/main.go`, например:
``` Go
    `_ "github.com/<you>/go-kata/internal/<task-name>"`
```
6. Убедиться, что `go test ./...` зелёный.
