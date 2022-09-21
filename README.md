# httpDump
Dump http request

```
# docker run --rm --name=httpdump -p 80:80 he11olx/httpdump
================ START httpDump http://127.0.0.1:80 ================
-----------[1655259956] [2022-06-15 10:25:56] recv from 127.0.0.1:52918-----------
GET / HTTP/1.1
Host: localhost
Accept: */*
User-Agent: curl/7.64.1


-----------[1655259980] [2022-06-15 10:26:20] recv from 127.0.0.1:52919-----------
POST / HTTP/1.1
Host: localhost
Accept: */*
Auht: 123
Content-Length: 7
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.64.1

a=1,b=2
```

```
[root@he11o ~]# curl localhost:80
GET / HTTP/1.1
Host: localhost
Accept: */*
User-Agent: curl/7.64.1

[root@he11o ~]# curl localhost:80 -XPOST -d 'a=1,b=2' -H 'Auht: 123'
POST / HTTP/1.1
Host: localhost
Accept: */*
Auht: 123
Content-Length: 7
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.64.1

a=1,b=2
```