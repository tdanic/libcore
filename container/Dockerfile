#Ambiente go oracle 
FROM golang:latest
LABEL maintainer="developers@thedataage.com"

#Utilidades
RUN apt-get update && apt-get upgrade -y
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends apt-utils

# Instala las dependencias necesarias
RUN apt-get update && apt-get install -y libaio1
RUN apt-get install build-essential -y
RUN apt-get install alien -y
RUN apt-get install pkg-config -y
RUN apt-get install rpm -y
RUN apt-get install git -y
RUN apt-get install openssh-server -y
RUN apt-get install nano -y

# Descarga la librerias de oracle cliente 12.1 y se instalan
RUN mkdir /usr/auxprogram && \
    cd /usr/auxprogram && \
    git clone https://github.com/tdanic/oracle_client12.git

RUN cd  /usr/auxprogram/oracle_client12 && \
    alien -i *.rpm
RUN cd  /usr/auxprogram/oracle_client12 && \
    rm *.rpm && \
    mkdir -p "/usr/lib/oracle/12.1/client64/network/admin"

# Desinstalacion de librerias no necesarias
RUN apt-get autoremove -y

# Definicion de variables de entorno pkg-config
RUN mkdir -p "/usr/local/share/pkgconfig"
COPY oci8.pc /usr/local/share/pkgconfig/oci8.pc 

# Definicion de variables de entorno de sistema
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
ENV PATH $PATH:$ORACLE_HOME/bin
ENV PKG_CONFIG pkg-config

# Eliminacion el directorio auxiliar
RUN rm -R  /usr/auxprogram

# Configuracion ssh
RUN sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
ENTRYPOINT service ssh start && bash