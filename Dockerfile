from golang:1.19.2-bullseye

WORKDIR /app

COPY . .

RUN go mod download
# for light sail container
# RUN GOOS=linux GOARCH=amd64 go build -o goapp ./cmd/myapp

# for local
RUN go build -o goapp ./cmd/myapp

EXPOSE 8080


CMD [ "./goapp" ]