#!/bin/bash

ENDPOINT=http://localhost:8080/part1

curl -X POST --data-binary @input.txt ${ENDPOINT}
