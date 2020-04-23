FROM centos:7.4.1708

#安装YUM源
RUN rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
RUN rpm -Uvh https://mirror.webtatic.com/yum/el7/webtatic-release.rpm

RUN yum -y install unzip
WORKDIR /usr/local/bin
ADD consul_1.6.1_linux_amd64.zip . 
RUN unzip consul_1.6.1_linux_amd64.zip

EXPOSE 8500
ENTRYPOINT ["/usr/local/bin/consul","agent","-dev","-client","0.0.0.0"]
