## Sample grapphQL go requests

### Create User
`curl X POST -H 'Content-Type: application/json' -d '{"query":"mutation {createUser(id:\"Test1\",userrole:\"driver\",latitude:12.554,longitude:35.22) { id,userrole,latitude,longitude } }"}'  http://localhost:8080/v1`

### Get user location
`curl -g 'http://localhost:8080/v1?query={user(id:"Test1",userRole:"driver"){latitude,longitude,type}}'`

### Find Nearby
`curl -g 'http://localhost:8080/v1?query={nearbyUser(latitude:14.554,longitude:44.56){latitude,longitude,id}}'`

### Set user location (update location)
`curl -X POST -H 'Content-Tpe: application/json' -d '{"query": "mutation { setLocation(id:\"Test1\", userrole:\"driver\",latitude:12.556,longitude:35.24)}"}'  http://localhost:8080/v1`