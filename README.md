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

## Documentation
TODO