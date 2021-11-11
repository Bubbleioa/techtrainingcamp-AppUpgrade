FROM ubuntu:latest
COPY techtrainingcamp-AppUpgrade /root/server
COPY ./public/index.html /root/public/index.html
COPY ./redis.conf /root/redis.conf
EXPOSE 8080
EXPOSE 11451
ENV IS_DOCKER 1
RUN apt-get update
RUN apt-get install -y mysql-server
RUN apt-get install -y redis 
RUN systemctl start mysql
RUN mysql -e "CREATE DATABASE app;"
RUN mysql -e "CREATE USER 'test'@'127.0.0.1' IDENTIFIED BY '123456';"
RUN mysql -e "grant all privileges on *.* to 'test'@'127.0.0.1' identified by '123456' WITH GRANT OPTION ;"
RUN mysql -u test -p123456 -e "use app; CREATE TABLE rules(id int UNSIGNED AUTO_INCREMENT,aid INT UNSIGNED,hit_count INT UNSIGNED DEFAULT 0,download_count INT UNSIGNED DEFAULT 0,  platform CHAR(16),download_url VARCHAR(128),update_version_code	VARCHAR(128),device_list TEXT,md5	VARCHAR(128),max_update_version_code	VARCHAR(128),min_update_version_code	VARCHAR(128),max_os_api	TINYINT UNSIGNED,min_os_api	TINYINT UNSIGNED,cpu_arch	TINYINT UNSIGNED,channel	VARCHAR(128),title	VARCHAR(256),update_tips	VARCHAR(1024),enabled	BOOLEAN DEFAULT true,create_date DATETIME DEFAULT CURRENT_TIMESTAMP,PRIMARY KEY ( id ));"  
RUN redis-server --version 
RUN redis-server /root/redis.conf
CMD /root/server
