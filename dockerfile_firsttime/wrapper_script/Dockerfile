FROM centos:latest
RUN yum install -y httpd iproute && yum clean all
RUN echo "Hello Apache." > /var/www/html/index.html
COPY ./script/httpdstart.sh  /usr/local/bin
RUN  chmod 755 /usr/local/bin/httpdstart.sh
ENTRYPOINT [ "/usr/local/bin/httpdstart.sh" ]