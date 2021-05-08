FROM ubuntu:20.04

RUN rm -f /etc/localtime
RUN ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY registry /home
COPY config.json /home
RUN mkdir /home/logs
WORKDIR /home

# CMD ["java", "-jar", "backend-0.0.1.jar", "--spring.profiles.active=online"]
