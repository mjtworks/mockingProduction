#!/bin/bash
# randomly issue different curl commands
requests=("curl -X GET localhost:7070/" "curl -X POST localhost:7070/" "curl -X GET localhost:7070/fail" "curl -X GET localhost:7070/redirect_me")

for i in `seq 1 500`
do
  selectedexpression=${requests[$RANDOM % ${#requests[@]} ]}
  `$selectedexpression` >/dev/null 2>&1 
done

