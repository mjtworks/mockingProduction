#!/bin/bash
# randomly issue different curl commands
requests=("curl -X GET localhost:8080/" "curl -X POST localhost:8080/" "curl -X GET localhost:8080/fail" "curl -X GET localhost:8080/redirect_me")
seed=$$$(date +%s)

for i in `seq 1 500`
do
  selectedexpression=${requests[$seed % ${#requests[@]} ]}
  `$selectedexpression`
done
