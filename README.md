# wizeline-go-bootcamp

- This project is a simple user service. 
- Tech stacks: Golang, Postgres (Docker)
- I used golang-migrate to migrate db
- I used echo framework for providing routers
- I used sqlc for interact with DB

## How to run:
- View Make file for commands 
- If you want to run this project:
    - make db //NOTE: to create a container for db
    - make createdb //NOTE: create a db in the above container
    - make miup //NOTE: migrate db 
    - make server //NOTE: run the server 
    - then use Postman to run:
        - GET: localhost:8000 //NOTE: get Hello, World!
        - POST: localhost:8000/users with json //NOTE: create a new user
            {
                "username": "anything_you_like",
                "email": "anything_you_like@email.com"
            }
        - GET: localhost:8000/anything_you_like //NOTE: get created user 
        