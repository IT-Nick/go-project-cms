# Система управления корпоративными данными

[![Go Report Card](https://goreportcard.com/badge/github.com/)](https://goreportcard.com/report/github.com/)

## Больше имнофрмации о процессе разработки

Вы можете найти больше информации о цели проекта и этапах разработки программы по ссылкам ниже:

- [Прототип проекта](https://github.com/IT-Nick/demo)
- [Презентация прототипа](https://github.com/skoroDobavlyNavernoe)
- [Жизненный цикл разработки программы (Kanban-доска)](https://trello.com/)

# Установка (Docker)

Создание и Загрузка:

```bash
docker login docker.pkg.github.com -u <USERNAME> -p <GITHUB_TOKEN>
docker build -t  docker.pkg.github.com/IT-Nick/go-project-cms/rcms:latest .
# создали контейнер
docker push docker.pkg.github.com/IT-Nick/go-project-cms/rcms:latest
# выполнили push
```

Выгрузка и Запуск:

```bash
docker pull docker.pkg.github.com/IT-Nick/go-project-cms/rcms:latest
docker run docker.pkg.github.com/IT-Nick/go-project-cms/rcms:latest
```


### Установка нового проекта для анализа на SonarCloud

- В _SonarCloud_:
    - Нажмите _Plus_ в Правом верхнем углу
    - _Analyze New Project_
    - Нажмите на ссылку _GitHub app configuration_
    - Настроить SonarCloud
    - Выберите репозиторий и Сохраните
    - Вернитесь к анализу проекта
    - Отметьте галочкой добавленный репозиторий
    - Ниажмите Настроить
    - Нажмите настроить с помощью Travis
    - Скопируйте команду для шифрования токена Travis
    - Запустите `travis encrypt --com <TOKEN_YOU_COPPIED>`
    - Заполните поле `secure` в `.travis.yml` полученной строкой
    - Push
- В Travis:
    - Установите `DOCKER_USERNAME`
    - Установите `DOCKER_PASSWORD` вашим токеном Github

### Установка CodeClimate
- Заходим на <https://codeclimate.com/github/repos/new>
- Добавляем репозиторий
- Заходим на вкладку Тестовое покрытие
- Копируем идентификатор Reporter ID
- Переходим в Travis и открываем Настройки для репозитория
- Добавляем переменную окружения: name: `CC_TEST_REPORTER_ID`, value: _Скопировать из CodeClimate_
