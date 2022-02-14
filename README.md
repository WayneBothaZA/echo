# echo

Simple echo REST service in golang

## usage

To call the service, POST this JSON message to the /echo endpoint

```code
{
    "message": "hello world"
}
```

## curl commandline

To call the service directly, just replace the IP and port with that of your service

```code
curl -X "POST" -H "Content-Type: application/json" -d '{"message":"hello world"}' 127.0.0.1:8080/echo
```
