# BookShop
- Тренировочный api + веб-сервер на Go.
## Development
- Установка air 
```bash
    go install github.com/air-verse/air@latest
```
- Билд контейнеров без сидирования бд
```makefile
    make init-dev
```
- C сидированием
```makefile
    make init-dev-seed
```
## Production
- Те же команды, но заменить init-dev на init-prod