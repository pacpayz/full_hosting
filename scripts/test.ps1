docker build -t test_full_hosting .

docker run -it -p 8080:8080 --name full_hosting_container test_full_hosting

docker rm full_hosting_container

docker rmi test_full_hosting