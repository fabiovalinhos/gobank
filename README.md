# GoBank
GoBank is a API make in Golang using Docker container Postgres as database  
This api was create following a tutorial by Anthony GG from Youtube

Command used for create Docker container Postgres  
For create Docker Postgres image, were used this default command from Postgres' documentation 
https://hub.docker.com/_/postgres

> docker run --name some-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres

For rerun image
docker start some-postgres(nome da image)