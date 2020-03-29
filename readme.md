# Initial Notes
*Official Readme to come*

## Set up Hosts file
sudo vi /etc/hosts
127.0.0.1 influxdb-service

## Start an Influx
docker run -d -expose -p 8086:8086 --name influxdb -e INFLUXDB_ADMIN_PASSWORD=PASSWORD influxdb/influxdb

## Setting up DB
curl -G http://influxdb-service:8086/query --data-urlencode "q=CREATE DATABASE CardPrices"

## Create the K8s Configmap
kubectl create configmap card-list-configmap --from-file=cardList.json