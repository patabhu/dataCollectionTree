data collection tree apis.
----------------------------------------------------
To start the server manually follow the steps below.
1. cd dataCollectionTree in go path.
2. source vars.env for environmental variables for http host port and db configs.
`source vars.env`
3. in dataCollectionTree run main.go file.
`go run main.go`
4. check-out endpoints in controller in dataCollectionTree.go
---------------------------------------------------------------------------------------
Running 2 different Docker containers follow below instructions:-

For Go server container
1. Run cmd "sudo docker build -t dataCollectionTree ." from in your dataCollectionTree dir.

2. Run your Go server image in container with cmd "sudo docker run -it --rm -e HTTP_HOST=0.0.0.0:10000 -p 10000:10000 --network=bridge --name dataCollectionTree -d dataCollectionTree".

3. Go server images will be running in container name as dataCollectionTree.
