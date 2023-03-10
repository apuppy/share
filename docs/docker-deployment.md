# docker deployment

## docker build
```shell
docker build -t share-backend .
```

## docker run
```shell
docker run -it --rm --name share-backend-demo -p 8888:8080 share-backend:latest
```