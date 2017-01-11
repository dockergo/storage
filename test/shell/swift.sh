#!/bin/bash


#keystone
curl -i http://192.168.20.6:5000/v3/auth/tokens -X GET -s -H "Content-Type: application/json" -d "
        {
            \"auth\": {
                \"identity\": {
                    \"methods\": [
                        \"password\"
                    ],
                    \"password\": {
                        \"user\": {
                            \"domain\": {
                                \"name\": \"test\"
                            },
                            \"name\": \"wps\",
                            \"password\": \"12345\"
                        }
                    }
                },
                \"scope\": {
                    \"project\":{
                        \"domain\": {
                            \"name\": \"test\"
                        },
                        \"name\": \"wps\"
                    }
                }
            }
        }"


#account
curl -i $publicURL?format=json -X GET -H "X-Auth-Token: $token"

curl -i $publicURL -X POST -H "X-Auth-Token: $token" -H "X-Account-Meta-Book: MobyDick" -H "X-Account-Meta-Subject: Literature"
curl -i $publicURL -X POST -H "X-Auth-Token: $token" -H "X-Account-Meta-Subject: AmericanLiterature"
curl -i $publicURL -X POST -H "X-Auth-Token: $token" -H "X-Remove-Account-Meta-Subject: x"

curl -i $publicURL -X HEAD -H "X-Auth-Token: $token"


#container
curl -i $publicURL/steven -X PUT -H "Content-Length: 0" -H "X-Auth-Token: $token"
curl -i $publicURL/marktwain -X POST
-H "X-Auth-Token: $token"
-H "X-Container-Meta-Author: MarkTwain"
-H "X-Container-Meta-Web-Directory-Type: text/directory"
-H "X-Container-Meta-Century: Nineteenth"
curl -i $publicURL/marktwain -X HEAD -H "X-Auth-Token: $token"
curl -i $publicURL/steven -X DELETE -H "X-Auth-Token: $token"


#object
curl -i $publicURL/marktwain/goodbye -X GET -H "X-Auth-Token: $token"
curl -i $publicURL/janeausten/goodbye -X GET -H "X-Auth-Token: $token"
curl -i $publicURL/janeausten/helloworld.txt -X PUT -d "Hello" -H "Content-Type: text/html; charset=UTF-8" -H "X-Auth-Token: $token"

curl -i $publicURL/marktwain/goodbye -X COPY -H "X-Auth-Token: $token" -H "Destination: janeausten/goodbye"
curl -i $publicURL/janeausten/goodbye -X PUT -H "X-Auth-Token: $token" -H "X-Copy-From: /marktwain/goodbye" -H "Content-Length: 0"

curl -i $publicURL/marktwain/helloworld -X DELETE -H "X-Auth-Token: $token"

curl $publicURL/marktwain/goodbye --head -H "X-Auth-Token: $token"

curl -i $publicURL/marktwain/goodbye -X POST -H "X-Auth-Token: $token" -H "X-Object-Meta-Book: GoodbyeColumbus"
curl -i $publicURL/marktwain/goodbye -X POST -H "X-Auth-Token: $token" -H "X-Object-Meta-Book: GoodbyeOldFriend"

curl -i -http://10.1.0.220/v1/AUTH_7c9c920f39de4fd8a9bdf807a3f11b5d/container1/bbb -X PUT -T "storage.toml" -H "X-Auth-Token: 1267a227d9994da4a9ea4977a5f5ee04"
