FROM golang:1.19 AS build

WORKDIR /go/src/prefixfinder

COPY . .

# Downloads all the dependencies in advance
RUN go mod download

# Install Jade/Pug template-to-Go code generator
RUN go install github.com/Joker/jade/cmd/jade@latest
RUN jade -writer -pkg=main similarwords.pug

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving only required files to the 'final image' to make image smaller
FROM alpine:latest as release

WORKDIR /app

COPY --from=build /go/src/prefixfinder/app .
COPY --from=build /go/src/prefixfinder/enable1.txt .
COPY --from=build /go/src/prefixfinder/views /app/views

RUN chmod +x /app/app

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

CMD ["/app/app"]
