# RestconfThread

use golang to send multiple concurrent RESTCONF calls to NSO

```
FMARANGI-M-309Y:cfe_migration fmarangi$ go run call.go 
Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 46.205394078s
Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 46.205474677s
  Received At: 2020-10-30 18:09:43.85723 +0100 CET m=+46.206401191
  Body       :
 {
  "tailf-ncs:output": {
    "result": false,
    "info": "Connection to R4 timed out"
  }
}

  Received At: 2020-10-30 18:09:43.857263 +0100 CET m=+46.206433761
  Body       :
 {
  "tailf-ncs:output": {
    "result": false,
    "info": "Connection to R3 timed out"
  }
}

Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 46.211600168s
  Received At: 2020-10-30 18:09:43.863357 +0100 CET m=+46.212527873
  Body       :
 {
  "tailf-ncs:output": {
    "result": false,
    "info": "Connection to R5 timed out"
  }
}

Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 46.217776559s
  Received At: 2020-10-30 18:09:43.86955 +0100 CET m=+46.218721536
  Body       :
 {
  "tailf-ncs:output": {
    "result": false,
    "info": "Connection to R1 timed out"
  }
}

Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 46.251791316s
  Received At: 2020-10-30 18:09:43.903612 +0100 CET m=+46.252783232
  Body       :
 {
  "tailf-ncs:output": {
    "result": false,
    "info": "Connection to R2 timed out"
  }
}

Execution Time:  46.2522692s
FMARANGI-M-309Y:cfe_migration fmarangi$ 
FMARANGI-M-309Y:cfe_migration fmarangi$ go run call.go 
Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 5.296616778s
  Received At: 2020-10-30 18:09:54.349394 +0100 CET m=+5.297868666
  Body       :
 {
  "tailf-ncs:output": {
    "result": true
  }
}

Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 5.323241569s
  Received At: 2020-10-30 18:09:54.37602 +0100 CET m=+5.324494396
  Body       :
 {
  "tailf-ncs:output": {
    "result": true
  }
}

Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 5.404569696s
  Received At: 2020-10-30 18:09:54.457348 +0100 CET m=+5.405821847
  Body       :
 {
  "tailf-ncs:output": {
    "result": true
  }
}

Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 5.537498229s
  Received At: 2020-10-30 18:09:54.590338 +0100 CET m=+5.538811105
  Body       :
 {
  "tailf-ncs:output": {
    "result": true
  }
}

Response Info:
  Error      : <nil>
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 5.911627173s
  Received At: 2020-10-30 18:09:54.964413 +0100 CET m=+5.912883701
  Body       :
 {
  "tailf-ncs:output": {
    "result": true
  }
}

Execution Time:  5.912183933s
```


```
FMARANGI-M-309Y:logs fmarangi$ tail -f localhost\:8080.access 
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R4/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R3/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R5/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R1/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R2/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
```
