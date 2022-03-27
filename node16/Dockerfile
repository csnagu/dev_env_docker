FROM node:16-slim

RUN apt update && apt install -y \
    curl \
    git

# Clean
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
