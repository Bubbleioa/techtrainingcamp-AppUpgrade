FROM centos:7
COPY grayReleaseProject /root/server
EXPOSE 8080
EXPOSE 11451
CMD /root/server