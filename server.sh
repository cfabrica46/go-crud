#!/bin/bash

#Signin
response=$(curl -sb -X POST http://localhost:8080/api/v1/signin -d '{"username":"cfabrica46","password":"01234"}')

token=$(echo "$response" | jq -r '.content')

#SignUp
#response=$(curl -X POST http://localhost:8080/api/v1/signup -d '{"username":"cfabrica46","password":"01234","email":"cfabrica46@gmail.com"}')

#ShowUsers
# curl -X GET http://localhost:8080/api/v1/users

#Profile
#curl -X GET http://localhost:8080/api/v1/user -H "Authorization: $token"

#Delete
# curl -X DELETE http://localhost:8080/api/v1/user -H "Authorization: $token"
