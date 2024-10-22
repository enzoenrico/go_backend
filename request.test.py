import requests

r = requests.post(
    "http://localhost:5000/users",
    json={"name": "testuser", "password": "testpassword", "email": "testemail"},
)
print(r.text)
