package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/valve"
	customMiddleware "github.com/sgulics/go-chi-example/pkg/middleware"
	"github.com/sgulics/go-chi-example/pkg/routes"
	"github.com/sgulics/go-chi-example/pkg/services"
	"github.com/sgulics/go-chi-example/pkg/stores"
	"github.com/sgulics/go-chi-example/pkg/templating"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)


func main() {

	valv := valve.New()
	baseCtx := valv.Context()

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		// disable, as we set our own
		DisableTimestamp: false,
	}

	//tm := template_manager.NewTemplateManager("templates_oild/layouts/", "templates_oild/")
	tm, err := templating.NewTemplateManager("templates/", "templates/layouts", ".gohtml", true)
	//err := tm.LoadTemplates()
	//err := tm.FindAndParseTemplates()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	// RequestID is a middleware that injects a request ID into the context of each
	// request. A request ID is a string of the form "host.example.com/random-0001",
	// where "random" is a base62 random string that uniquely identifies this go
	// process, and where the last number is an atomically incremented request
	// counter.
	r.Use(middleware.RequestID)
	r.Use(customMiddleware.NewStructuredLogger(logger))
	// Recoverer is a middleware that recovers from panics, logs the panic (and a
	// backtrace), and returns a HTTP 500 (Internal Server Error) status if
	// possible. Recoverer prints a request ID if one is provided.
	r.Use(middleware.Recoverer)
	// URLFormat is a middleware that parses the url extension from a request path and stores it
	// on the context as a string under the key `middleware.URLFormatCtxKey`. The middleware will
	// trim the suffix from the routing path and continue routing.
	//r.Use(middleware.URLFormat)
	//r.Use(render.SetContentType(render.ContentTypeJSON))
	service := services.NewArticlesService(stores.NewMemoryStore(), logger)

	r.Mount("/admin", routes.AdminRoutes(tm))
	r.Mount("/monitors", routes.MonitorRoutes())
	r.Mount("/v1/articles", routes.NewArticleResource(service).ArticleRoutes())


	srv := http.Server{Addr: ":3333", Handler: chi.ServerBaseContext(baseCtx, r)}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			fmt.Println("shutting down..")

			// first valv
			valv.Shutdown(20 * time.Second)

			// create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			// start http shutdown
			srv.Shutdown(ctx)

			// verify, in worst case call cancel via defer
			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()
	srv.ListenAndServe()
}

