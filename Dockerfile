FROM golang:1.16

RUN echo $GOPATH
RUN mkdir -p $GOPATH/src/github.com/bapung
WORKDIR $GOPATH/src/github.com/bapung
RUN git clone https://github.com/bapung/terraform-provider-idcloudhost
WORKDIR $GOPATH/src/github.com/bapung/terraform-provider-idcloudhost
RUN make install
