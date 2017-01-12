import time
from flask import Flask, jsonify
from faker import Factory

app = Flask(__name__)
fake = Factory.create()


def query_data():
    time.sleep(0.1)
    return fake.name()


@app.route("/")
def root():
    return jsonify({ "data": query_data() })


if __name__ == '__main__':
    app.run()

