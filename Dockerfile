FROM ubuntu:20.04

COPY registry /home
COPY config.json /home
WORKDIR /home

# CMD ["java", "-jar", "backend-0.0.1.jar", "--spring.profiles.active=online"]
