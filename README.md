# Testify-handlers

Toolkit for testing http handlers in a verbose way.

```go
func TestAPIGet(t *testing.T) {
  // create your request 
  handlers.NewRequest(yourHandler).GET("/jobs").
    // then assert your expectations
    Assert().
      Status(http.StatusOK).          
      Header("Allow-Origin: *").
      JsonBody(`[{"id": 1}]`).
      JsonConformsTo([]*models.Job).
      Test(t)
}

// or

func TestPostForm(t *testing T) {
  handlers.NewRequest(yourHandler).FormUrlEncodedMap(map[string]string{
    "field": "value"
  }).Assert().
    Status(http.StatusCreated).   .
    ContentType("text/html").
    Test(t)
} 
```

Using testify-handlers has following advantages:
- offers more flexibility to create request (you can choose to set from values or not, set some headers or not)
- leads to less repeat of code (if you want to test multiple things, such as body, status code, set headers) you will create the request only once
- packs some common testing methods (diff json regardless of indents, set proper headers for form sending, etc.)

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

## Assert 

Once you created a needed request call on it `.Assert()` to get get an object where you can specify assertions.

```go
func TestPostForm(t *testing T) {
  handlertest.Call(yourHandler).FormUrlEncodedMap(map[string]string{
    "field": "value"
  }).Assert().
    Status(http.StatusCreated).   .
    ContentType("text/html").
    Test(t)
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

## Examples

TODO

### Get a list of objects

### Create an object based on form data

### File upload

