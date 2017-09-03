go get github.com/lib/pq
go get -u -d github.com/mattes/migrate/cli
go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli
chmod +x ./bin/wait-for-it.sh
./bin/wait-for-it.sh postgres:5432 -s --timeout=0 -- echo "Pg is up"
migrate -database postgres://docker:docker@postgres:5432/test?sslmode=disable -path ./db/migrations up
