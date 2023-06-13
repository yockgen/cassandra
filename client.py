
from cassandra.cluster import Cluster

def connect_to_cassandra():
    #cluster = Cluster(['10.233.96.108'])  # Replace with Cassandra contact points
    cluster = Cluster(['10.233.8.211'], port=31631)  # Replace with Cassandra contact points
    session = cluster.connect('my_keyspace')  # Replace with your keyspace name
    return session

def create_table(session):
    session.execute(
        """
        CREATE TABLE IF NOT EXISTS my_table (
            id UUID PRIMARY KEY,
            name text,
            age int
        )
        """
    )

def insert_data(session):
    session.execute(
        """
        INSERT INTO my_table (id, name, age) VALUES (uuid(), 'John', 25)
        """
    )
    session.execute(
        """
        INSERT INTO my_table (id, name, age) VALUES (uuid(), 'Jane', 30)
        """
    )

def fetch_data(session):
    rows = session.execute(
        """
        SELECT * FROM my_table
        """
    )
    for row in rows:
        print(f"ID: {row.id}, Name: {row.name}, Age: {row.age}")

def main():
    session = connect_to_cassandra()
    #create_table(session)
    #insert_data(session)
    fetch_data(session)

if __name__ == '__main__':
    main()

