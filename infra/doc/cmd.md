### Миграции
1. Экспортируем переменные
```bash
source ./goose_env
```
2. Создать миграцию
```bash
goose -dir db/migrations create alter_book_table sql
```
3. Миграции вперед и назад
```bash
goose -dir db/migrations up
goose -dir db/migrations down
```
4. Если не экспортировать переменные
```bash
goose -dir db/migrations postgres "postgresql://postgres:postgres@PostgreSQL-16:5432/books?sslmode=disable" up
```