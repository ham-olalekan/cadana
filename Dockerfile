FROM golang:1.16-alpine

WORKDIR /cadana/

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /cadana/ .

EXPOSE 8081
ENTRYPOINT [ "/cadana/" ]