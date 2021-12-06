FROM golang:1.16.6 AS base

RUN apt-get update && apt-get install -y --no-install-recommends postgresql-client \
	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /src

COPY go.mod go.sum /src/

RUN go mod download

COPY . /src/

CMD [ "sh", "-c", "./build/test.sh" ]
