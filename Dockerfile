FROM golang:1.20-bullseye AS bin-builder

WORKDIR /x

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o bin/app

FROM gcr.io/distroless/static-debian11:nonroot

WORKDIR /x

COPY --from=js-builder /x/dist dist
COPY --from=bin-builder /x/bin/app app
COPY public public
COPY garden.yaml garden.yaml

CMD ["/x/app", "serve"]