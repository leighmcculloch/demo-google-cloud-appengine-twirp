package app

import "context"

type key struct{}

var keyCountry key

func contextWithCountry(ctx context.Context, country string) context.Context {
	return context.WithValue(ctx, &keyCountry, country)
}

func contextCountry(ctx context.Context) string {
	return ctx.Value(&keyCountry).(string)
}
