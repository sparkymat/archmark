FROM golang:1.17-stretch
RUN wget https://github.com/Y2Z/monolith/releases/download/v2.6.1/monolith-gnu-linux-x86_64 -O /bin/monolith
RUN chmod 0755 /bin/monolith
ENV MONOLITH_PATH=/bin/monolith
COPY archmark-linux /bin/archmark
WORKDIR /
EXPOSE 8080
CMD ["/bin/archmark"]
