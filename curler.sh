#!/bin/bash
for i in `seq 1 50`
do
  curl -X GET localhost:8080/
  curl -X POST localhost:8080/
  curl -X GET localhost:8080/fail
done

