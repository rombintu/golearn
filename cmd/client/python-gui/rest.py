import requests

class NewRequest:

    def __init__(self, data={}):
        self.data = data

    def postRequest(self, path, tokenre=False):
        payload = {}
        err = None

        if not tokenre:
            self.data["token"] = ""

        resp = requests.post(
            f'{self.data["server"]}/{path}', 
            json=self.data, 
            headers={"token": self.data["token"]}
        )
        payload = resp.json()
        payload["server"] = self.data["server"]

        try:
            err = payload["error"]
        except:
            pass

        return payload, err

    def getRequest(self, path, tokenre=False):
        payload = {}
        err = None        

        if not tokenre:
            self.data["token"] = ""


        resp = requests.get(
            f'{self.data["server"]}/{path}?id={self.data["ID"]}&type={self.data["role"]}', 
            headers={"token": self.data["token"]},
        )
        payload = resp.json()
        
        if type(payload) == "dict":
            payload["server"] = self.data["server"]
        try:
            err = payload["error"]
        except:
            pass

        return payload, err