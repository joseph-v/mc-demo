# Multi-Cloud Demo

This project shows how to use multi-cloud project for a streaming usecase

## Setting up demo

```bash
#Clone repo and start multi-cloud
cd $GOPATH/src/$USER
git clone https://github.com/joseph-v/multi-cloud.git
cd multi-cloud
git checkout UpdatedClient

#Create soft link for the opensds dependacy
ln -s $GOPATH/src/sodafoundation $GOPATH/src/opensds

#build and start multi-cloud
make docker
docker-compose up -d

#Get AKSK generated from OpenSDS Dashboard UI
export OS_ACCESS_KEY=ZNRJARg7wkfm9wxzuIeD

$Export environment variables
export MULTI_CLOUD_IP=192.168.20.158
export MICRO_SERVER_ADDRESS=:8089
export OS_AUTH_AUTHSTRATEGY=keystone

#Export environment variable
export OPENSDS_ENDPOINT=http://192.168.20.158:50040
export OPENSDS_AUTH_STRATEGY=keystone
export OS_AUTH_URL=http://192.168.20.158/identity
export OS_USERNAME=admin
export OS_PASSWORD=opensds@123
export OS_TENANT_NAME=admin
export OS_PROJECT_NAME=admin
export OS_USER_DOMIN_ID=default

export KAFKA_URL=localhost:9092
export KAFKA_TOPIC=streams-wordcount-processor-output
export KAFKA_GROUP_ID=TestGroupID

#Clone demo repo and start demo
git clone https://github.com/joseph-v/mc-demo.git
cd mc-demo
go run demo.go gelato_client.go
```