docker build . -t go-web-app

docker run --name go-api --net springfield_default -p 8080:8080 --env-file ./.env.staging -d go-web-app