export PROJECT_ID=demotwirp
export REGION=us-central

# Run the app locally using the AppEngine development server.
run:
	go install && dev_appserver.py --port=8080 app.yaml

# Curl the local development server using a json request.
runtestjson:
	curl --request POST --header "Content-Type: application/json" --data '{}' http://localhost:8080/twirp/geo.Geo/Get

# Curl the local development server using a protobuf request.
runtestprotobuf:
	echo "" | protoc --proto_path=../rpc/geo/ --encode geo.GetRequest ../rpc/geo/geo.proto \
		| curl -s --request POST --header "Content-Type: application/protobuf" --data-binary @- http://localhost:8080/twirp/geo.Geo/Get \
		| protoc --proto_path=../rpc/geo/ --decode geo.GetResponse ../rpc/geo/geo.proto

# Deploy the app to AppEngine.
deploy:
	gcloud --project $(PROJECT_ID) app deploy -q

# Create the AppEngine project and setup the application.
create-projects:
	gcloud projects create $(PROJECT_ID)
	gcloud --project $(PROJECT_ID) app create --region $(REGION)
