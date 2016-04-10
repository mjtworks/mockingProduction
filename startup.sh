#!/bin/bash

# start prometheus on localhost:9090/ [metrics, graphs]
./prometheus -config.file=/home/phelan/Repositories/mockingProduction/prometheus/prometheus.yml -storage.local.path=/tmp/data

# start grafana on localhost:3000
sudo service grafana-server start

# start http server
./http_server -port=7070

