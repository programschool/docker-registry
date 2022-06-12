FROM ubuntu:20.04

RUN rm -f /etc/localtime
RUN ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN mkdir /home/logs
COPY registry /home
COPY config.json /home
WORKDIR /home

# CMD ["java", "-jar", "backend-0.0.1.jar", "--spring.profiles.active=online"]
