FROM centos:7
COPY techtrainingcamp-AppUpgrade /root/server
COPY ./public/index.html /root/public/index.html
EXPOSE 8080
EXPOSE 11451
RUN ls /root
CMD /root/server
