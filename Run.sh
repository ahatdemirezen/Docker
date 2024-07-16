docker image build -f Dockerfile -t ascii .
docker container run -p:8080:8080 --detach --name asciicontainer ascii
docker ps -a
echo To stop container, run "Docker stop <container_name>".