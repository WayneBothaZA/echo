# echo

Simple echo REST service in golang

## Running

The call the service directly use just replace the IP and port with that of your service

```code 
curl -X "POST" -H "Content-Type: application/json" -d '{"message":"hello"}' 10.152.183.36:8080/echo
```
