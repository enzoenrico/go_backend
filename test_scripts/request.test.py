import requests as req

def get_user_data():
    res = req.get("https://random-data-api.com/api/v2/users").json()
    user_ob = {
        "id": res["id"],
        "name": res["first_name"] + " " + res["last_name"],
        "email": res["email"],
        "password": res["password"],
    }
    return user_ob


def query_api():
    user = get_user_data()
    res = req.post("http://localhost:5000/users", json=user)
    print(res.json())

query_api()