FROM golang as build

COPY . /build
WORKDIR /build

# Results 4.34 MB file
# RUN CGO_ENABLED=0 go build .

# Results 3.06 MB file
RUN CGO_ENABLED=0 go build -a -ldflags '-s' .

FROM scratch
COPY --from=build /build/gontti .
ENTRYPOINT ["./gontti"]

