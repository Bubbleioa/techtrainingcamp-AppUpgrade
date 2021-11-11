FROM centos:7
COPY techtrainingcamp-AppUpgrade /root/server
COPY ./public/index.html /root/public/index.html
EXPOSE 8080
EXPOSE 11451
ENV IS_DOCKER 1
RUN /etc/init.d/mysqld start &&\  
    mysql -e "grant all privileges on *.* to 'root'@'%' identified by '123456' WITH GRANT OPTION ;"&&\  
    mysql -e "grant all privileges on *.* to 'root'@'localhost' identified by '123456' WITH GRANT OPTION ;"&&\ 
    mysql -u root -p123456 -e "show databases;"  
CMD /root/server
