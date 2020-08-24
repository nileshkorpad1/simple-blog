# Simple Blog Using GO, Docker, Swagger, MongoDB

Backend API for a simple blog system.

### Requirements

```bash
1. Go
2. MongoDB
2. Docker
```

## How to Run Application

#### checkout Application in $GOPATH/src/github.com/<UserName>
```python
 https://github.com/nileshkorpad1/simple-blog.git
```

#### Then Run Below Commands

```bash
$ go build
$ go run main.go
```

### Using Docker Compose To Setup the images:

```python
   $ docker-compose build
```

### To Start the development server:

```python
   $ docker-compose up
```

### To stop the development server:

```python
   $ docker-compose stop
```


### Stop Docker development server and remove containers:

```python
   $ docker-compose down
```

## To generate APIs using Swagger

#### Requirement:

1. go-swagger

#### Documentation

[https://goswagger.io/](https://goswagger.io/)

1. To install Swagger: `alias swagger='docker run --rm -e GOPATH=${GOPATH}:/go -v $(pwd):$(pwd) -w $(pwd) -u $(id -u):$(id -u) stratoscale/swagger:v1.0.21'`

2. To validate Swagger yml file: `swagger validate swagger/swagger.yaml`

3. To generate Go Server components and docs: `swagger generate server -A Simple-Blog -f swagger/swagger.yaml`

## Testing

```python
   $ go test -v
```

### Demo

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
