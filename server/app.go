package app

import (
	"context"
	"net/http"

	"github.com/leighmcculloch/demo-google-cloud-appengine-twirp/rpc/geo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	handler := geo.NewGeoServer(&GeoServerImpl{}, nil)
	http.Handle("/", withContext(withCountry(handler)))
}

type GeoServerImpl struct{}

func (g *GeoServerImpl) Get(ctx context.Context, r *geo.GetRequest) (*geo.GetResponse, error) {
	log.Infof(ctx, "Received 'Get' request, params: %#v", r)

	country := contextCountry(ctx)

	log.Infof(ctx, "Raw country: %q", country)

	if country == "" {
		country = "ZZ"
	}

	log.Infof(ctx, "Final country: %q", country)

	resp := &geo.GetResponse{
		Country: country,
	}

	log.Infof(ctx, "Returning response: %#v", resp)

	return resp, nil
}

func withCountry(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		country := r.Header.Get("X-AppEngine-Country")
		ctx := r.Context()
		ctx = contextWithCountry(ctx, country)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func withContext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
