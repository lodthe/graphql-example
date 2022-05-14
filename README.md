# GraphQL Example

A simple Golang application that implements GraphQL API and stores data in PostgreSQL.

GraphQL models are generated using [99designs/gqlgen](github.com/99designs/gqlgen).

## Startup

You can run both the server and the storage via `docker-compose up -d`. 
After that, the GraphiQL interface will be available at localhost:80.

## API

The full API schema is available [here](./api/schema.graphqls).

### Query
```graphql
query {
  # Select one match by id
  match(id: "371b9bd1-1bc3-4da8-819c-6780b5096061") {
    id
    comments {
      id
      text
    }
  }
    
  # Filter existing matches
  matches(isFinished: true, limit: 2, offset: 5) {
    id
    createdAt
    isFinished
    comments {
      id
      text
    }
    scoreboard {
      players {
        username
        role
        isAlive
        kills
      }
    }
  }
}
```

### Mutation

```graphql
mutation {
  # Add a comment to an existing match
  createComment(matchId: "dfc2a88c-a41f-4618-a864-f73d0fcc19e4", text: "blah blah") {
    id,
    text
  }
}
```