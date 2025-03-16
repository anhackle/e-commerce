## Go to the project directory
```
cd cungsao
```

## Run docker compose to initialize MySQL, Redis and Kafka
```
docker compose up -d
```

## Make migrations
```
make upse
```

## Update config filename
```
mv config/production.yaml.example config/production.yaml
```

## Build
```
go build -o cungsao cmd/server/main.go
```

## Run
```
./cungsao
```