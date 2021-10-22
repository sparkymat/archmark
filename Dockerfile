FROM golang:1.17-stretch
COPY . /app
WORKDIR /app
RUN make archmark-linux
RUN mv archmark-linux /bin/archmark
EXPOSE 8080
CMD ["/bin/archmark"]
