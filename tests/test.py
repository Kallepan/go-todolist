import requests
import json
URL = "http://localhost:8080/tasks"

def test_post():
    data = json.dumps({
        "name": "Create a new task",
        "description": "Create a new task in the database",
        "completed": False,
    })


    r = requests.post(URL, data=data, headers={"Content-Type": "application/json"})

    assert r.status_code == 201


def test_delete():
    r = requests.delete(URL + "/2")
    assert r.status_code == 200

def main():
    #test_post()
    test_delete()


if __name__ == "__main__":
    main()