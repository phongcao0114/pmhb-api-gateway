# pmhb-api-gateway
Access: localhost:10000


`Create a new book:`

mutation{
  create_book(
    name:"book004"
    author:"A.U"
  )
}

`Update book by id:`

mutation{
  update_book(
    id: "3527285a-64e7-46b3-95d6-2b9d849b6504"
    author:"B.C"
  )
}

`Delete book:`

mutation{
  delete_book(id:"114d4f44-af05-48d0-8650-41406619958b")
}

`Create a new book:`

mutation{
  create_book(
    name:"book004"
    author:"A.U"
  )
}

`Get book list:`

query{
  books{
    id,
    name,
    author
  }
}

`Get book by id:`

query{
  book_by_id(
    id:"045ce937-2122-4804-8d05-e30f6e07acfc"
  ){
    id
    name
    author
  }
}