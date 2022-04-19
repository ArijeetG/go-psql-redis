Hi

stacks used: 
    - golang
    = postgreSql
    - docker
    - redis


routes: 
    - /signup  -->
            method: POST
            body (JSON): {
                "Name" : <string>,
                "Password: : <string>,
                "User_type" : <string> (any arbitraty value)
            }
    - /login  -->
            method: POST
            body (JSON) : {
                "Name" : <string>,
                "Password" : <string>
            }
    - /hello  -->
            description: View details for a user against a name
            method: POST
            body (JSON) : {
                "Name" : <string>
            }
    
    - /editInfo -->
            description: Edit details for a user agains a name
            method: POST
            body (JSON) : {
                "Name" : <string>,
                "Address" : <string>,
                "Phone" : <string>
            }
    - /delInfo --> 
            description: Delete details for a user against a name
            method: POST,
            body (JSON) : {
                "Name": <string>,
                "Password": <password>
            }

setup:
    -> clone the repository
    -> go to /database and run docker compose up
    -> go to /redis and run docker compose up
    -> go to /server and run docker compose up
    -> endpoints will be accessible at http://localhost:4000/
