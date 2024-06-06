# database image
$DB_IMAGE = "postgres:16.3-alpine3.20"

# database name
$DB_NAME = "postgres"

# database type
$DB_TYPE = "postgres"

# database username
$DB_USER = "postgres"

# database password
$DB_PWD = "Postgres123"

# psql URL
$IP = "127.0.0.1"

$PSQLURL = "$($DB_TYPE)://$($DB_USER):$($DB_PWD)@$($IP):5432/$($DB_NAME)"

# sqlc yaml file
$SQLC_YAML = "./sqlc.yaml"

function postgresup {
    Write-Host "Creating the container, postgres database, and bound volume"
    docker run --name test-postgres -v ${PWD}:/usr/share/chapter_01 -e POSTGRES_PASSWORD=$DB_PWD -p 5432:5432 -d $DB_IMAGE
}

function postgresdown {
    Write-Host "Stop the container"
    docker stop test-postgres || true
    Write-Host "Removing the container, postgres database, and its bound volume"
    docker rm test-postgres || true
    docker volume rm $(docker volume ls -q) || true
}

function psql {
    Write-Host "Connecting to the postgres database"
    docker exec -it test-postgres psql $PSQLURL
}

function createdb {
    Write-Host "Connecting to the postgres database and creating the schema, table, and indexes"
    docker exec -it test-postgres psql $PSQLURL -c "\i /usr/share/chapter_01/db/schema.sql"
}

function teardown_recreate {
    postgresdown
    Start-Sleep -Seconds 5
    postgresup
    Start-Sleep -Seconds 80
    createdb
}

function generate {
    Write-Host "Generating Go models with sqlc"
    sqlc generate -f $SQLC_YAML
}

function teardown_recreate_generate {
    teardown_recreate
    generate
}

function postgresup_createdb_generate {
    postgresup
    Start-Sleep -Seconds 80
    createdb
    generate
}

function build {
    Write-Host "Building database main sample app"
    Write-Host "Building for Linux environments: sampledb"
	go build -o sampledb .
    Write-Host "Building for Windows environments. sampledb.exe"
    go build -o sampledb.exe .
}
