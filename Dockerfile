FROM golang:latest
ENV WORKDIR /app
ENV PROJECTDIR /app/src/gin-sample/
ENV GOPATH $WORKDIR
ENV GOBIN $WORKDIR/bin
WORKDIR $WORKDIR

ADD . $WORKDIR

RUN apt-get update && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    apt-get -y install mysql-client && \
    echo "export PATH=$PATH:$GOBIN" > /root/.bashrc && \
    cd $PROJECTDIR && $GOBIN/dep ensure


CMD ["go", "run", "main.go"]

