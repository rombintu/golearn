import sys 
from PyQt5 import QtWidgets, QtGui, QtCore
import requests
from datetime import datetime

import rest

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
        self.req = rest.NewRequest(profileData)

        print(self.profileData)

        self.my_id.setText(str(profileData["ID"]))
        self.my_token.setText(profileData["token"])
        self.labelAcc.setText(f"Аккаунт ({profileData['account']})")
        self.my_role.setText(profileData['role'])

        self.lineFName.setText(profileData["first_name"])
        self.lineLName.setText(profileData["last_name"])
        self.lineAddr.setText(profileData["address"])
        self.lineMail.setText(profileData["mail"])
        self.linePhone.setText(profileData["phone"])
        self.lineBthday.setText(str(profileData["date_of_birth"]))

        self.refreshDataBtn.clicked.connect(self.UpdateUserProfile)
        self.deleteAccountBtn.clicked.connect(self.DeleteUserProfile)

    def UpdateUserProfile(self):
        user = {
            "first_name": self.lineFName.text(),
            "last_name": self.lineLName.text(),
            "address": self.lineAddr.text(),
            "mail": self.lineMail.text(),
            "phone": self.linePhone.text(),
            "date_of_birth": self.lineBthday.text(),
        }

        self.req.data = {**self.req.data, **user}

        payload, err = self.req.getRequest("user/update", tokenre=True)
        if err != None:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err))
            return
        print(payload)
        okWin = QtWidgets.QMessageBox.about(self, "Уведомление", "Данные успешно обновлены")
    
    def DeleteUserProfile(self):
        payload, err = self.req.postRequest("user/delete", tokenre=True)
        if err != None:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err))
            return
        print(payload)
        okWin = QtWidgets.QMessageBox.about(self, "Уведомление", "Аккаунт удален")
        self.close()

class AppMain(QtWidgets.QMainWindow, golearn_main.Ui_MainWindow):
    def __init__(self, args):
        super().__init__()
        self.args = args
        self.setupUi(self)

        self.req = rest.NewRequest(args)
        
        if self.args["role"] == "user":
            self.pushOpenAdmin.hide()
        # BTNS
        self.pushMyProfile.clicked.connect(self.OpenMyProfile)

        # MENU
        self.actionExit.triggered.connect(self.Logout)
        self.actionby_Nickolsky.triggered.connect(openGithub)
    
    def Log(self, item, err=False):
        flag = ""
        if err:
            flag = "ERROR"
        time = datetime.now().strftime("%H:%M:%S")
        self.plainTextEdit.appendPlainText(f"[{time}] {flag} {item}")

    def OpenMyProfile(self):
        
        payload, err = self.req.getRequest("user", tokenre=True)
        if err != None:
            self.Log(err, True)
            return
        payload["token"] = self.args["token"]
        self.profileWidjet = WidgetMyProfile(payload)
        self.profileWidjet.show()
        self.Log("Запрос данных аккаунта")

    def Logout(self):
        self.windowAuth = AppAuth()
        self.close()
        self.windowAuth.show()

    
class AppAuth(QtWidgets.QMainWindow, golearn_auth.Ui_MainWindow):
    def __init__(self):
        super().__init__()
        self.setupUi(self)
        self.nextBtn.clicked.connect(self.Auth)
        self.regBtn.clicked.connect(self.Reg)

        self.req = rest.NewRequest()

        self.actionGithub.triggered.connect(openGithub)
        self.actionNoAccess.triggered.connect(openNoAccess) # TODO

    def showMain(self, args):
        self.windowMain = AppMain(args=args)
        self.windowMain.show()

    def SkipAuth(self, data):
        self.req.data = data
        payload, err = self.req.postRequest("auth")
        return payload, err

    def Auth(self):
        data = {}
        data["server"] = self.lineServer.text()
        data["account"] = self.lineLogin.text()
        data["password"] = self.linePassword.text()
        roleCheck = data["account"].split(":")
        data["role"] = "user"
        if roleCheck[-1] == "admin":
            data["role"] = roleCheck[-1]
            data["account"] = roleCheck[0]

        self.req.data = data
        payload, err = self.req.postRequest("auth")

        print(payload)
        
        if err != None:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err))
            return
        self.showMain(payload)
        self.close()

    def Reg(self):
        data = {}
        data["server"] = self.lineServer.text()
        data["account"] = self.lineLogin.text()
        data["password"] = self.linePassword.text()

        self.req.data = data
        payload, err = self.req.postRequest("user")

        if err != None:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err))
            return
        
        okWin = QtWidgets.QMessageBox.about(self, "Уведомление", "Регистрация прошла успешно")

def main():
    app = QtWidgets.QApplication(sys.argv) 
    windowAuth = AppAuth()
    # windowAuth.show()

    # DEVELOPMENT
    data = {"server": "http://localhost:5000", "account": "admin", "password": "admin", "role": "admin", }
    # data = {"server": "http://localhost:5000", "account": "user1", "password": "user1", "role": "user", }
    payload, err = windowAuth.SkipAuth(data)
    print(payload)
    if err != None:
        print(err)
    windowAuth.showMain(payload)
    windowAuth.close()

    app.exec_()

if __name__ == '__main__':
    main()