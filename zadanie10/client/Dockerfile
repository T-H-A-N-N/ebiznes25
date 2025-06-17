FROM ubuntu:22.04

ENV TZ="Europe/Warsaw"
ENV DEBIAN_FRONTEND=noninteractive
ENV NODE_VERSION=18

RUN apt-get update && \
    apt-get install -y curl ca-certificates gnupg && \
    curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION}.x | bash - && \
    apt-get install -y nodejs && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN node -v && npm -v

WORKDIR /app
COPY . .

RUN npm install

EXPOSE 3000
CMD ["npm", "start"]

