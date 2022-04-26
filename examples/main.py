import requests
from os import path

pathfile = path.join(path.dirname(path.abspath(__file__)), "temp.txt")
server = "http://localhost:8080"


def set():
    data = {
        'file': (pathfile, open(pathfile, 'rb'))
    }
    response = requests.post(f'{server}/setFile', files=data)
    return response.json()


def get(uuid_number):
    response = requests.post(f'{server}/getFile', json={"name": uuid_number})
    return response.content


# create file
with open('temp.txt', 'w') as f:
    f.write('Hello World!')


result = set()
print(result["name"])
print(get(result["name"]))
