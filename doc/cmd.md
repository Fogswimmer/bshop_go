### Миграции
1. Экспортируем переменные
```bash
source ./goose_env
```
2. Миграции вперед и назад
```bash
goose -dir db/migrations up
goose -dir db/migrations down
```