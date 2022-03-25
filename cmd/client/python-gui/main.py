import sys 
from PyQt5 import QtWidgets, QtGui, QtCore
import requests, copy
from datetime import datetime

import rest

from lib import profile, golearn_auth, golearn_main

def openGithub():
    QtGui.QDesktopServices.openUrl(QtCore.QUrl("https://github.com/rombintu/golearn"))

def openNoAccess():
    QtGui.QDesktopServices.openUrl(QtCore.QUrl("https://google.com")) # TODO

class WidgetMyProfile(QtWidgets.QWidget, profile.Ui_Form):
    def __init__(self, context):
        super().__init__()
        self.setupUi(self)
        self.context = context
        self.req = rest.NewRequest(context)

        print(self.context)

        self.my_id.setText(str(context["ID"]))
        self.my_token.setText(context["token"])
        self.labelAcc.setText(f"Аккаунт ({context['account']})")
        self.my_role.setText(context['role'])

        self.lineFName.setText(context["first_name"])
        self.lineLName.setText(context["last_name"])
        self.lineAddr.setText(context["address"])
        self.lineMail.setText(context["mail"])
        self.linePhone.setText(context["phone"])
        self.lineBthday.setText(str(context["date_of_birth"]))

        self.refreshcontextBtn.clicked.connect(self.UpdateUserProfile)
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

        self.req.context = {**self.req.context, **user}

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
    def __init__(self, data):
        super().__init__()
        self.setupUi(self)
        self.context = data

        copy_data = copy.copy(data)
        self.req = rest.NewRequest(copy_data)
        self.courses = []
        self.GetAllCourses()

        print("PROFILE >> ", self.context)
        # clicked = QtCore.pyqtSignal()

        if self.context["role"] == "user":
            self.pushOpenAdmin.hide()
        # BTNS
        self.pushMyProfile.clicked.connect(self.OpenMyProfile)
        self.pushRefreshCourses.clicked.connect(self.GetAllCourses)

        # MENU
        self.actionExit.triggered.connect(self.Logout)
        self.actionby_Nickolsky.triggered.connect(openGithub)

        # Courses
        self.listWidget.currentRowChanged.connect(lambda i: self.ViewCourse(i))

    def GetAllCourses(self):
        self.Log(message="Обновление курсов")
        payload, err = self.req.getRequest("course/all")
        if err != None:
            self.Log(err, True)
            return
        self.courses = payload

        self.listWidget.clear()
        self.plainTextAboutCourse.clear()
        self.lineTags.clear()

        if self.courses:
            for crs in self.courses:
                self.listWidget.addItem(crs["title"])
        else:
            self.listWidget.addItem("Курсы не найдены, попробуйте обновить")
        self.Log(message="Курсы обновлены")
    
    def ViewCourse(self, index):
        curItem = self.courses[index]
        if curItem["is_active"]:
            self.pushGetMeCourse.setEnabled()
        self.lineTags.setText(curItem["tags"])
        self.plainTextAboutCourse.setPlainText(curItem["about"])

    def Log(self, message, err=False):
        flag = ""
        if err:
            flag = "ERROR"
        time = datetime.now().strftime("%H:%M:%S")
        self.plainTextEdit.appendPlainText(f"[{time}] {flag} {message}")

    def OpenMyProfile(self):
        payload, err = self.req.getRequest("user", tokenre=True)
        if err != None:
            self.Log(err, True)
            return
        # payload["token"] = self.context["token"]

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

    def showMain(self, context):
        print("SHOW MAIN >> ", context)
        self.windowMain = AppMain(context)
        self.windowMain.show()

    def SkipAuth(self, context):
        self.req.context = context
        payload, err = self.req.postRequest("auth")
        print("SKIP CONTEXT >> ", payload)
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

        self.req.context = data
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

        self.req.context = data
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