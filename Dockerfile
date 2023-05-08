FROM registry.access.redhat.com/ubi9-minimal:9.1 as builder
RUN microdnf install -y golang && microdnf clean all
WORKDIR /go
COPY . .
RUN go build -o color-demo .

FROM registry.access.redhat.com/ubi9-micro:9.1
WORKDIR /srv/color-demo
COPY --from=builder /go/templates ./templates
COPY --from=builder /go/color-demo /usr/bin/color-demo
ENV BASE_PATH /srv/color-demo
CMD /usr/bin/color-demo
