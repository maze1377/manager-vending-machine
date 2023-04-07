FROM golang:1.20 as build_image

WORKDIR /opt

ADD Makefile go.mod go.sum /opt/
RUN make dependencies
ADD . /opt/
RUN make vendingd

FROM ubuntu:22.04 as runner

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG=C.UTF-8 LC_ALL=C.UTF-8
ENV TZ 'Asia/Tehran'
RUN apt-get update --fix-missing && \
    apt-get upgrade -y && \
    apt-get install -y ca-certificates && \
    echo $TZ > /etc/timezone && \
    apt-get install -y tzdata && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    apt-get clean

COPY --from=build_image /opt/vendingd /bin/vendingd

EXPOSE 10000
ENTRYPOINT ["/bin/vendingd"]
CMD ["version"]
