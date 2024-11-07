## Initialize database(Docker)
```
docker-compose up -d
```

## Update config filename
```
cd manage-money && mv config/prouction_example.yaml config/production.yaml
```

## Build
```
cd manage-money && go build -o moneymanage cmd/server/main.go
```

## Run
```
./moneymanage
```