FROM golang:1.12.7
MAINTAINER 121

RUN sed -i 's/http:\/\/archive\.ubuntu\.com\/ubuntu\//http:\/\/mirrors\.ustc\.edu\.cn\/ubuntu\//g' /etc/apt/sources.list
# 添加supervisor配置文件
RUN apt-get update
RUN apt-get -y install supervisor


RUN go get

#COPY server.conf /etc/supervisor/conf.d
# 初始化文件夹 添加主程序二进制执行文件
#RUN mkdir /logs
# RUN mkdir /go/src/liuma
# ADD config.yaml /go/src/liuma
#ADD code/ /go/src/liuma

CMD [ "supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf" ]



