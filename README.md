# Subscription Management API
Tech Stack: Go (Fiber), MongoDB

Description: A simple web application that provides basic Create, Read, Update, and Delete operations for managing users. Built with the Fiber web framework and connected to a MongoDB database, 
this API allows you to perform CRUD operations on a user model, including adding new users, retrieving user details, updating existing user information, and deleting users.

## API Endpoints

| Method | Route       | Action                         |
|--------|-------------|--------------------------------|
| GET    | /users      | Get the list of users          |
| GET    | /user/:id   | Get a particular user          |
| POST   | /user       | Create a new user              |
| POST   | /user/:id   | Update an existing user info   |
| DELETE | /user/:id   | Delete an existing user        |

## Usage
**Note: Install make to run the following commands**
1. Install dependencies:
```
make deps
```
2. Build the project:
```
make build
```
3. Run the project:
```
make run
```
4. Clean Build Artifacts(Optional):
```
make clean
```

## Check List
- [x] Models
- [x] Routes
- [x] Controllers
- [x] Workflows
- [ ] Test files(To be added in the future or help me if you can)
