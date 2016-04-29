#!/bin/bash

while [ 1 ]
do
  curl -X GET localhost:7070/cause_500 >/dev/null 2>&1
  sleep $[ ( $RANDOM % 5 )  + 1 ]s
done

