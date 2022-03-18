import sys 
from PyQt5 import QtWidgets, QtGui, QtCore
import requests
from datetime import datetime

from lib import profile, golearn_auth, golearn_main

def openGithub():
    QtGui.QDesktopServices.openUrl(QtCore.QUrl("https://github.com/rombintu/golearn"))

def openNoAccess():
    QtGui.QDesktopServices.openUrl(QtCore.QUrl("https://google.com")) # TODO

class WidgetMyProfile(QtWidgets.QWidget, profile.Ui_Form):
    def __init__(self, profileData):
        super().__init__()
        self.setupUi(self)
        self.profileData = profileData
        print(self.profileData)
        self.my_id.setText(str(profileData["ID"]))
        self.my_token.setText(profileData["password"])
        self.labelAcc.setText(f"Аккаунт ({profileData['account']})")
        self.my_role.setText(profileData['role'])

        self.lineFName.setText(profileData["first_name"])
        self.lineLName.setText(profileData["last_name"])
        self.lineAddr.setText(profileData["address"])
        self.lineMail.setText(profileData["mail"])
        self.linePhone.setText(profileData["phone"])
        self.lineBthday.setText(str(profileData["date_of_birth"]))

        
class AppMain(QtWidgets.QMainWindow, golearn_main.Ui_MainWindow):
    def __init__(self, args):
        super().__init__()
        self.args = args
        self.setupUi(self)

        # self.model = QtGui.QStandardItemModel()
        # self.listAudit.setModel(self.model)
        # self.logger = self.listAudit.addI
        if self.args["role"] == "user":
            self.pushOpenAdmin.hide()
        # BTNS
        self.pushMyProfile.clicked.connect(self.OpenMyProfile)

        # MENU
        self.actionExit.triggered.connect(self.ExitProgramm)
        self.actionby_Nickolsky.triggered.connect(openGithub)
    
    def Log(self, item, err=False):
        # self.model.appendRow(QtGui.QStandardItem(item))
        flag = ""
        if err:
            flag = "ERROR"
        time = datetime.now().strftime("%H:%M:%S")
        self.plainTextEdit.appendPlainText(f"[{time}] {flag} {item}")

    def OpenMyProfile(self):
        params={"id": self.args["ID"], "type": self.args["role"]}, 
        headers={"token": self.args["password"]}
        payload, err = self.getRequest(self.args, "user", params, headers)
        if err != None:
            self.Log(err, True)
            return
        self.profileWidjet = WidgetMyProfile(payload)
        self.profileWidjet.show()
        self.Log("Запрос данных аккаунта")

    def ExitProgramm(self):
        self.windowAuth = AppAuth()
        self.close()
        self.windowAuth.show()

    def getRequest(self, data, path, params, headers):
        uri = data["server"]
        # data.pop("server")
        payload = {}
        err = None
        try:
            resp = requests.get(
                f'{uri}/{path}?id={data["ID"]}&type={data["role"]}', 
                # params=params,
                headers=headers,
            )
            payload = resp.json()
        except Exception as err:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err))
            return {}, str(err)

        try:
            err = payload["error"]
        except:
            pass
        return payload, err

class AppAuth(QtWidgets.QMainWindow, golearn_auth.Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)
        self.nextBtn.clicked.connect(self.golearnAuth)
        self.regBtn.clicked.connect(self.golearnReg)

        self.actionGithub.triggered.connect(openGithub)
        self.actionNoAccess.triggered.connect(openNoAccess) # TODO

    def showMain(self, args):
        self.windowMain = AppMain(args=args)
        self.windowMain.show()

    def postRequest(self, data, path):
        uri = data["server"]
        # data.pop("server")
        payload = {}
        err = None
        try:
            resp = requests.post(f'{uri}/{path}', json=data)
            payload = resp.json()
            payload["server"] = data["server"]
        except Exception as err:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err))
            return {}, str(err)

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
        self.close()

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
    # windowAuth.show()

    # DEVELOPMENT
    data = {"server": "http://localhost:5000", "account": "admin", "password": "admin", "role": "admin", }
    # data = {"server": "http://localhost:5000", "account": "user1", "password": "user1", "role": "user", }
    payload, err = windowAuth.auth(data)
    print(payload)
    if err != None:
        self.label.setText(f"Вход* [{err}]")
        return
    windowAuth.showMain(payload)
    windowAuth.close()

    app.exec_()

if __name__ == '__main__':
    main()