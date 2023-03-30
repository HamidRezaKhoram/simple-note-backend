# Go API with Gorm and Postgres
###This repository is the backend of our [Simple-note](https://simple-note.hamidrezakhoramrokh.ir/)
This is a Go API using [Gorm](https://gorm.io/) as an ORM and [Postgres](https://www.postgresql.org/) as a database.
## Developers
This API was developed by [Daniel Ibanescu](https://github.com/zazu7765) and [Hamidreza Khoramrokh](https://github.com/HamidRezaKhoram/).
## Installation

1. Clone this repository.
2. Install the dependencies using `go mod download`.
3. Set the environment variables in `.env.example` and rename it to `.env`.
4. Run the server using `go run main.go`.

## Models

### User
- `Name` (string): the name of the user.
- `Email` (string): the email address of the user.
- `Password` (string): the password of the user.
- `Notebooks` (array of Notebook): the notebooks owned by the user.

### Notebook
- `UserID` (uint): the ID of the user who owns the notebook.
- `Title` (string): the title of the notebook.
- `Description` (string): the description of the notebook.
- `Notes` (array of Note): the notes in the notebook.

### Note
- `NotebookID` (uint): the ID of the notebook that the note belongs to.
- `Title` (string): the title of the note.
- `Content` (string): the content of the note.

## Migrations

The `migrateAll` function in `main.go` is used to migrate the database. It migrates the User, Notebook, and Note models.

## API Endpoints

### Authentication

- POST `/auth/signup`: create a new user account.
- POST `/auth/login`: log in to an existing user account.

### User

- GET `/user`: get the authenticated user's information.
- PUT `/user`: update the authenticated user's information.
- DELETE `/user`: delete the authenticated user's account.

### Notebook

- GET `/notebooks`: get all notebooks owned by the authenticated user.
- GET `/notebooks/all`: get all notebooks.
- POST `/notebooks`: create a new notebook.
- PUT `/notebooks`: update a notebook.
- DELETE `/notebooks`: delete a notebook.

### Note

- GET `/notes/:id`: get a specific note.
- GET `/notes/all`: get all notes in a notebook.
- POST `/notes`: create a new note.
- PUT `/notes/:id`: update a note.
- DELETE `/notes/:id`: delete a note.

### Public

- GET `/`: a public endpoint that always returns success.

Note: All endpoints except for `/public` require authentication using a JSON Web Token (JWT) in the `Authorization` header.

