FROM quay.io/projectquay/golang:1.18
WORKDIR /go/src/github.com/ocp-tigers/rhte22-devsecops-app/
COPY main.go . 
COPY go.mod .
COPY static ./static
RUN chown -R 9999:9999 ./static/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=0 /go/src/github.com/ocp-tigers/rhte22-devsecops-app/main .
COPY --from=0 /go/src/github.com/ocp-tigers/rhte22-devsecops-app/static ./static
EXPOSE 8080
USER 9999
CMD ["/main"]
