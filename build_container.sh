app_version=`cat Version | head -1`
docker build -t japa-tracker:$app_version .
