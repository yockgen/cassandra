# Accessing Cassandra via Golang client

Configure:
```
hspe@BA38RDL00616:/data/cassandra/go-client$ cat config.yaml
cassandra:
  host: 10.49.76.106
  port: 30200
  username: admin
  password: intel@2023
```


Run this:
```
hspe@BA38RDL00616:/data/cassandra/go-client$ go run . "select id,name,metric,host from collection_config;"
```
