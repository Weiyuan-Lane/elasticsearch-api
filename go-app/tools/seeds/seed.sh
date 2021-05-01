#!/bin/sh

index=users

curl -X POST 'http://localhost:8080/indices' -d "{\"id\":\"$index\"}"

jq -c '.[]' tools/seeds/data.json | while read data; do
  curl -X POST "http://localhost:8080/indices/$index/documents" -d "$data"
done

