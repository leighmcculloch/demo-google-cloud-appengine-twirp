# demo-cloud-appengine-twirp

A demonstration of how to use [twirp](https://github.com/twitchtv/twirp) on Google Cloud AppEngine.

See the [Makefile](Makefile) for setup, run and deploy commands.

This won't necessarily always be deployed and running at the URL below, but if it is you can try curling directly:

```
curl \
	--request POST \
	--header "Content-Type: application/json" \
	--data '{}' \
	https://demotwirp.appspot.com/twirp/app.Geo/Get
```
