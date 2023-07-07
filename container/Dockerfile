#Ambiente go oracle 
FROM golang
LABEL maintainer="developers@thedataage.com"

#Utilidades
RUN apt-get update 
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends apt-utils

#dependencia de la imagen
RUN apt-get install vim -y && \
    apt-get install curl -y && \
    apt-get install wget -y && \
    apt-get install gzip -y && \
    apt-get install unzip -y && \
    apt-get install zip -y && \
    apt-get install p7zip-full -y && \
    apt-get install inotify-tools -y && \
    apt-get install software-properties-common -y && \
    apt-get install libaio1 -y && \
    apt-get install alien -y 

#instalacion de rpm
RUN apt-get install rpm -y

#instalacion de git
RUN apt-get install git -y

#oracle cli 12.1
RUN mkdir /usr/auxprogram && \
    cd /usr/auxprogram && \
    git clone https://github.com/tdanic/dockerglash.git

RUN cd  /usr/auxprogram/dockerglash && \
    7z e jdkclient.7z.001 
RUN cd  /usr/auxprogram/dockerglash && \
    rm -R jdkclient.7z.0* 
RUN cd  /usr/auxprogram/dockerglash && \
    alien -i *.rpm

RUN cd  /usr/auxprogram/dockerglash && \
    rm *.rpm && \
    rm -R Comprimer && \
    rm README.md && \
    mkdir -p "/usr/lib/oracle/12.1/client64/network/admin"

RUN rm -R  /usr/auxprogram

#creacion de variable oci
RUN mkdir -p "/usr/local/share/pkgconfig"
COPY configssh/oci8.pc /usr/local/share/pkgconfig/oci8.pc 

#variables de entorno
ENV PKG_CONFIG_PATH /usr/local/share/pkgconfig
ENV ORACLE_HOME /usr/lib/oracle/12.1/client64
ENV LD_LIBRARY_PATH $LD_LIBRARY_PATH:$ORACLE_HOME/lib:$ORACLE_HOME
ENV OCI_HOME /usr/lib/oracle/12.1/client64/lib
ENV OCI_LIB_DIR $ORACLE_HOME/lib
ENV OCI_INC_DIR /usr/include/oracle/12.1/client64
ENV OCI_INCLUDE_DIR /usr/include/oracle/12.1/client64
ENV OCI_VERSION 12

ENV C_INCLUDE_PATH /usr/include:/usr/include/oracle/12.1/client64:/usr/local/include
ENV TNS_ADMIN $ORACLE_HOME/network/admin
ENV TNSADMIN $ORACLE_HOME/network/admin
ENV GOBIN $GOPATH/bin
ENV GCCGO gccgo
ENV CC gcc
ENV CXX g++
ENV CGO_ENABLED 1
ENV PKG_CONFIG pkg-config
ENV GO111MODULE on
ENV GIT_TERMINAL_PROMPT 1
ENV GOPROXY direct
ENV GOSUMDB off
ENV  PATH $PATH:$ORACLE_HOME/bin

#depuracion de la imagen
RUN  apt-get remove --purge unzip -y && \
     apt-get remove --purge p7zip-full -y && \    
     apt-get remove --purge alien -y && \
     apt-get remove --purge rpm -y 

RUN  apt-get autoremove -y

#configuracion ssh
RUN apt-get update
RUN apt-get install sudo -y

USER root
ENV  SSH_MASTER_USER master
ENV  SSH_MASTER_PASS abc123

RUN apt-get update && apt-get install -y openssh-server
RUN mkdir /var/run/sshd
RUN echo 'root:abc123' | chpasswd
RUN sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd
EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]