FROM golang:1.22-bookworm

RUN useradd --create-home --shell /bin/bash jenkins
USER jenkins

RUN mkdir /home/jenkins/project
WORKDIR /home/jenkins/project

COPY ./go/* /home/jenkins/project
RUN go build -ldflags "-s -w"

CMD [ "/home/jenkins/project/go-vs-php" ]
