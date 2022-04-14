package main
import (
    "testing"
   // "github.com/stretchr/testify/assert"
    //"github.com/stretchr/testify/require"
   // "os"
    //"net"
    "net/http"
    "net/http/httptest"
    "github.com/go-chi/chi/v5/middleware"
   // "fmt"
   // "syscall"
   // "io/ioutil"
)


func Test_Main(t *testing.T) {
    testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("[]"))
      })

      req := httptest.NewRequest("GET", "/ping", nil)
      rr := httptest.NewRecorder()

      handler := middleware.DefaultLogger(testHandler)
      handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
    }
//
//     got, _ := ioutil.ReadAll(rr.Body)
//     fmt.Printf("%s",string(got))
// //     want := "pong"
// //
// //     if got != want {
// //         t.Errorf("got %q, want %q", got, want)
// //     }
}