# dgraph-recommender
Setting up and running Dgraph requires us to use Docker container. After installation run the following commands:

dgraph zero:
>docker run -it -p 5080:5080 -p 6080:6080 -p 8080:8080 -p 9080:9080 -p 8000:8000 -v ~/dgraph:/dgraph --name dgraph dgraph/dgraph dgraph zero

dgraph:
>docker exec -it dgraph dgraph server --lru_mb 2048 --zero localhost:5080

ratel ui:
>docker exec -it dgraph dgraph-ratel

To access the web ui visit: http://localhost:8000/

To load data into our graph it is more ideal to use Bulk load. However, in case we need to use Live load the command is:
>  dgraph live -r <path-to-rdf-gzipped-file> -s <path-to-schema-file> -d <dgraph-server->address:grpc_port> -z <dgraph-zero-address:grpc_port>

Discussion page with support engineers from Dgraph: https://discuss.dgraph.io/t/recommendation-engine-tutorial-help/2778/19

useful docker commands:

description:
> docker ps -a

to re-start container:
> docker start container_name
