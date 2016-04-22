#!/bin/bash

# start alertmanager
cd /home/phelan/go/src/github.com/prometheus/alertmanager
./alertmanager -config.file=/home/phelan/Repositories/mockingProduction/alerting/alertmanager_config.yml

# start prometheus on localhost:9090/ [metrics, graphs]
/opt/prometheus-0.17.0.linux-386/prometheus -config.file=/home/phelan/Repositories/mockingProduction/prometheus/prometheus.yml -storage.local.path=/tmp/data -alertmanager.url=http://localhost:9093/

# start grafana on localhost:3000
sudo service grafana-server start

# start http server
http_server/http_server -port=7070

# start curler
hammers/curler.sh

