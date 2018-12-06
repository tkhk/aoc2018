#!/bin/bash

ENDPOINT=http://localhost:8080/part2

curl -X POST --data-binary @input.txt ${ENDPOINT}
