CREATE KEYSPACE IF NOT EXISTS telemetry
WITH replication = {
   'class': 'SimpleStrategy',
   'replication_factor': 1
};

DROP TABLE collection_config;
CREATE TABLE IF NOT EXISTS collection_config (
            sysid UUID PRIMARY KEY,
            id text,
			name text,			
            interval int,
			latency int,
			metric text,
			host text,
			tsp TIMESTAMP,
			qry text,
			updated_by text
        );
