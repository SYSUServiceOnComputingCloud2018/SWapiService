FROM guyton/centos6
RUN yum install -y mysql mysql-server
RUN /etc/init.d/mysqld start &&\
mysql -e "grant all privileges on *.* to 'root'@'%' identified by '123456';" &&\
mysql -e "grant all privileges on *.* to 'root'@'localhost' identified by '123456';"
EXPOSE 3306
CMD ["mysqld_safe"]
