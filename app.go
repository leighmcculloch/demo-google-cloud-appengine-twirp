package app

import (
	"context"
	"net/http"

	"github.com/leighmcculloch/demo-google-cloud-appengine-twirp/internal/ctxvalues"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	handler := NewGeoServer(&GeoServerImpl{}, nil)
	http.Handle("/", injectAppEngineContextHandler(handler))
}

type GeoServerImpl struct{}

func (g *GeoServerImpl) Get(ctx context.Context, r *GetRequest) (*GetResponse, error) {
	log.Infof(ctx, "Received 'Get' request, params: %#v", r)

	country := ctxvalues.Country(ctx)

	log.Infof(ctx, "Raw country: %q", country)

	if country == "" {
		country = "ZZ"
	}

	log.Infof(ctx, "Final country: %q", country)

	resp := &GetResponse{
		Country: country,
	}

	log.Infof(ctx, "Returning response: %#v", resp)

	return resp, nil
}

func injectAppEngineContextHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		log.Infof(ctx, "Received %s with headers: %#v", r.URL.Path, r.Header)

		country := r.Header.Get("X-AppEngine-Country")
		ctx = ctxvalues.WithCountry(ctx, country)

		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)

		log.Infof(ctx, "Finished handling")
	})
}
