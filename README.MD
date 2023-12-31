# Сервис на `Go` в облаке за пять минут

Как расположиться бессерверно в `Яндекс облаке` с использованием библиотеки `Chi`.
----

Это простой шаблон, который позволит запустить первую версию вашего приложения в облаке. Делается за минуту:
+ Загрузите `go.mod` и `main.go` в вашу функцию
+ Поставьте точка входа - `main.Router`
    ![точка входа](./assets/code.png "Вот сюда")
+ Поставьте галочку `публичная функция`
    ![галочка](./assets/public.png "Вот сюда")


Потом перейдите на `https://functions.yandexcloud.net/d4e3...`(написан там где галочку ставили). Вы победили!

## Стоп! Это ещё не все

Чтобы у нас работало, что-то кроме одного маршрута, мы добавим `Api Gateway`. Там все очень просто. 

Это делается на тот, случай если ты решишь, что тебе нужен ещё например такой адрес `/info`

```go
router.Get(/info, infoHandler)
```

1. Запомнить идентификатор функции. Или в адресе `https://functions.yandexcloud.net/<идентификатор>`, или в списке всех функций
    ![запомнить функцию](./assets/func-id.png "Где находится идетификатор")
1. Перейдем в сервис `Api Gateway` в яндекс облаке и создадим новый гейт
    ![создать шлюз](./assets/gateway-create.png "Где находится кнопка")
2. Запишем такую конфигурацию в окне создания шлюза, и поменяем строчку `function_id` на полученный в первом пунтке идентификатор

    ```openapi
    openapi: 3.0.0
    info:
      title: Sample API
      version: 1.0.0
    paths:
      /{url+}:
        x-yc-apigateway-any-method:
          parameters:
          - explode: false
            in: path
            name: url
            required: false
            style: simple
          x-yc-apigateway-integration:
            function_id: твой_идентификатор
            tag: $latest
            type: cloud_functions
    ```

    ![заполнить конфигурацию](./assets/gateway-config.png "Куда писать")
4. Теперь твоя функция работает по такому адресу `https://d5dphasdln.apigw.yandexcloud.net`, его можно посмотреть в поле служебный домен на открывшейся странице
    ![Адрес шлюза](./assets/gateway-url.png)

Готово! 🙌

## Что дальше?

Форкните на гитхабе и установите себе локально репозиторий, чтобы можно было работать с кодом. 

![Форкните](https://docs.github.com/assets/cb-79331/mw-1440/images/help/repository/fork_button.webp)

Ещё вот есть интересный [гайд по git](https://habr.com/ru/articles/541258/) или прекрасный [курс на Яндекс.Практикуме](https://practicum.yandex.ru/git-basics/)

### Установить локально

1. зайдите в папку
2. скачайте в нее репозиторий
    ```bash
    git clone https://github.com/<ваше_имя>/<ваш_репозиторий> .
    ```

### Локальный запуск

Чтобы сервер можно было запускать локально добавьте в конце `main.go` такой код

```go
func main(){
    http.ListenAndServe(":8080", router)
}
```

### Планы

В дальнейшем планируется добавить полезных функций для чтения контекста, и прочего.
Записать видео, как создать функцию

# Автор

Дмитрий Фроленко 2023
