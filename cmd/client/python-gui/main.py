import sys 
from PyQt5 import QtWidgets, QtGui, QtCore
import copy
from datetime import datetime

import rest, widgets

from lib import golearn_auth, golearn_main

def openGithub():
    QtGui.QDesktopServices.openUrl(QtCore.QUrl("https://github.com/rombintu/golearn"))

def openNoAccess():
    QtGui.QDesktopServices.openUrl(QtCore.QUrl("https://google.com")) # TODO

class AppMain(QtWidgets.QMainWindow, golearn_main.Ui_MainWindow):
    def __init__(self, data):
        super().__init__()
        self.setupUi(self)
        self.context = data
        self.req = rest.NewRequest(copy.copy(data))
        self.courses = []
        self.GetAllCourses()

        print("PROFILE >> ", self.context)
        
        print("COURSES >> ", self.courses)
        if self.context["role"] == "user":
            self.pushOpenAdmin.hide()
        # BTNS
        self.pushMyProfile.clicked.connect(self.OpenMyProfile)
        self.pushOpenAdmin.clicked.connect(self.OpenActions)
        self.pushRefreshCourses.clicked.connect(self.GetAllCourses)
        self.pushMyCourses.clicked.connect(self.OpenMyCourses)

        # MENU
        self.actionExit.triggered.connect(self.Auditout)
        self.actionby_Nickolsky.triggered.connect(openGithub)

        # Courses
        self.listWidget.currentRowChanged.connect(lambda i: self.ViewCourse(i))

    def GetAllCourses(self):
        self.Audit(message="Обновление курсов")
        payload, err = self.req.getRequest("course/all")
        if err != None:
            self.Audit(err, True)
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
        self.Audit(message="Курсы обновлены")
    
    def ViewCourse(self, index):
        try:
            curItem = self.courses[index]
        except:
            return
        if curItem["is_active"]:
            self.pushGetMeCourse.setEnabled()
        self.lineTags.setText(curItem["tags"])
        self.plainTextAboutCourse.setPlainText(curItem["about"])

    def Audit(self, message, err=False):
        flag = ""
        if err:
            flag = "ERROR"
        time = datetime.now().strftime("%H:%M:%S")
        self.plainTextEdit.appendPlainText(f"[{time}] {flag} {message}")

    def OpenMyProfile(self):
        self.req.context = self.context
        payload, err = self.req.getRequest("user", tokenre=True)
        if err != None:
            self.Audit(err, True)
            return
        payload["token"] = self.context["token"]
        payload["server"] = self.context["server"]

        self.profileWidjet = widgets.WidgetMyProfile(payload)
        self.profileWidjet.show()
        self.Audit("Запрос данных аккаунта")

    def OpenMyCourses(self):
        dialog = widgets.DialogOpenMyCourses(self, courses=self.courses)
        dialog.exec_()

    def OpenActions(self):
        self.adminWidget = widgets.WidgetActions(self.context)
        self.adminWidget.show()
        self.Audit('Режим: "Администрирование"')

    def Auditout(self):
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