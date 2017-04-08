package server

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"context"

	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "server: ", log.Lshortfile)
}

func Init(addr string, proxy string) {

	/*
		var (
			httpAddr     = flag.String("http.addr", ":8000", "Address for HTTP (JSON) server")
			consulAddr   = flag.String("consul.addr", "", "Consul agent address")
			retryMax     = flag.Int("retry.max", 3, "per-request retries to different instances")
			retryTimeout = flag.Duration("retry.timeout", 500*time.Millisecond, "per-request timeout, including retries")
		)*/

	var serverLogger kitlog.Logger
	serverLogger = kitlog.NewLogfmtLogger(os.Stderr)
	serverLogger = kitlog.With(serverLogger, "listen", addr, "caller", kitlog.DefaultCaller)
	

	fieldKeys := []string{"method", "error"}

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "dfk",
		Subsystem: "parse_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "dfk",
		Subsystem: "parse_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "dfk",
		Subsystem: "parse_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	var svc ParseService
	svc = parseService{}
	svc = proxyingMiddleware(context.Background(), proxy, serverLogger)(svc)
	svc = statsMiddleware("18")(svc)
	svc = cachingMiddleware()(svc)
	//	svc = resultCachingMiddleware()(svc)
	svc = loggingMiddleware(serverLogger)(svc)
	svc = robotsTxtMiddleware()(svc)
	svc = instrumentingMiddleware(requestCount, requestLatency, countResult)(svc)

	getHTMLHandler := httptransport.NewServer(
		makeGetHTMLEndpoint(svc),
		decodeGetHTMLRequest,
		encodeResponse,
	)

	marshalDataHandler := httptransport.NewServer(
		makeMarshalDataEndpoint(svc),
		decodeMarshalDataRequest,
		encodeResponse,
	)

	checkServicesHandler := httptransport.NewServer(
		makeCheckServicesEndpoint(svc),
		decodeCheckServicesRequest,
		encodeCheckServicesResponse,
	)

	router := httprouter.New()
	router.Handler("POST", "/app/gethtml", getHTMLHandler)
	router.Handler("POST", "/app/marshaldata", marshalDataHandler)
	router.Handler("POST", "/app/chkservices", checkServicesHandler)
	//router.Handler("GET", "/", http.FileServer(http.Dir("web/")))
	//router.ServeFiles("/static/*filepath", http.Dir("web/static"))
	router.ServeFiles("/static/*filepath", http.Dir("web/static"))

	//router.HandlerFunc("GET", "/test1", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "___TEST___")
	//})
	router.HandlerFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	router.Handler("GET", "/metrics", stdprometheus.Handler())

	serverLogger.Log("msg", "HTTP", "addr", addr)
	serverLogger.Log("err", http.ListenAndServe(addr, router))
}