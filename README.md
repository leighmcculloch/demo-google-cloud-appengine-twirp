# demo-cloud-appengine-twirp

A demonstration of how to use [twirp](https://github.com/twitchtv/twirp) on Google Cloud AppEngine.

## Server

See the [server/Makefile](server/Makefile) for setup, run and deploy commands.

This won't necessarily always be deployed and running at the URL below, but if it is you can try curling directly:

```
curl \
	--request POST \
	--header "Content-Type: application/json" \
	--data '{}' \
	https://demotwirp.appspot.com/twirp/geo.Geo/Get
```

## Client

The client talks to the server using the protobuf format.

```
go run client/main.go https://demotwirp.appspot.com
```
