# cassandra
Cassandra Installation on K8s:    
https://kubernetes.io/docs/tutorials/stateful-application/cassandra/
```
kubectl apply -f persistent.yaml
kubectl apply -f deployment.yaml
kubectl apply -f svc.yaml

```

# access via cqlsh via service port

```
root@node1:/data/cassandra# cqlsh 10.233.8.211 31631
Connected to K8Demo at 10.233.8.211:31631.
[cqlsh 5.0.1 | Cassandra 3.11.2 | CQL spec 3.4.4 | Native protocol v4]
Use HELP for help.
cqlsh>
```
# use cql syntax (sql alike)
```
CREATE KEYSPACE my_keyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE my_keyspace;

CREATE TABLE my_table (
    id UUID PRIMARY KEY,
    name text,
    age int
);

INSERT INTO my_table (id, name, age) VALUES (uuid(), 'John', 25);
INSERT INTO my_table (id, name, age) VALUES (uuid(), 'Jane', 30);

SELECT * FROM my_table;
```

# access via client
```
python3 client.py
```
