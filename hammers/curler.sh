#!/bin/bash
# randomly issue different curl commands
requests=("curl -X GET localhost:7070/" "curl -X POST localhost:7070/" "curl -X GET localhost:7070/fail" "curl -X GET localhost:7070/redirect_me" "curl -X GET localhost:7070/cause_500")

while [ 1 ]
do
  selectedexpression=${requests[$RANDOM % ${#requests[@]} ]}
  `$selectedexpression` >/dev/null 2>&1 
  sleep $[ ( $RANDOM % 5 )  + 1 ]s
done

