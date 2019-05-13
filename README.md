
# gin-gorm-sample

## How To Run

```
docker-compose up
```

or

```
docker-compose run -v $PWD:/app -p 8080:8080 app bash
ENV=development go run cmd/db/migrate/main.go
ENV=development go run cmd/app/main.go
```


## Query

```
mutation createTodo {
  createTodo(input:{text: "first", userId: 1}) {
    id
    text
    createdAt
    updatedAt
  }
}

mutation createUser {
  createUser(input:{name: "first"}) {
    id
    name
    createdAt
    updatedAt
  }
}

mutation updateTodo {
  updateTodo(input:{id: 4, text: "hogehoge"}) {
    text
    done
    createdAt
    updatedAt
  }
}

mutation deleteTodo {
  deleteTodo(input: 1) {
    text
    done
  }
}


query findTodos {
  	todos {
      text
      done
      userID
      user {
        id
        name
      }
    }
}

query findTodo {
  	todo(input: { id: 2}) {
      text
      done
      userID
      user {
        id
        name
      }
    }
}

query findUsers {
  	users {
      name
    }
}
```
