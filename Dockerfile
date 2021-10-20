FROM golang:1.17-stretch
COPY archmark-linux /bin/archmark
WORKDIR /
EXPOSE 8080
CMD ["/bin/archmark"]
