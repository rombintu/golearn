import requests

def getToken(creds):
    resp = requests.post(f'{creds["server"]}/auth', data=creds)
    payload = resp.json()
    err = None
    try:
        err = payload["error"]
    except:
        pass

    return payload, err

print(getToken({"server": "http://localhost:5000", "account": "admin", "password": "admin", "role": "admin", }))
class NewRequest:

    def __init__(self, context={}):
        self.context = context


    def postRequest(self, path, tokenre=False):
        payload = {}
        err = None

        if not tokenre:
            self.context["token"] = ""

        resp = requests.post(
            f'{self.context["server"]}/{path}', 
            json=self.context, 
            headers={"token": self.context["token"]}
        )
        payload = resp.json()
        payload["server"] = self.context["server"]

        try:
            err = payload["error"]
        except:
            pass

        return payload, err

    def getRequest(self, path, tokenre=False):
        payload = {}
        err = None        

        if not tokenre:
            self.context["token"] = ""

        # print("GET >>", self.context)
        resp = requests.get(
            f'{self.context["server"]}/{path}?id={self.context["ID"]}&type={self.context["role"]}', 
            headers={"token": self.context["token"]},
        )
        payload = resp.json()
        
        if type(payload) == "dict":
            payload["server"] = self.context["server"]
        try:
            err = payload["error"]
        except:
            pass

        return payload, err