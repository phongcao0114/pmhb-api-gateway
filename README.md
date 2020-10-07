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

###Login:
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

###Get book list:
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
###Get book by id:
Query:

	query($id: String!) {
		book_by_id(id:$id) {
			id
			name
		}
	}

Query variables:

	{
		"id": "cc4116db-638c-4ae1-b9d9-208b6fd15391"
	}
HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}
###Create a new book:
Query:

	mutation ($name: String!,$author: String!){
		create_book(
			name:$name
			author:$author
		)
	}
Query variables:

	{
		"name": "Book1",
		"author":"B.A"
	}
HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}
###Update a book:
Query:

	mutation ($id: String,$author: String){
		update_book(
			id: $id
			author:$author
		)
	}
Query variables:

	{
		"id": "169349ad-035a-4873-bc5b-012195d4a984"
		"author":"X.B"
	}
HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}
###Delete a book:
Query:

	mutation($id: String!){
		delete_book(id:$id)
	}
Query variables:

	{
		"id": "169349ad-035a-4873-bc5b-012195d4a984"
	}
HTTP Headers:

	{
		"Authorization":"Bearer <token>"
	}
