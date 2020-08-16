#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache gcc musl-dev git
RUN go get -d -v ./... 
WORKDIR /app/src
ADD . . 
RUN GOOS=linux go build -a -o ./app -- code.go

#final stage
FROM alpine:latest
RUN apk --no-cache add gcc musl-dev ca-certificates
WORKDIR /app
COPY --from=builder app/src/app .
CMD ./app
LABEL Name=tg Version=0.0.1
EXPOSE 8000


# #build stage\nFROM golang:alpine AS builder\nWORKDIR /app\nCOPY . .\nRUN apk add --no-cache gcc musl-dev git\nRUN go get -d -v ./... \nWORKDIR /app/src\nADD . . \nRUN GOOS=linux go build -a -o ./app -- code.go\n\n#final stage\nFROM alpine:latest\nRUN apk --no-cache add gcc musl-dev ca-certificates\nWORKDIR /app\nCOPY --from=builder app/src/app .\nCMD ./app\nLABEL Name=tg Version=0.0.1\nEXPOSE 8080