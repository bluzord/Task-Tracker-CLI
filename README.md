# 📌 Task Tracker CLI

Простой **CLI task tracker**, написанный на **Go**, с хранением задач в JSON-файле.

---

## Возможности и особенности

* Добавление, удаление и обновление задач
* Изменение статусов задач (`todo` → `in-progress` → `done`)
* Вывод задач (в том числе фильтрация по статусу)
* Цветной и аккуратно отформатированный вывод
* Хранение данных в локальном JSON-файле

---

## Установка

1. Склонируйте репозиторий
```bash
git clone https://github.com/bluzord/Task-Tracker-CLI
```
2. Перейдите в папку с проектом
```bash
cd ./Task-Tracker-CLI
```
3. Соберите исполняемый файл
```bash
go build . # Для Windows

go build -o task-cli . # Для Linux
```

4. (опционально для Linux) Добавьте файл в PATH:
```bash
mv task-cli /usr/local/bin/
```

---

## Использование
```bash
task-cli add [description]              # добавить задачу
task-cli update [id] [description]      # обновить описание задачи
task-cli delete [id]                    # удалить задачу
task-cli list                           # показать все задачи
task-cli list [done|todo|in-progress]   # показать задачи по статусу
task-cli mark-in-progress [id]          # отметить задачу как "в работе"
task-cli mark-done [id]                 # отметить задачу как "выполнено"
```

---

## Пример работы программы
```bash
$ task-cli add "Test 1"
Task added: [1] Test 1
$ task-cli add "Test 2"
Task added: [2] Test 2
$ task-cli add "Test 3"
Task added: [3] Test 3
$ task-cli list
[3] (todo)           Test 3                                             | created 2025-11-01 16:16:51
[2] (todo)           Test 2                                             | created 2025-11-01 16:16:48
[1] (todo)           Test 1                                             | created 2025-11-01 16:16:39


$ task-cli mark-done 2
Task marked: [2] (done)
$ task-cli mark-in-progress 1
Task marked: [1] (in-progress)
$ task-cli list 
[3] (todo)           Test 3                                             | created 2025-11-01 16:16:51
[1] (in-progress)    Test 1                                             | updated 2025-11-01 16:17:32
[2] (done)           Test 2                                             | updated 2025-11-01 16:17:20


$ task-cli update 1 "Update 1"
Task updated: [1] Update 1
$ task-cli list               
[3] (todo)           Test 3                                             | created 2025-11-01 16:16:51
[1] (in-progress)    Update 1                                           | updated 2025-11-01 16:19:23
[2] (done)           Test 2                                             | updated 2025-11-01 16:17:20


$ task-cli delete 2
Task deleted: [2] Test 2
$ task-cli list    
[3] (todo)           Test 3                                             | created 2025-11-01 16:16:51
[1] (in-progress)    Update 1                                           | updated 2025-11-01 16:19:23


$ task-cli list todo
[3] (todo)           Test 3                                             | created 2025-11-01 16:16:51
```