# shorty
shorty



## Running Application on Docker Container

This application you can run on docker container.

### Prerequisite :
* Docker
* OS Linux Docker Image, 
* Golang version 1.14 or latest

### Environment Variable available


### build docker image
```bash
$ docker build -t {IMAGE_NAME} -f deployment/dockerfiles/dockerfile--dev .
# example
$ docker build -t shorty -f deployment/dockerfiles/dockerfile-dev .
```

### run docker container after build the image
```bash
# example run http serve:
$ docker run -i --name shorty -p 8081:8081 -t shorty

```


## Running on your local machine

Linux or MacOS

## Installation guide
#### 1. install go version 1.14.2     
```bash
# please read this link installation guide of go
# https://golang.org/doc/install
```

#### 2. Create directory workspace    
```bash
run command below: 
mkdir $HOME/go
mkdir $HOME/go/src
mkdir $HOME/go/pkg
mkdir $HOME/go/bin
mkdir -p $HOME/go/src/github.com/aryayunanta-ralali/shorty
chmod -R 775 $HOME/go
cd $HOME/go/src/github.com/ralali
export GOPATH=$HOME/go
```    
```bash
# edit bash profile in $HOME/.bash_profile        
# add below to new line in file .bash_profile         
    PATH=$PATH:$HOME/bin:$HOME/go/bin
    export PATH  
    export GOPATH=$HOME/go 
# run command :
source $HOME/.bash_profile
```

#### 3. Build the application    
```bash
# run command :
cd $HOME/go/src/github.com/aryayunanta-ralali/shorty
git clone -b development https://github.com/aryayunanta-ralali/shorty .
cd $HOME/go/src/github.com/aryayunanta-ralali/shorty
go mod tidy && go mod download && go mod vendor
cp config/app.yaml.tpl config/app.yaml     
# edit config/app.yaml with server environtment
go build

# run application after build or create on supervisord 
./shorty serve-http
```


### 4. Health check Route PATH
```bash
{{host}}/in/health
```


#### Postman Collection
```go
```

### Database Migration
migration up
```bash
go run main.go db:migrate up
```

migration down
```bash
go run main.go db:migrate down
```

migration reset
```bash
go run main.go db:migrate reset
```

migration reset
```bash
go run main.go db:migrate reset
```

migration redo
```bash
go run main.go db:migrate redo
```

migration status
```bash
go run main.go db:migrate status
```

create migration table
```bash
go run main.go db:migrate create {table-name} sql

# example
go run main.go db:migrate create users sql
```

to show all command
```bash
go run main.go db:migrate
```

## run docker compose on your local machine
```bash
docker-compose -f deployment/docker-compose.yaml --project-directory . up -d --build
```