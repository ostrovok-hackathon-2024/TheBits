FROM python:3.10

WORKDIR /app
COPY requirements.txt /app
RUN pip install -r requirements.txt

COPY . /app
ENTRYPOINT [ "python" ]
CMD [ "flask_app.py" ]  

EXPOSE 5005