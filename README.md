# Introduction
This database is designed for high load situations and has a main architecture of Redis and PostgreSQL. Due to the need to generate a large number of personal lists at once, the personal list in PostgreSQL is designed using the List data structure. We first check if the head of the ID -> ListKey exists in Redis. If there is such a list, we update the ID -> ListKey in Redis. When a user visits our database, we output ten pieces of data to them in a lazy tag-like manner, which ensures the throughput performance of the database under heavy traffic.

# Endpoints
```/modify/{uid}```

This endpoint allows users to input data into the database. The request should be sent as a POST request with the following parameters:

- {uid}: the unique identifier for the user
- {content}: an array of data to be inputted into the database

### Example request:

```bash
POST /modify/12345
{
    "content": ["data1", "data2"]
}
```

```/get/{uid}/{page}```

This endpoint allows users to retrieve data from the database. The request should be sent as a GET request with the following parameters:

- {uid}: the unique identifier for the user
- {page}: the page number of the data to be retrieved

Example request:
```
GET /get/12345/2
```

## Port
The port number for this database is 8006. Please make changes to main.go accordingly.

## Deployment
Please refer to the docker-compose.yml file for deployment information. To start the program, run docker-compose up followed by go run main.go. To stop the program, use kill -9 PID.