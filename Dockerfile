FROM golang:1.20.2 as builder

ARG BACKENDTYPE OAIKEY K8SCONFIG
WORKDIR /app
COPY ./ /app

RUN go get . && go run main.go
RUN wget -qO- https://github.com/k8sgpt-ai/k8sgpt/releases/download/v0.1.4/k8sgpt_Linux_x86_64.tar.gz  | tar xvzf -


FROM golang:1.20.2

WORKDIR /app
COPY --from=builder /app/ /app/

RUN mkdir /root/.kube/ && cp k8sgptconfig /root/.kube/config && cp k8sgpt.yaml /root/.k8sgpt.yaml


ENTRYPOINT ["./k8sgpt", "analyze", "--explain ", "--namespace=default", " --filter=Pod"]