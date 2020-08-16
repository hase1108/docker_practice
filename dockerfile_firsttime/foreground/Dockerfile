FROM centos:latest
RUN yum install -y httpd iproute && yum clean all
RUN echo "Hello Apache." > /var/www/html/index.html
RUN systemctl enable httpd
ENTRYPOINT [ "/usr/sbin/httpd","-DFOREGROUND"]
