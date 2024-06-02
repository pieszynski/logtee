#!/bin/bash

i=0;
while true;
do
  echo "{\"@timestamp\":\"$(date -I'seconds')\",\"message\":\"log line $i\",\"i\":$i}"
  i=$((i+1));
  sleep 2;
done 