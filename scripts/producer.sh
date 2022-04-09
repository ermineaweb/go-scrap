#!/bin/bash

count=1
broker="kafka:9092"
topic="stream_start"

echo "ponce" | kcat -b "kafka:9092" -P -t "stream_start" && \
echo "zerator" | kcat -b "kafka:9092" -P -t "stream_start" && \
echo "asmongold" | kcat -b "kafka:9092" -P -t "stream_start" && \
echo "aminematue" | kcat -b "kafka:9092" -P -t "stream_start" && \
echo "otplol_" | kcat -b "kafka:9092" -P -t "stream_start"

echo "aminematue" | kcat -b "kafka:9092" -P -t "stream_stop"

while :
do
    printf -v date '%(%Y/%m/%d-%H:%M:%S)T' -1 
    printf -v message "message n°%s - %s" "$count" "$date"
    printf "topic '%s'\n'%s'\n" "$topic" "$message"
    # printf "%s" "$message" | kcat -b $broker -P -t $topic
    printf "%s" "zerator" | kcat -b $broker -P -t $topic
    count=$((count + 1))
    sleep 3600
done