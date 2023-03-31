FROM golang:1.21 as builder

WORKDIR /app
COPY main.go /app

RUN go run main.go
RUN wget -qO- https://github.com/k8sgpt-ai/k8sgpt/releases/download/v0.1.4/k8sgpt_Linux_x86_64.tar.gz  | tar xvzf -


FROM gcr.io/distroless/static

COPY --from=builder /app/k8sgpt* /app

RUN mkdir /root/.kube/ && cp k8sgptconfig /root/.kube/config && cp k8sgpt.yaml /root/.k8sgpt.yaml

ENTRYPOINT ["./k8sgpt", "analyze", "--explain ", "--namespace=default", " --filter=Pod"]