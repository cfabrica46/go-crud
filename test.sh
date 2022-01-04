#!/bin/bash

#Signin
# response=$(curl -sb -X POST http://localhost:8080/api/v1/signin -d '{"username":"cfabrica46","password":"01234"}')

#SignUp
# response=$(curl -X POST https://localhost:8081/api/v1/signup -d '{"username":"cfabrica46","password":"01234","email":"cfabrica46@gmail.com"}')

# token=$(echo "$response" | jq -r '.content')


#ShowUsers
curl -X GET https://localhost:8081/api/v1/users

#Profile
# curl -X GET https://localhost:8081/api/v1/user -H "Authorization: $token"

#Delete
# curl -X DELETE http://localhost:8080/api/v1/user -H "Authorization: $token"
