FROM centos:7
COPY techtrainingcamp-AppUpgrade /root/server
COPY /public/* /root/server/public
EXPOSE 8080
EXPOSE 11451
CMD /root/server