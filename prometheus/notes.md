wget "https://github.com/prometheus/prometheus/releases/download/0.17.0/prometheus-0.17.0.linux-386.tar.gz"
sudo cp prometheus-0.17.0.linux-386.tar.gz /opt/
tar xvfz prometheus-0.17.0.linux-386.tar.gz
sudo rm prometheus-0.17.0.linux-386.tar.gz
./prometheus -config.file=/home/phelan/Repositories/mockingProduction/prometheus/prometheus.yml -storage.local.path=/tmp/data
...
- start example targets
- edit yaml file to include example targets
- restart prometheus
- check out http://192.168.0.109:8080/metrics for the stats page (refresh for
update)
- visit http://192.168.0.109:9090/graph (the graph endpoint of the prometheus
service) to see visualizations and access the expression browser, where you can
run queries.

next steps will be to instrument the http server.
