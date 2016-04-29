Add alerts to prometheus so bad situations are detected:
- Create alerting_rules file that contains the rules you want to use to fire alerts
- Use promtool to check the syntax of the rules
$ /opt/prometheus-0.17.0.linux-386/promtool check-rules alerting_rules
Checking alerting_rules
  SUCCESS: 2 rules found
- Modify the prometheus config file to point at the rules file containing the alerts
- Edit prometheus.yml and add the rules filepath to the rule_files section.
- Alerting rules are configured in Prometheus in the same way as recording rules
  http://prometheus.io/docs/alerting/rules/
- Restart prometheus server
- Check out http://localhost:9090/alerts
- add a 500 handler so I could intentionally cause 500s
- Created a new script to just make 500 requests (still some 200’s from prometheus scraping)
- Checked out the alert endpoint and saw it was pending
- installed alertmanager dependencies
- fought with the gopath, now installed to  ~/go/src/github.com/prometheus/alertmanager
- made the config file, added my email address as the receiver
- ran alertmanager and prometheus and got an alert to fire
	$ ./alertmanager -config.file=/home/phelan/Repositories/mockingProduction/alerting/alertmanager_config.yml
	$ /opt/prometheus-0.17.0.linux-386/prometheus -config.file=/home/phelan/Repositories/mockingProduction/prometheus/prometheus.yml -storage.local.path=/tmp/data -alertmanager.url=http://localhost:9093/	
- installed sendmail
	$ sudo apt-get install sendmail
	$ sudo sendmailconfig
- trying to configure sendmail, can't send to gmail for some reason. 

this works:
$ echo "Just testing my sendmail gmail relay" | mail -s "Sendmail gmail Relay" phelan@ratqueens
this doesn't:
echo "Just testing my sendmail gmail relay" | mail -s "Sendmail gmail Relay" phelan.vendeville@gmail.com

- used the tutorial here to set it up:
https://linuxconfig.org/configuring-gmail-as-sendmail-email-relay
- and the tutorial here to fix my /etc/hosts
http://serverfault.com/questions/58363/my-unqualified-host-name-foo-bar-unknown-problem