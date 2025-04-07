# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [authors.proto](#authors-proto)
    - [Author](#authors-Author)
    - [AuthorList](#authors-AuthorList)
    - [AuthorRequest](#authors-AuthorRequest)
    - [CreateAuthorRequest](#authors-CreateAuthorRequest)
    - [CreateAuthorResponse](#authors-CreateAuthorResponse)
    - [Empty](#authors-Empty)
    - [StatusResponse](#authors-StatusResponse)
    - [UpdateAuthorRequest](#authors-UpdateAuthorRequest)
  
    - [AuthorService](#authors-AuthorService)
  
- [books.proto](#books-proto)
    - [Book](#books-Book)
    - [BookList](#books-BookList)
    - [BookRequest](#books-BookRequest)
    - [CreateBookRequest](#books-CreateBookRequest)
    - [CreateBookResponse](#books-CreateBookResponse)
    - [Empty](#books-Empty)
    - [StatusResponse](#books-StatusResponse)
    - [UpdateBookRequest](#books-UpdateBookRequest)
  
    - [BookService](#books-BookService)
  
- [combine.proto](#combine-proto)
    - [AuthorAndBook](#combine_book_and_author_service-AuthorAndBook)
    - [CombinedUpdateRequest](#combine_book_and_author_service-CombinedUpdateRequest)
  
- [Scalar Value Types](#scalar-value-types)



<a name="authors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## authors.proto



<a name="authors-Author"></a>

### Author



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name_author | [string](#string) |  |  |
| surname_author | [string](#string) |  |  |
| biography | [string](#string) |  |  |
| birthday | [string](#string) |  |  |






<a name="authors-AuthorList"></a>

### AuthorList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Authors | [Author](#authors-Author) | repeated |  |






<a name="authors-AuthorRequest"></a>

### AuthorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="authors-CreateAuthorRequest"></a>

### CreateAuthorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name_author | [string](#string) |  |  |
| surname_author | [string](#string) |  |  |
| biography | [string](#string) |  |  |
| birthday | [string](#string) |  |  |






<a name="authors-CreateAuthorResponse"></a>

### CreateAuthorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| author | [Author](#authors-Author) |  |  |






<a name="authors-Empty"></a>

### Empty







<a name="authors-StatusResponse"></a>

### StatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="authors-UpdateAuthorRequest"></a>

### UpdateAuthorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name_author | [string](#string) |  |  |
| surname_author | [string](#string) |  |  |
| biography | [string](#string) |  |  |
| birthday | [string](#string) |  |  |





 

 

 


<a name="authors-AuthorService"></a>

### AuthorService

Метод для получения книг

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAllAuthors | [Empty](#authors-Empty) | [AuthorList](#authors-AuthorList) |  |
| GetAuthorById | [AuthorRequest](#authors-AuthorRequest) | [Author](#authors-Author) |  |
| CreateAuthor | [CreateAuthorRequest](#authors-CreateAuthorRequest) | [CreateAuthorResponse](#authors-CreateAuthorResponse) |  |
| UpdateAuthor | [UpdateAuthorRequest](#authors-UpdateAuthorRequest) | [StatusResponse](#authors-StatusResponse) |  |
| DeleteAuthor | [AuthorRequest](#authors-AuthorRequest) | [StatusResponse](#authors-StatusResponse) |  |

 



<a name="books-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## books.proto



<a name="books-Book"></a>

### Book



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| author_id | [int32](#int32) |  |  |
| author | [string](#string) |  |  |
| year | [int32](#int32) |  |  |
| isbn | [string](#string) |  |  |






<a name="books-BookList"></a>

### BookList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| books | [Book](#books-Book) | repeated |  |






<a name="books-BookRequest"></a>

### BookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="books-CreateBookRequest"></a>

### CreateBookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| author_id | [int32](#int32) |  |  |
| year | [int32](#int32) |  |  |
| isbn | [string](#string) |  |  |






<a name="books-CreateBookResponse"></a>

### CreateBookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| book | [Book](#books-Book) |  |  |






<a name="books-Empty"></a>

### Empty







<a name="books-StatusResponse"></a>

### StatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="books-UpdateBookRequest"></a>

### UpdateBookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| author_id | [int32](#int32) |  |  |
| year | [int32](#int32) |  |  |
| isbn | [string](#string) |  |  |





 

 

 


<a name="books-BookService"></a>

### BookService

Метод для получения книг

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAllBooks | [Empty](#books-Empty) | [BookList](#books-BookList) |  |
| GetBookById | [BookRequest](#books-BookRequest) | [Book](#books-Book) |  |
| CreateBook | [CreateBookRequest](#books-CreateBookRequest) | [CreateBookResponse](#books-CreateBookResponse) |  |
| UpdateBook | [UpdateBookRequest](#books-UpdateBookRequest) | [StatusResponse](#books-StatusResponse) |  |
| DeleteBook | [BookRequest](#books-BookRequest) | [StatusResponse](#books-StatusResponse) |  |

 



<a name="combine-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## combine.proto



<a name="combine_book_and_author_service-AuthorAndBook"></a>

### AuthorAndBook



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| author | [authors.Author](#authors-Author) |  |  |
| book | [books.Book](#books-Book) |  |  |






<a name="combine_book_and_author_service-CombinedUpdateRequest"></a>

### CombinedUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| book_update | [books.UpdateBookRequest](#books-UpdateBookRequest) |  |  |
| author_update | [authors.UpdateAuthorRequest](#authors-UpdateAuthorRequest) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

