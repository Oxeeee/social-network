# Start
## How to run
1. Start database
```shell
docker compose up -d
```
2. Add config (ask the administrator)

3. 
```shell
make run
```

### swagger docs will be avalible on http://localhost:3000/swagger/index.html

## How to generate swagger docs
1.
```shell
go install github.com/swaggo/swag/cmd/swag@v1.8.12
```
2.
```shell
make swag
```

1. Аутентификация и авторизация (Auth) [done]
 • Регистрация (email / пароль / логин)
 • Логин / логаут
 • JWT Refresh + Access
 ---

2. Пользовательский профиль (User Profile)
 • Аватар
 • Ник / имя / описание
 • Возможность редактировать профиль

---

3. Лента и посты (Feed & Posts)
 • Создание поста (текст, изображение)
 • Удаление поста
 • Редактирование поста
 • Просмотр постов других пользователей
 • Лента: собственная, глобальная

---

4. Социальные действия
 • Лайки
 • Комментарии
 • Сохранения (в закладки)

---

7. Поиск по категориям

Спорт, IT и тд

Страницы: Профиль, Новости, Добавлениме поста
