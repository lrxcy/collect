#!/bin/bash
pageNum=$1
curl -XGET "http://127.0.0.1:8080/v1?page=${pageNum}"
