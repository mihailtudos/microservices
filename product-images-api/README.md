# Product images api

## Description

## Testing locally 

To store files:

```sh
    curl -X POST -H "Content-Type: application/octet-stream" --data-binary "@test.png" http://localhost:8080/images/1/test.png
```

To download files:

```sh
    curl -X POST -H "Content-Type: application/octet-stream" --data-binary "@test.png" http://localhost:8080/images/1/test.png
```
