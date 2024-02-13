ARG appName=app

FROM golang:1.20.5-alpine3.18 AS build
ARG appName
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY go.mod go.sum main.go ./
COPY /src src
COPY /docs docs
COPY /vendor vendor

ENV CGO_ENABLED=0
RUN go build -o ${appName}

FROM scratch
ARG appName

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy timezone data, which avaiable after installing tzdata
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta

COPY --from=build /app/${appName} /app

ENTRYPOINT [ "/app" ]