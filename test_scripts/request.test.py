import requests as req


def get_user_data():
    res = req.get("https://random-data-api.com/api/v2/users").json()
    user_ob = {
        "id": res["id"],
        "name": res["first_name"] + " " + res["last_name"],
        "email": res["email"],
        "password": res["password"],
    }
    # req.post("http://localhost:5000/users", json=user_ob)
    return user_ob


def query_user_api():
    user = get_user_data()
    # print("User: ")
    # print(user)
    res = req.post("http://localhost:5000/users", json=user)
    print(res.json())
    return user


def query_posts_api(user):
    post = {
        "id": 0,
        "title": "Test Title",
        "body": "Test Body",
        "user": user,
        "timestamp": 10,
    }
    res = req.post("http://localhost:5000/posts", json=post)
    print(res.json())


user = query_user_api()
query_posts_api(user)
