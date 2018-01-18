package ctxvalues

import "context"

type key struct{}

var keyCountry key

func WithCountry(ctx context.Context, country string) context.Context {
	return context.WithValue(ctx, &keyCountry, country)
}

func Country(ctx context.Context) string {
	return ctx.Value(&keyCountry).(string)
}
