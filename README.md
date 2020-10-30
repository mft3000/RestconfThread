# RestconfThread

use golang to send multiple concurrent RESTCONF calls to NSO

golang code:

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

show devices commit-queue
```                                                    WAITING  TRANSIENT             IS                              
ID             TAG            AGE  STATUS  DEVICES  FOR      ERRORS     COMPLETED  ATOMIC  NAME  REASON  NODE  ID  
-------------------------------------------------------------------------------------------------------------------
1604080758084  sync-from:565  1    locked  [ R3 ]   -        -          -          true                            
1604080758085  sync-from:566  1    locked  [ R4 ]   -        -          -          true                            
1604080758086  sync-from:562  1    locked  [ R1 ]   -        -          -          true                            
1604080758087  sync-from:564  1    locked  [ R5 ]   -        -          -          true                            
1604080758088  sync-from:563  1    locked  [ R2 ]   -        -          -          true                            

```

tail -f localhost\:8080.access 
```
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R4/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R3/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R5/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R1/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
127.0.0.1 - admin [30/Oct/2020:18:08:08 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R2/tailf-ncs:sync-from HTTP/1.1" 200 94 "-" "go-resty/2.3.0 (https://github.com/go-resty/resty)"
```





python code:
```
FMARANGI-M-309Y:cfe_migration fmarangi$ python3 call.py 
{
  "tailf-ncs:output": {
    "result": true
  }
}

{
  "tailf-ncs:output": {
    "result": true
  }
}

{
  "tailf-ncs:output": {
    "result": true
  }
}

{
  "tailf-ncs:output": {
    "result": true
  }
}

{
  "tailf-ncs:output": {
    "result": true
  }
}

```

show devices commit-queue
```
fmarangi@ncs# show devices commit-queue
                                                    WAITING  TRANSIENT             IS                              
ID             TAG            AGE  STATUS  DEVICES  FOR      ERRORS     COMPLETED  ATOMIC  NAME  REASON  NODE  ID  
-------------------------------------------------------------------------------------------------------------------
1604080703160  sync-from:547  5    locked  [ R1 ]   -        -          -          true                            

fmarangi@ncs# show devices commit-queue
System message at 2020-10-30 18:58:28...
Commit performed by admin via http using rest.
fmarangi@ncs# show devices commit-queue
                                                    WAITING  TRANSIENT             IS                              
ID             TAG            AGE  STATUS  DEVICES  FOR      ERRORS     COMPLETED  ATOMIC  NAME  REASON  NODE  ID  
-------------------------------------------------------------------------------------------------------------------
1604080708871  sync-from:550  0    locked  [ R2 ]   -        -          -          true                            

fmarangi@ncs# show devices commit-queue
System message at 2020-10-30 18:58:32...
Commit performed by admin via http using rest.
fmarangi@ncs# show devices commit-queue
                                                    WAITING  TRANSIENT             IS                              
ID             TAG            AGE  STATUS  DEVICES  FOR      ERRORS     COMPLETED  ATOMIC  NAME  REASON  NODE  ID  
-------------------------------------------------------------------------------------------------------------------
1604080712070  sync-from:553  0    locked  [ R3 ]   -        -          -          true                            

fmarangi@ncs# show devices commit-queue
System message at 2020-10-30 18:58:34...
Commit performed by admin via http using rest.
fmarangi@ncs# show devices commit-queue
                                                    WAITING  TRANSIENT             IS                              
ID             TAG            AGE  STATUS  DEVICES  FOR      ERRORS     COMPLETED  ATOMIC  NAME  REASON  NODE  ID  
-------------------------------------------------------------------------------------------------------------------
1604080714654  sync-from:556  0    locked  [ R4 ]   -        -          -          true                            

fmarangi@ncs# show devices commit-queue
System message at 2020-10-30 18:58:37...
Commit performed by admin via http using rest.
fmarangi@ncs# show devices commit-queue
                                                    WAITING  TRANSIENT             IS                              
ID             TAG            AGE  STATUS  DEVICES  FOR      ERRORS     COMPLETED  ATOMIC  NAME  REASON  NODE  ID  
-------------------------------------------------------------------------------------------------------------------
1604080717247  sync-from:559  0    locked  [ R5 ]   -        -          -          true                            

```

tail -f localhost\:8080.access 
```
127.0.0.1 - admin [30/Oct/2020:18:52:56 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R4/tailf-ncs:sync-from HTTP/1.1" 200 76 "-" "python-requests/2.22.0"
127.0.0.1 - admin [30/Oct/2020:18:52:59 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R5/tailf-ncs:sync-from HTTP/1.1" 200 76 "-" "python-requests/2.22.0"
127.0.0.1 - admin [30/Oct/2020:18:55:23 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R1/tailf-ncs:sync-from HTTP/1.1" 200 51 "-" "python-requests/2.22.0"
127.0.0.1 - admin [30/Oct/2020:18:55:26 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R2/tailf-ncs:sync-from HTTP/1.1" 200 51 "-" "python-requests/2.22.0"
127.0.0.1 - admin [30/Oct/2020:18:55:29 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R3/tailf-ncs:sync-from HTTP/1.1" 200 51 "-" "python-requests/2.22.0"
127.0.0.1 - admin [30/Oct/2020:18:55:32 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R4/tailf-ncs:sync-from HTTP/1.1" 200 51 "-" "python-requests/2.22.0"
127.0.0.1 - admin [30/Oct/2020:18:55:35 +0100] "POST /restconf/operations/tailf-ncs:devices/tailf-ncs:device=R5/tailf-ncs:sync-from HTTP/1.1" 200 51 "-" "python-requests/2.22.0"
```
