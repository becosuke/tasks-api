FROM scratch
ARG name
COPY ${name} /main
COPY zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
COPY ca-certificates.crt /ca-certificates.crt
ENV SSL_CERT_FILE /ca-certificates.crt
CMD ["/main"]
