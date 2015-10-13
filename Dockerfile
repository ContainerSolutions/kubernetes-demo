FROM scratch

EXPOSE 8080

COPY kubernetes-demo /

ENTRYPOINT ["/kubernetes-demo"]
