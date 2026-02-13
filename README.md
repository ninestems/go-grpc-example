### Быстрый старт
```shell
  make buf-install  
```
#### Генерация кода
```shell
  make server-gen
  make client-gen
```
#### Запуск сервера
```shell
  make run-server
```
#### Запуск клиента
```shell
  make run-client 
```

### Как это устроено
Сервер и клиент имеют независимо сгенерированный код из одних и тех же proto файлов:
- Сервер: `proto/hello/v1/*.proto` → `internal/controller/proto`
- Клиент: `github.com/company/common-proto.git` → `internal/client/proto`

Это позволяет:
- Серверу всегда использовать актуальные локальные proto
- Клиенту работать с версией proto из общего репозитория
- Избежать конфликтов зависимостей

### Дополнительная информация по buf настройкам
В примере файлы разнесены по разным директориям 

| **Аспект**           | **buf.yaml**                                                | **buf.gen.yaml**                                                           |
|----------------------|-------------------------------------------------------------|----------------------------------------------------------------------------|
| **Назначение**       | Конфигурация модуля и его зависимостей                      | Конфигурация генерации кода                                                |
| **Основная функция** | Определяет ЧТО есть в модуле и его требования               | Определяет КАК генерировать код из proto                                   |
| **Содержит**         | Зависимости, правила линтинга, проверки на breaking changes | Плагины, выходные директории, настройки go_package                         |
| **Обязательность**   | Обязателен для публичных модулей в Buf Schema Registry      | Опционален, нужен только для генерации кода                                |
| **Где лежит**        | Всегда в корне модуля (рядом с proto)                       | Может лежать в любом месте (у вас в bufconfig/server/ и bufconfig/client/) |
| **Зависимости**      | Указывает внешние модули (например, googleapis)             | Не указывает зависимости, только входные данные                            |
| **Входные данные**   | Определяет, какие proto входят в модуль                     | Определяет, откуда брать proto (локально или из git)                       |
| **Плагины**          | Не содержит                                                 | Содержит список плагинов для генерации                                     |
| **Линтинг**          | Настраивает правила линтинга                                | Не настраивает линтинг                                                     |
| **Breaking changes** | Настраивает проверки                                        | Не участвует                                                               |
| **Пример**           | Указывает, что модуль зависит от googleapis                 | Указывает, что нужен protoc-gen-go и protoc-gen-go-grpc                    |

#### Сервер
##### buf.gen
```yaml
# Файл: bufconfig/server/buf.yaml
version: v2                              # Используем версию v2 конфигурации Buf (более новая и стабильная)
deps:                                    # Внешние зависимости proto, которые нужны для компиляции
  - buf.build/grpc-ecosystem/grpc-gateway    # Зависимость для grpc-gateway (нужна если используем HTTP REST)
  - buf.build/googleapis/googleapis          # Зависимость с стандартными google proto (timestamp, duration и т.д.)
modules:                                 # Описание модулей в этом репозитории
  - path: proto                           # Путь до директории с proto файлами
lint:                                    # Настройки линтера для proto файлов
  use:                                    # Какие правила линтера использовать
    - STANDARD                            # Использовать стандартный набор правил
breaking:                                # Настройки проверки на breaking changes
  use:                                    # Какие правила использовать для обнаружения breaking changes
    - FILE                                # Проверять на уровне файлов (не только пакетов)
```
##### buf.yaml.gen
```yaml
# Файл: bufconfig/server/buf.gen.yaml
version: v2                              # Версия конфигурации генерации
clean: true                              # Очищать выходную директорию перед генерацией
managed:                                 # Управление настройками go_package
  enabled: true                           # Включить управление
  disable:                                 # Отключить для определенных модулей
    - module: buf.build/googleapis/googleapis    # Не менять настройки для google proto
  override:                                # Переопределить настройки
    - file_option: go_package_prefix       # Какое поле переопределяем
      value: github.com/ninestems/go-grpc-example/internal/controller/proto  # Префикс для go_package
plugins:                                 # Плагины для генерации кода
  - local: protoc-gen-go                  # Использовать локально установленный плагин
    out: internal/controller/proto         # Куда генерировать .pb.go файлы
    opt: paths=source_relative              # Опции: сохранять относительные пути импортов
  - local: protoc-gen-go-grpc              # Плагин для генерации gRPC кода
    out: internal/controller/proto          # Куда генерировать _grpc.pb.go файлы
    opt: paths=source_relative               # Опции: сохранять относительные пути импортов
inputs:                                  # Входные данные (откуда брать proto)
  - directory: proto                       # Из локальной директории proto
```
#### Client
##### buf.yaml
Клиент не нуждается в buf.gen, так как не проводит валидацию/линтинг
##### buf.gen.yaml
```yaml
# Файл: bufconfig/client/buf.gen.yaml
version: v2                              # Версия конфигурации генерации
clean: true                              # Очищать выходную директорию перед генерацией
managed:                                 # Управление настройками go_package
  enabled: true                           # Включить управление
  disable:                                 # Отключить для определенных модулей
    - module: buf.build/googleapis/googleapis    # Не менять настройки для google proto
  override:                                # Переопределить настройки
    - file_option: go_package_prefix       # Какое поле переопределяем
      value: github.com/ninestems/go-grpc-example/internal/client/proto  # Префикс для go_package
plugins:                                 # Плагины для генерации кода
  - local: protoc-gen-go                  # Использовать локально установленный плагин
    out: internal/client/proto              # Куда генерировать .pb.go файлы (клиентская часть)
    opt: paths=source_relative              # Опции: сохранять относительные пути импортов
  - local: protoc-gen-go-grpc              # Плагин для генерации gRPC кода
    out: internal/client/proto              # Куда генерировать _grpc.pb.go файлы (клиентская часть)
    opt: paths=source_relative               # Опции: сохранять относительные пути импортов
inputs:                                  # Входные данные (откуда брать proto)
  - git_repo: https://github.com/company/common-proto.git  # Из удаленного git репозитория
    subdir: proto                          # Из поддиректории proto в том репозитории
    branch: master                            # Из ветки main
```