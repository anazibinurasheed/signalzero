Assessment
===

Implement 2 API endpoints using Gin framework and the endpoints should store and retrieve data using MongoDB.
## Requirments



### POST API

- Send the user details such as user_id, username, password in the request body.
- Recieve it from the backend, then insert it to a collection.


##### request
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

- Retreive username as query param from the request.
- If query is empty show all users.
- If query exists then retrieve all the users who have the name with same prefix. 


##### request

```
curl --location 'http://localhost:8080?username=new'
```



### Tested Platform
  ```
  OS: Ubuntu LTS
  Go: version go1.22.1 linux/amd64
  ```
