FROM golang:1.17-stretch
COPY . /app
WORKDIR /app
RUN make archmark-linux
RUN mv archmark-linux /bin/archmark
ENV MONOLITH_PATH=/bin/monolith
EXPOSE 8080
CMD ["/bin/archmark"]
