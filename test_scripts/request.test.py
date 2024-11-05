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

def get_token():
    res = req.post("http://localhost:5000/login", json={"name": "asdasd"})
    print(res.json())
    return res.json()



def query_user_api():
    user = get_user_data()
    # print("User: ")
    # print(user)
    res = req.post("http://localhost:5000/users", json=user, headers={"Authorization": "Bearer " + get_token()['token']}) 
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


get_token()
user = query_user_api()
query_posts_api(user)
