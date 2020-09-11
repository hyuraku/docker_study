- docker-compose
```
docker-compose build .
docker-compose run -p 2224 app ./main

docker ps 
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                     NAMES
hogehoge3        golang_app          "./main"                 hogehoge seconds ago       Up hogehoge seconds        0.0.0.0:32769->2224/tcp   hogehoge

ssh -p 32769 localhost
```

- docker
```
docker build -t golang:0.2
docker run -it --rm -p 2224 --name test3 golang:0.2

 docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                     NAMES
hogehoge        golang:0.2          "./main"                 About hogehoge minute ago   Up About hogehoge minute   0.0.0.0:32770-->2224/tcp   test3

ssh -p 32770 localhost
```

To bypass this error, regnerate your host key with
`ssh-keygen -R "[localhost]:2222"`