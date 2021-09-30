# Simple launch service

### To generate a protobuf run 
```sh
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    launch_grpc/launch.proto
```

## How to run the application 

### Locally

Server :
```sh
$ export GRAPHQL_URL=https://api.spacex.land/graphql/
$ export SERVER_PORT=50051
$ go run launch_server/main.go
```

Client
```sh
$ export LAUNCH_ADDRESS=localhost:50051
$ go run launch_client/main.go 4 2
```
After running the client you will see the output like below  
Note: some time external service was unable to connect and give error `could not get the launches: rpc error: code = DeadlineExceeded desc = context deadline exceeded`  

```
2021/09/30 10:02:14 Started Launces
2021/09/30 10:02:14 total 4 launches received
2021/09/30 10:02:14 Launces: [launch_success:true mission_name:"Starlink-15 (v1.0)" launch_success:true mission_name:"Sentinel-6 Michael Freilich" launch_success:true mission_name:"Crew-1" launch_success:true mission_name:"GPS III SV04 (Sacagawea)"]
2021/09/30 10:02:14 Launce received for id 2 as s: mission_name:"DemoSat"
```

### Using docker-compose

```sh
$ docker-compose up
```
output :
```
Starting grpc-test_launch_service_1 ... done
Starting grpc-test_launch_client_1  ... done
Attaching to grpc-test_launch_service_1, grpc-test_launch_client_1
launch_service_1  | 2021/09/29 17:32:33 server listening at [::]:50051
launch_client_1   | 2021/09/29 17:32:35 Launces: [launch_success:true  mission_name:"Starlink-15 (v1.0)" launch_success:true  mission_name:"Sentinel-6 Michael Freilich" launch_success:true  mission_name:"Crew-1" launch_success:true  mission_name:"GPS III SV04 (Sacagawea)" launch_success:true  mission_name:"Starlink-14 (v1.0)" launch_success:true  mission_name:"Starlink-13 (v1.0)" mission_name:"Starlink-12 (v1.0)" launch_success:true  mission_name:"Starlink-11 (v1.0)" launch_success:true  mission_name:"SAOCOM 1B, GNOMES-1, Tyvak-0172" launch_success:true  mission_name:"Starlink-10 (v1.0) & SkySat 19-21"]
launch_client_1   | 2021/09/29 17:32:35 could not get the launch: rpc error: code = DeadlineExceeded desc = context deadline exceeded
```

Note: some if you are not able to call the service and get out put as, try again
```
Creating grpc-test_launch_service_1 ... done
Creating grpc-test_launch_client_1  ... done
Attaching to grpc-test_launch_service_1, grpc-test_launch_client_1
launch_service_1  | 2021/09/29 17:32:10 server listening at [::]:50051
launch_client_1   | 2021/09/29 17:32:12 could not get the launches: rpc error: code = DeadlineExceeded desc = context deadline exceeded
```