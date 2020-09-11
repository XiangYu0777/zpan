FROM debian:10

#RUN echo \
##    deb http://mirrors.aliyun.com/debian buster main \
##    deb http://mirrors.aliyun.com/debian buster-updates main \
##    deb http://mirrors.aliyun.com/debian-security buster/updates main \
#    deb http://mirrors.aliyun.com/debian/ buster main non-free contrib \
#    deb-src http://mirrors.aliyun.com/debian/ buster main non-free contrib \
#    deb http://mirrors.aliyun.com/debian-security buster/updates main \
#    deb-src http://mirrors.aliyun.com/debian-security buster/updates main \
#    deb http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib \
#    deb-src http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib \
#    deb http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib \
#    deb-src http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib \
#    > /etc/apt/sources.list

RUN apt-get update \
    && apt-get install --fix-missing -y ca-certificates telnet procps curl

ENV APP_HOME /zpan
WORKDIR $APP_HOME

RUN curl -sSf https://dl.saltbo.cn/install.sh | sh -s zpan

CMD ["zpan", "server"]