wget "https://github.com/prometheus/prometheus/releases/download/0.17.0/prometheus-0.17.0.linux-386.tar.gz"
sudo cp prometheus-0.17.0.linux-386.tar.gz /opt/
tar xvfz prometheus-0.17.0.linux-386.tar.gz
sudo rm prometheus-0.17.0.linux-386.tar.gz
./prometheus -config.file=/home/phelan/Repositories/mockingProduction/prometheus/prometheus.yml -storage.local.path=/tmp/data

left off: https://prometheus.io/docs/introduction/getting_started/, "Starting prometheus"
