FROM sagemath/sagemath:latest-py3

ARG DEBIAN_FRONTEND=noninteractive
#RUN sed -i 's/archive.ubuntu.com/linux.yz.yamagata-u.ac.jp/g' /etc/apt/sources.list

# COPY ./requirements.txt /opt/requirements.txt
# RUN pip3 install -r /opt/requirements.txt

CMD [ "/opt/proxy-cmd", "/opt/app", ":9999", "python3", "affine.py" ]
