<!-- .slide: class="with-code-bg-dark" -->

# Tests HTTP

## Les tests end-to-end avec httptest.Server

```go
func TestEndToEnd(t *testing.T) {
    ts := httptest.NewServer(
        http.HandlerFunc(
            func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "Hello, client")
            }
        )
    )
    defer ts.Close()
    res, err := http.Get(ts.URL)
    if err != nil {
        t.Error(err)
    }
    // ...
}
```

Notes:
OFU

Negroni implémente http.Handler
