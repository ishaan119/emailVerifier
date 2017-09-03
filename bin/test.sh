echo "Removing exited containers"
docker rm -f $(docker ps -aqf'status=exited')

docker-compose -f 'docker-compose.test.yml' -p ci up --build --abort-on-container-exit
exit $(docker wait ci_newsletter-engine_1)
