Installing grafana
(see http://docs.grafana.org/installation/debian/)

# add graphana repo to sources
$ vim /etc/apt/sources.list
deb https://packagecloud.io/grafana/stable/debian/ wheezy main

# grab package cloud key
$ curl https://packagecloud.io/gpg.key | sudo apt-key add -

# update and install
$ sudo apt-get update
$ sudo apt-get install grafana

# run the server
sudo service grafana-server start

# check out the interface
http://192.168.0.107:3000
# change the password
http://192.168.0.107:3000/admin/users
# look at the config file
$ sudo vim /etc/grafana/grafana.ini

# start prometheus and the webserver
$ ./prometheus -config.file=/home/phelan/Repositories/moc
kingProduction/prometheus/prometheus.yml -storage.local.path=/tmp/data
$ ./http_server -port=7070
# start graphana
$ sudo service grafana-server start

# configure prometheus data source
http://prometheus.io/docs/visualization/grafana/

Adding dashboards
- http://docs.grafana.org/guides/gettingstarted/

