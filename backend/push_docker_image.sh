#!/bin/bash
docker buildx build --push --platform linux/arm/v7,linux/arm64/v8,linux/amd64 -t rof20004/carrefour-backend:latest .