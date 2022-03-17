import sys 
from PyQt5 import QtWidgets
import requests

import golearn_auth
import golearn_main

class AppMain(QtWidgets.QMainWindow, golearn_main.Ui_MainWindow):
    def __init__(self, args):
        super().__init__()
        self.args = args
        self.setupUi(self)
        self.my_id.setText(str(args["ID"]))
        self.my_token.setText(args["password"])
        self.labelAcc.setText(f"Аккаунт ({args['account']})")
        self.my_role.setText(args['role'])
        
        self.actionExit.triggered.connect(self.ExitProgramm)

    def ExitProgramm(self):
        self.windowAuth = AppAuth()
        self.hide()
        self.windowAuth.show()


class AppAuth(QtWidgets.QMainWindow, golearn_auth.Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)
        self.nextBtn.clicked.connect(self.golearnAuth)
        self.regBtn.clicked.connect(self.golearnReg)
    
    def showMain(self, args):
        self.windowMain = AppMain(args=args)
        self.windowMain.show()

    def postRequest(self, data, path):
        uri = data["server"]
        data.pop("server")
        payload = {}
        err = None
        try:
            resp = requests.post(f'{uri}/{path}', json=data)
            payload = resp.json()
        except Exception as err:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err.args[-1]))
            return {}, str(err.args[-1])

        try:
            err = payload["error"]
        except:
            pass
        return payload, err

    def auth(self, data):
        payload, err = self.postRequest(data, "auth")
        return payload, err

    def registration(self, data):
        payload, err = self.postRequest(data, "user")
        return payload, err

    def golearnAuth(self):
        data = {}
        data["server"] = self.lineServer.text()
        data["account"] = self.lineLogin.text()
        data["password"] = self.linePassword.text()
        roleCheck = data["account"].split(":")
        data["role"] = "user"
        if roleCheck[-1] == "admin":
            data["role"] = roleCheck[-1]
            data["account"] = roleCheck[0]

        payload, err = self.auth(data)
        print(payload)
        if err != None:
            self.label.setText(f"Вход* [{err}]")
            return
        self.showMain(payload)
        self.hide()

    def golearnReg(self):
        data = {}
        data["server"] = self.lineServer.text()
        data["account"] = self.lineLogin.text()
        data["password"] = self.linePassword.text()
        payload, err = self.registration(data)
        if err != None:
            self.label.setText(f"Вход* [{err}]")
            return
        self.showMain(payload)
        self.hide()

def main():
    app = QtWidgets.QApplication(sys.argv) 
    windowAuth = AppAuth()
    windowAuth.show()

    app.exec_()

if __name__ == '__main__':
    main()