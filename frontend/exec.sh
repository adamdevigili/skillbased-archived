#!/usr/bin/env bash
docker build -t skillbased-frontend .
docker run --init -p 3000:3000 -it skillbased-frontend
