FROM --platform=$BUILDPLATFORM golang:1.27.4-alpine AS build
ARG TARGETOS TARGETARCH
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o /sample .

FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /sample /sample
EXPOSE 8080
ENTRYPOINT ["/sample"]
