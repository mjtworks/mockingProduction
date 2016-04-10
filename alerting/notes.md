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

Next: Configure prometheus to send these bad conditions to alertmanager, so alerts can actually get delivered
