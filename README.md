# go-manufacturer

В базе данных pgsql хранятся производители

manufacturer ( id varchar, details jsonb)

В поле details есть json в котором в корне структуры есть поле needUpdate (bool)

Написать скрипт который бы доставал тех производителей у которых needUpdate=true

батчами, по N элементов и отправлял их в http rest api (которое нужно тоже написать), получал подтверждение  от api и помечал  как needUpdate false в базе.

Будет плюсом:

- docker-compose для решения
- замер латенси ответов http rest api и запросов в pgsql
- настроенный dashboard для этих метрик
