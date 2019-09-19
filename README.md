# Testify-handlers

Toolkit for testing http handlers in a verbose way.

```go
func TestAPIGet(t *testing.T) {
  // create your request 
  handlertest.Call(yourHandler).GET("/jobs").
    // then assert your expectations
    Assert(t).
      Status(http.StatusOK).          
      Header("Allow-Origin: *").
      JsonBody(`[{"id": 1}]`).
      JsonUnmarshallsTo([]*models.Job)
}

// or

func TestPostForm(t *testing T) {
  handlertest.Call(yourHandler).FormUrlEncodedMap(map[string]string{
    "field": "value"
  }).Assert(t).
    Status(http.StatusCreated).   .
    ContentType("text/html")
} 
```

## Semver

This library will follow [semantic versioning](https://semver.org/) and next versions will be released as tags.

For now it is in alpha/exploratory phase. v1 will be released after community feedback is sourced.


## Request

Test request is created with `handlertest.Call(YourHttpHandler)`. 

Then you can set how the request should look like by chaining methods. 
```
handlertest.Call(YourHttpHandler).POST("/jobs").Json(`{"name": "test"}`)
```

Methods below are part of `Request`.

#### Methods

- `Method("PUT")` - Call handler with a given method.
- `POST()` - shorthand for a POST method
- `GET()` - shorthand for a GET method

#### Url

- `Url("/jobs?status=SUCCESSFUL` - set request's URL

#### Body

- ```Json(`{"name": "test"}`)``` - sets given json as body and add `Content-Type: application/json` header

TODO custom body

#### Forms

A set of methods that encodes form values in a body. 
Methods creating forms of `Content-Type: application/x-www-form-urlencoded` 
- `FormUrlEncoded(values url.Values)` - general method to set form fields
- `FormUrlEncodedMap(values map[string]string)` - a shorthand accepting single strings as field values

Methods creating forms of `Content-Type: multipart/form-data`:
- `FormMultipart(fields url.Values)` - general method to set form fields that are not files
- `FormMultipartMap(values map[string]string)` - shorthand for the above
- `File(field string, fileName string, content string)` - adds a file to a given field
- `FileReader(field string, fileName string, content io.Reader)` - adds a file to a given field taking content from a reader
- `Files(fields map[string]map[string]string)` sets all files at once
- `FileReaders(fields map[string]map[string]io.Reader)` - sets all files at once

#### Headers

- `Header(key string, value string)` - sets a header
- `ContentType(contentType string)` - shorthand for setting `Content-Type`

#### Context

*TBDone*

#### Custom

To implement custom modifications of response not covered by this lib use the below function. Please also file an issue to support your case if it is general enough.
- `Custom(func(request *http.Request))`

## Assert 

Once you created a needed request call on it `.Assert(t)` to get get an object where you can specify assertions.

```go
func TestPostForm(t *testing T) {
  handlertest.Call(yourHandler).FormUrlEncodedMap(map[string]string{
    "field": "value"
  }).Assert(t).
    Status(http.StatusCreated).   .
    ContentType("text/html")
} 
```

#### Status

- `Status(statusCode int)` - assert that response has specific HTTP Status Code

#### Headers

- `Header(key string, value string)` - assert that specific header is set
- `HeaderMissing(key string)` - assert that specific header is not set
- `ContentType(contentType string)` - assert that response is of specific `Content-Type`

#### Body

There is one general function that lets you assert that handler provided a specific body:
- `Body(func(t *testing.T, body []byte))`
		
To support asserting for json response common in API development there are following assertions that tests that Content-Type is set right and offer different ways to assert for the body contents: 
- ```JsonBody(`[{"id": 1}]`)``` - providing it as a string. Indentation here doesn't play a role and there will be an option to show diff between expected and actual values.
- `JsonUnmarshallsTo([]Obj{})` - there is simple assertion that tests unmarshalling 
- `JsonMatches(func(t *testing.T, ret []Obj)` - in case it's hard to predict your whole response, 
or you don't want to test the whole response, you might use this function 
to get an unmarshalled response body and test for specific values.

#### Context

#### Custom

To implement custom assertions you might need that is not covered by this lib use below function. Please also file an issue to support your case if it is general enough
```go
a.Custom(func(t *testing.T, response *http.Response) {
    t.Error("I don't like this response")
})
```
## Examples

### Get a list of filtered objects

```go
func TestListFilter(t *testing.T) {
  // TODO to some inserts to test DB
  // create your request 
  handlertest.Call(yourHandler).GET("/products?category=a").
    // then assert your expectations
    Assert(t).
      Status(http.StatusOK).          
      JsonMatches(func(t *testing.T, products []Product) {
        // unmarshalling of JSON objects is done for you
        if len(products) == 0 {
          t.Errorf("Expected to have some products returned")
        }   
        _, p := range products {
          if p.category != "a" {
            t.Errorf("Expected filter to return products only of category %s, but got %s", "a", p.category)
          }       
        } 
      })
}
```

### File upload

```go
func TestUploadAttachments(t *testing T) {
	// create request
    handlertest.Call(blog.uploadAttachments).
        POST("/attachments").
        FormMultipartMap(map[string]string{
            "post_id": "1"
        }).
        File("files[]", "img1.jpg", "contents").
        // then assert your expectations
        Assert(t).
        Status(http.StatusCreated).   .
        ContentType("text/html")
} 
```