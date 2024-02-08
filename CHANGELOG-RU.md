## 0.0.1

## Добавлено

### Библиотека

#### Ядро

* Добавлены типы ошибок, форматы и статусы (раздел core/info).
* Добавлена структура описание ответа от сервисом api.
* Добавлено обобщённое выполнения запросов по api и получение результатов.
* Добавлен интерфейс ApiMethod описывающий требования к методом api.

#### Api

* Добавлен базовый метод описывающий запрос по api.
* Добавлены методы взаимодействия с DNS: *getData*, *changeRecords*.
* Реализован создатель записей для запроса на изменение dns записей через метод *changeRecords*.

### Репозиторий

* Добавлено описание README на русском и английском языке.
* Настроена работа CI.
* Добавлен пример применения библиотеки в другом репозитории.