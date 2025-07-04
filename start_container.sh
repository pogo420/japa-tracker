app_version=`cat Version | head -1`
docker run -d --name japa-tracker-v$app_version --env-file=.env -p 8080:8080 -e DB_HOST=host.docker.internal japa-tracker:$app_version
