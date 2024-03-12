# Match Words

The objective of the app is that users can create an account and gain points for matching words with their definitions. 

When they enter to the app, they can see the top scores and the option to play.

# Technical Info

This project uses microservices architecture.

Project's backend is developed with Golang, and project's frontend is developed using React (with tailwind).

To run the app you'd need to use Kubernetes or docker.

# Task List

- [] Auth Service.
- [] Gateway Service 
- [] User Service
- [] Words Service
- [] Client
- [] Logs Service
- []
- []
- []
- []
- []

# Services logic

### Auth

This service manages the authorizations of the backend. If your client or other source makes a request to a protected route it will ask for a token and validate it.

Otherwise if the client wants to get a token, he needs to login and the token will be sent to him in the response.

If user doesn't exist, it can register to the database (if user registries, it will send a request to user's service, because it needs to save the data there too).