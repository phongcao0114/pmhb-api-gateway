# pmhb-api-gateway
An implementation of GraphQL in Go.
This is an example which defines GraphQL queries for basic CRUD functions in pmhb-book-service with the resulting output printed in JSON format.

## Getting Started
```bash
go get https://github.com/phongcao0114/pmhb-api-gateway
```
```bash
go get https://github.com/phongcao0114/pmhb-book-service
```

## Running
1. Import Database: ./pmhb-book-service/book.sql
2. Run pmhb-book-service
3. Run pmhb-api-gateway
4. Access: 
http://localhost:10000/

## Queries

####Login:
Query:

	mutation ($username: String!,$password: String!){
		login(
			username: $username
			password: $password
		)
	}
Query variables:

	{
		"username": "user1",
		"password": "P@ssw0rd"
	}
Sample response:

	{
		"data": {
			"login": <token>
		}
	}

####Get book list:
Query:

	query{
		books{
			id,
			name,
			author
		}
	}

HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}
####Get book by id:
Query:
Query variables:
HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}
####Create a new book:
Query:
Query variables:
HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}
####Update a book:
Query:
Query variables:
HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}
####Delete a book:
Query:
Query variables:
HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}

