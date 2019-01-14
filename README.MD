Простой блокирующий FCGI-HTTP сервер реализующий паттерн Producer-Consumer. 
Чтение блокируется если нет данных для чтения и запись блокируется если нет читающего клиента. Все хедеры копируются из запроса.

Для запуска неоходимо передать в первом параметре порт, на котором будет запущен сервер. 

Пример: sq.exe 8080

Запись методом PUT

```
➜  ~ curl -v -X PUT http://localhost:8080 -d 'Hello' -H "Content-type: application/text" -H "AAA: 1"
* Rebuilt URL to: http://localhost:8080/
*   Trying localhost...
* TCP_NODELAY set
* Connected to localhost (localhost) port 8080 (#0)
> PUT / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.58.0
> Accept: */*
> Content-type: application/text
> AAA: 1
> Content-Length: 5
> 
* upload completely sent off: 5 out of 5 bytes
< HTTP/1.1 200 OK
< Date: Fri, 11 Jan 2019 12:54:39 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
➜  ~ 
```

Чтение методом GET

```
curl -v -X GET http://localhost:8080
Note: Unnecessary use of -X or --request, GET is already inferred.
* Rebuilt URL to: http://localhost:8080/
*   Trying localhost...
* TCP_NODELAY set
* Connected to localhost (localhost) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.58.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Aaa: 1
< Accept: */*
< Content-Length: 5
< Content-Type: application/text
< User-Agent: curl/7.58.0
< Date: Fri, 11 Jan 2019 12:54:39 GMT
< 
* Connection #0 to host localhost left intact
Hello%    
```