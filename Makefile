export PROJECT_ID=demotwirp
export REGION=us-central

# Run the app locally using the AppEngine development server.
run:
	dev_appserver.py --port=8080 app.yaml

# Curl the local development server.
runtest:
	curl --request POST --header "Content-Type: application/json" --data '{}' http://localhost:8080/twirp/app.Geo/Get

# Build the twirp server and generated code from the proto definition.
protoc: protoc-deps
	protoc --go_out=. --twirp_out=. ./service.proto

# Deploy the app to AppEngine.
deploy:
	gcloud --project $(PROJECT_ID) app deploy -q

# Create the AppEngine project and setup the application.
create-projects:
	gcloud projects create $(PROJECT_ID)
	gcloud --project $(PROJECT_ID) app create --region $(REGION)

# Install protoc build dependencies.
protoc-deps:
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/twitchtv/twirp/protoc-gen-twirp
