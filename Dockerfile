FROM golang:latest AS build-stage
RUN mkdir /build 
ADD ./server /build
ADD go.mod /build
ADD go.sum /build
ADD main.go /build
WORKDIR /build 
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./main.go


FROM scratch as production-stage
COPY --from=build-stage /build/main /app/
WORKDIR /app
EXPOSE 9010
CMD ["./main"]