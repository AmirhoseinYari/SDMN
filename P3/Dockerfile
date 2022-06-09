# use python as base image
FROM python:3.8.10

#
ENV PYTHONUNBUFFERED 1

# set the work directory
WORKDIR /code

# copy script.sh from host to /root container
COPY HTTP_server.py /code/

# run the script on startup
CMD ["python", "HTTP_server.py"]
