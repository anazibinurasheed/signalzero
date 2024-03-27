Assessment
===

Implement 2 API endpoints using the Gin framework, and the endpoints should store and retrieve data from MongoDB.

## Requirements

### POST API

- Send the user details such as user_id, username, password in the request body.
- Receive it on the backend, then insert it into a MongoDB collection.

##### Request

```
curl --location 'http://localhost:8080/' \
--header 'Content-Type: application/json' \
--data '{

"user_id":"1",
"username":"new person",
"password":"1234"
}'
```
### GET API

- Retrieve username as a query parameter from the request.
- If the query is empty, show all users stored in the 'users' collection.
- If the query exists, retrieve all the users whose names start with the specified prefix.
- Give response with the retrieved user information.

##### Request

```
curl --location 'http://localhost:8080?username=new'
```



### Tested Platform
  ```
  OS: Ubuntu LTS
  Go: version go1.22.1 linux/amd64
  ```
