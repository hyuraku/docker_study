FROM python:3.8.4-slim
ENV PYTHONUNBUFFERED 1
RUN mkdir /code
WORKDIR /code
ADD requirements.txt /code/
RUN apt-get update && apt-get install python3-psycopg2 libghc-postgresql-libpq-dev -y
RUN pip install -r requirements.txt
ADD . /code/
