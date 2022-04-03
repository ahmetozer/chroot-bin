FROM golang:1.17 as build
WORKDIR /WORKDIR
COPY ["main.go", "go.mod", "./"] 
RUN go build -ldflags="-w -s" -o /usr/bin/chroot-bin

FROM scratch
COPY --from=build /usr/bin/chroot-bin /usr/bin/chroot-bin
ENTRYPOINT ['/usr/bin/chroot-bin']
LABEL org.opencontainers.image.source="https://github.com/ahmetozer/chroot-bin"
