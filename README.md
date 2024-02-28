# skillbox_final
final project on the Golang course at Skillbox education platform

## Шаблон под финальную работу

Симулятор предпочтительно запускать из его собственной директории

```
danil@FunBox:~/diplomtemp$ cd simulator

danil@FunBox:~/diplomtemp/simulator$ go run main.go
```

После запуска в директории simuator будут сгенерированы файлы с тестовыми данными, именно их мы и будем забирать нашим сервисом для дальнейшей обработки

```
simulator/sms.data
simulator/voice.data
simulator/email.data
simulator/billing.data
```

Другая часть данных будет доступна нам по http запросу на порт 8383

```
http://127.0.0.1:8383/mms
http://127.0.0.1:8383/support"
http://127.0.0.1:8383/accendent"
```

Для демонстрации конечного результата можно открыть HTML страницу из директории симулятора

```
status_page.html
```

Для работы с нашим сервисом агрегатором создана директория web, страница index.html будет делать запрос на http://127.0.0.1:8282/api

По этому адресу сервис должен отдать всю собранную информацию.


