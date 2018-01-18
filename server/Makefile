export PROJECT_ID=demotwirp
export REGION=us-central

# Run the app locally using the AppEngine development server.
run:
	go install && dev_appserver.py --port=8080 app.yaml

# Curl the local development server.
runtest:
	curl --request POST --header "Content-Type: application/json" --data '{}' http://localhost:8080/twirp/geo.Geo/Get

# Deploy the app to AppEngine.
deploy:
	gcloud --project $(PROJECT_ID) app deploy -q

# Create the AppEngine project and setup the application.
create-projects:
	gcloud projects create $(PROJECT_ID)
	gcloud --project $(PROJECT_ID) app create --region $(REGION)
