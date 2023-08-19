// Package router
package router

import (
	"context"
	"encoding/json"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/v1/endpoint"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/v1/short_url"
	"github.com/aryayunanta-ralali/shorty/pkg/generator"
	"github.com/gorilla/mux"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/bootstrap"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/handler"
	"github.com/aryayunanta-ralali/shorty/internal/middleware"
	"github.com/aryayunanta-ralali/shorty/internal/ucase"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/routerkit"

	//"github.com/aryayunanta-ralali/shorty/pkg/mariadb"
	//"github.com/aryayunanta-ralali/shorty/internal/repositories"
	//"github.com/aryayunanta-ralali/shorty/internal/ucase/example"

	ucaseContract "github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
)

type router struct {
	config *appctx.Config
	router *routerkit.Router
}

// NewRouter initialize new router wil return Router Interface
func NewRouter(cfg *appctx.Config) Router {
	bootstrap.RegistryValidatorRules(cfg)
	bootstrap.RegistryMessage()
	bootstrap.RegistryLogger(cfg)
	bootstrap.RegistrySnowflake()

	return &router{
		config: cfg,
		router: routerkit.NewRouter(routerkit.WithServiceName(cfg.App.AppName)),
	}
}

func (rtr *router) handle(hfn httpHandlerFunc, svc ucaseContract.UseCase, mdws ...middleware.MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				w.WriteHeader(http.StatusInternalServerError)
				res := appctx.Response{
					Name: consts.ResponseInternalFailure,
				}

				res.BuildMessage()
				logger.Error(logger.SetMessageFormat("error %v", string(debug.Stack())))
				json.NewEncoder(w).Encode(res)
				return
			}
		}()

		var st time.Time
		var lt time.Duration

		st = time.Now()

		ctx := context.WithValue(r.Context(), "access", map[string]interface{}{
			"path":      r.URL.Path,
			"remote_ip": r.RemoteAddr,
			"method":    r.Method,
		})

		req := r.WithContext(ctx)
		defaultLang := rtr.defaultLang(req.Header.Get(consts.HeaderLanguageKey))

		if status := middleware.FilterFunc(rtr.config, req, mdws); status != consts.MiddlewarePassed {
			rtr.response(w, appctx.Response{
				Name: status,
				Lang: defaultLang,
			})

			return
		}

		resp := hfn(req, svc, rtr.config)

		resp.Lang = defaultLang

		rtr.response(w, resp)

		lt = time.Since(st)
		logger.AccessLog("access log",
			logger.Any("tag", "go-access"),
			logger.Any("http.path", req.URL.Path),
			logger.Any("http.method", req.Method),
			logger.Any("http.agent", req.UserAgent()),
			logger.Any("http.referer", req.Referer()),
			logger.Any("http.status", resp.GetCode()),
			logger.Any("http.latency", lt.Seconds()),
		)
	}
}

// response prints as a json and formatted string for DGP legacy
func (rtr *router) response(w http.ResponseWriter, resp appctx.Response) {

	w.Header().Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)

	defer func() {
		resp.BuildMessage()
		w.WriteHeader(resp.GetCode())
		json.NewEncoder(w).Encode(resp)
	}()

	return

}

// Route preparing http router and will return mux router object
func (rtr *router) Route() *routerkit.Router {

	root := rtr.router.PathPrefix("/").Subrouter()
	in := root.PathPrefix("/in/").Subrouter()
	inV1 := in.PathPrefix("/v1/").Subrouter()

	// open tracer setup
	bootstrap.RegistryOpenTracing(rtr.config)

	db := bootstrap.RegistryMariaMasterSlave(rtr.config.WriteDB, rtr.config.ReadDB, rtr.config.App.Timezone)

	// use case
	healthy := ucase.NewHealthCheck()

	repoShortUrl := repositories.NewShortUrlsRepo(db)

	insertShortUrl := short_url.NewInsertShortUrl(generator.GenerateInt64, repoShortUrl)
	detailShortUrl := short_url.NewDetailShortUrl(repoShortUrl)
	listEndpoint := endpoint.NewGetList()

	// healthy
	in.HandleFunc("/", rtr.handle(
		handler.HttpRequest,
		listEndpoint,
	)).Methods(http.MethodGet)

	in.HandleFunc("/health", rtr.handle(
		handler.HttpRequest,
		healthy,
	)).Methods(http.MethodGet)

	inV1.HandleFunc("/short-urls", rtr.handle(
		handler.HttpRequest,
		insertShortUrl,
	)).Methods(http.MethodPost)

	inV1.HandleFunc("/short-urls/{short_code:[a-zA-Z0-9]{1,255}}", rtr.handle(
		handler.HttpRequest,
		detailShortUrl,
	)).Methods(http.MethodGet)

	in.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tmpl, _ := route.GetPathTemplate()
		consts.Endpoints = append(consts.Endpoints, tmpl)
		return nil
	})

	// this route for example rest, please delete
	// example list
	//inV1.HandleFunc("/example", rtr.handle(
	//    handler.HttpRequest,
	//    el,
	//)).Methods(http.MethodGet)

	//inV1.HandleFunc("/example", rtr.handle(
	//    handler.HttpRequest,
	//    ec,
	//)).Methods(http.MethodPost)

	//inV1.HandleFunc("/example/{id:[0-9]+}", rtr.handle(
	//    handler.HttpRequest,
	//    ed,
	//)).Methods(http.MethodDelete)

	return rtr.router

}

func (rtr *router) defaultLang(l string) string {
	if len(l) == 0 {
		return rtr.config.App.DefaultLang
	}

	return l
}
