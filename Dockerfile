FROM scratch
MAINTAINER Ondrej Barta <ondrej@ondrej.it>

ADD gwets /

EXPOSE 8080

ENTRYPOINT ["/gwets"]
