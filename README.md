# Setup and Installation

## Clone the repository
```
git clone git@github.com:tbson/base-go.git
cd base-go
cp docker/.env.example docker/.env # then update the content
cp docker/volumes/nginx/conf.d/default.conf.example docker/volumes/nginx/conf.d/default.conf # then update the content
```

## Install SSL certificate

Install `docker/volumes/nginx/ssl/localca.pem`

## Build docker images
```
cd docker
./exec build
```

## Run services
```
./exec restart
```

## Migrate database
```
./exec migrate
```

## Seed database
```
./exec command initdata
```

## Seed database
```
./exec command syncrolespems
```

## Update hosts files
Add this line to hosts file
```
127.0.0.1       simplepm.test
```

## Run backend server
```
./exec bserver
```

## Run frontend server
```
./exec fserver
```
## Build backend code
```
./exec bbuild
```

## Build frontend code
```
./exec yarn build
```

## Make migrations
```
./exec makemigrations
```


