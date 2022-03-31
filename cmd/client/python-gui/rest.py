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

class NewRequest:

    def __init__(self, context={}):
        self.context = context

    def downloadFile(self, path, filename):
        resp = requests.get(
            f'{self.context["server"]}/{path}?filename={filename}',
            headers={"token": self.context["token"]},
        )
        return resp.content

    def uploadFile(self, path, fileData):
        payload = {}        
        err = None
        resp = requests.post(
            f'{self.context["server"]}/{path}',
            headers={"token": self.context["token"]},
            files=fileData,
        )
        payload = resp.json()
        payload["server"] = self.context["server"]

        try:
            err = payload["error"]
        except:
            pass

        return payload, err

    def postRequest(self, path, tokenre=False):
        payload = {}        
        err = None

        if not tokenre:
            self.context["token"] = ""

        
        resp = requests.post(
            f'{self.context["server"]}/{path}', 
            json=self.context, 
            headers={"token": self.context["token"]},
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

        resp = requests.get(
            f'{self.context["server"]}/{path}?id={self.context["ID"]}&type={self.context["role"]}', 
            headers={"token": self.context["token"]},
        )
        
        if resp.status_code >= 400:
            err = "Status code 404"
            return {}, err

        payload = resp.json()
        
        if type(payload) == "dict":
            payload["server"] = self.context["server"]
        try:
            err = payload["error"]
        except:
            pass

        return payload, err