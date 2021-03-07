# xavier

A client/server application which stores information passed via REST API and can be recalled via the CLI Client.

## Design

### Software Stack

- Server: Golang / SQLite (MVP)
- Client: Golang

### Server

- REST Server which receives input from a source via HTTP Calls, stores in the appropriate schema and enables the CLI to pull the data.

### CLI

- Displays input based on query provided to the server in a user-friendly version
