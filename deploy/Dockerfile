#build
FROM golang:alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN apk add upx && CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o nishcal && upx --best --lzma nishcal


#final-stage
FROM alpine
WORKDIR /app
RUN apk update && apk add --no-cache ffmpeg
CMD [ "./nishcal" ]
COPY --from=build /app/nishcal .