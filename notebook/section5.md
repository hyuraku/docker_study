- docker network ls
  show all docker network
- docker network create [name]
  create a new docker network
- docker network inspect [name]
  show info of one docker network
- docker container run -d --name new_nginx --network [name] nginx
  run docker container with assigned docker network network
- docker network connect [network id] [container id]
  connect certain contaner to network