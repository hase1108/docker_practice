FROM centos:latest
RUN yum install -y epel-release
RUN yum install -y supervisor httpd iproute && yum clean all
RUN echo "Hello Apache." > /var/www/html/index.html
COPY supervisord.conf /etc/
ENTRYPOINT ["/usr/bin/supervisord"]