package main
import (
  "log"
   "net/http"
   "github.com/pkg/errors"
   "github.com/go-chi/chi/v5"
   "github.com/go-chi/chi/v5/middleware"
   "github.com/jessevdk/go-flags"
   "strings"
)

type Server struct {
	PinSize        int
	MaxPinAttempts int
	WebRoot        string
	Version        string
	Port           string
}

type Options struct {
    Port string `short:"p" long:"port" env:"SERVER_PORT" default:"8080" description:"Port web server"`
}

func main() {
    var opts Options
    parser := flags.NewParser(&opts, flags.Default)
    _, err := parser.Parse()
    if err != nil {
        log.Fatal(err)
    }

    srv := Server {
        PinSize:   1,
        WebRoot:   "/",
        Version:   "1.0",
        Port: opts.Port,
    }

    if err := srv.Run(); err != nil {
        log.Printf("[ERROR] failed, %+v", err)
    }
}

func (s Server) Run() error {
    log.Printf("[INFO] Activate rest server")
    log.Printf("[INFO] Port: %s", s.Port)

	if err := http.ListenAndServe(":"+s.Port, s.routes()); err != http.ErrServerClosed {
		return errors.Wrap(err, "server failed")
	}

	return nil
}

func (s Server) routes() chi.Router {
	router := chi.NewRouter()

    router.Use(middleware.Logger)
    router.Use(Ping)
    router.Route("/", func(r chi.Router) {
    })

	return router
}

func Ping(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && strings.HasSuffix(strings.ToLower(r.URL.Path), "/ping") {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("pong"))
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}





