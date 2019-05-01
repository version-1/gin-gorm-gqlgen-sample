FROM golang:latest
ENV ROOTDIR /app
ENV WORKDIR /app/src/gin-sample
ENV GOPATH $ROOTDIR
ENV GOBIN $ROOTDIR/bin
WORKDIR $WORKDIR
VOLUME $ROOTDIR
ADD . $ROOTDIR

RUN apt-get update && \
    apt-get -y install mysql-client && \
    echo "export PATH=$PATH:$GOBIN" > /root/.bashrc && \
    cd $WORKDIR && $GOBIN/dep ensure

CMD ["go", "run", "cmd/app/main.go"]

