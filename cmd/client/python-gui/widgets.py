from PyQt5 import QtWidgets, QtGui
import copy
from datetime import datetime

import rest, requests
from tools import zipper
from lib import profile, actions

const_Actions = {
    "create_course": "Создать курс",
    "delete_course": "Удалить курс"
}

const_success = "Операция прошла успешно"
const_notice = "Уведомление"

class DialogCreateCourse(QtWidgets.QDialog):
    def __init__(self, parent=None, title="title"):
        super().__init__(parent)

        self.setWindowTitle(title)

        QBtn = QtWidgets.QDialogButtonBox.Ok | QtWidgets.QDialogButtonBox.Cancel

        self.buttonBox = QtWidgets.QDialogButtonBox(QBtn)
        self.buttonBox.accepted.connect(self.accept)
        self.buttonBox.rejected.connect(self.reject)
        
        self.layout = QtWidgets.QVBoxLayout()

        self.layout.addWidget(QtWidgets.QLabel("Введите данные для отправки"))
        
        
        self.lineName = QtWidgets.QLineEdit(self)
        self.lineName.setPlaceholderText("Имя")
        self.layout.addWidget(self.lineName)

        self.lineTags = QtWidgets.QLineEdit(self)
        self.lineTags.setPlaceholderText("Теги")
        self.layout.addWidget(self.lineTags)

        self.plainAbout = QtWidgets.QPlainTextEdit(self)
        self.plainAbout.setPlaceholderText("Описание")
        
        self.dataFile = None
        self.pushOpenFile = QtWidgets.QPushButton(self)
        self.pushOpenFile.setText("Прикрепить файл")
        self.pushOpenFile.clicked.connect(self.getCourseFile)
        
        self.layout.addWidget(self.plainAbout) 
        self.layout.addWidget(self.pushOpenFile)
        self.layout.addWidget(self.buttonBox)
        self.setLayout(self.layout)
    
    def getCourseFile(self):
        options = QtWidgets.QFileDialog.Options()
        options |= QtWidgets.QFileDialog.DontUseNativeDialog
        openFile = QtWidgets.QFileDialog.getOpenFileName(self, "Выберите файл")[0]
        if openFile:
            self.dataFile = openFile

    def createCourse(self):
        return {
            "title": self.lineName.text(),
            "about": self.plainAbout.toPlainText(),
            "tags": self.lineTags.text(),
            "file": self.dataFile
        }

class DialogOpenMyCourses(QtWidgets.QDialog):
    def __init__(self, parent=None, courses=[]):
        super().__init__(parent)
        self.courses = courses
        self.setWindowTitle("MyCourses")
        self.layout = QtWidgets.QVBoxLayout()
        self.layout.addWidget(QtWidgets.QLabel("Мои курсы"))
        
        self.items = QtWidgets.QListWidget(self)
        
        for crs in self.courses:
            self.items.addItem(crs["title"])

        self.pushDload = QtWidgets.QPushButton(self)
        self.pushDload.setText("Скачать .zip")
        self.pushDload.clicked.connect(self.downloadCourse)

        self.layout.addWidget(self.items)
        self.layout.addWidget(self.pushDload)
        self.setLayout(self.layout)

    def downloadCourse(self):
        if not self.courses:
            return 
        options = QtWidgets.QFileDialog.Options()
        options |= QtWidgets.QFileDialog.DontUseNativeDialog
        saveFile = QtWidgets.QFileDialog.getSaveFileName(self, 'Сохранить')[0]
        if saveFile:
            # files = []
            # for crs in self.courses:
            #     files.append(
                    
            #     ) TODO
            zipper.Create_zip_file(saveFile, self.courses)

class WidgetActions(QtWidgets.QWidget, actions.Ui_Form):
    def __init__(self, context):
        super().__init__()
        self.setupUi(self)
        self.context = context
        self.req = rest.NewRequest(copy.copy(context))
        self.GetActions()

        print(self.context)

        # Actions
        self.listWidget.itemClicked.connect(lambda item: self.OpenAction(item))

    def GetActions(self):
        self.listWidget.clear()
        for v in const_Actions.values():
            self.listWidget.addItem(v)

    def OpenAction(self, item):
        i = item.text()
        if i == const_Actions["create_course"]:
            dialog = DialogCreateCourse(self, title=i)
            if dialog.exec_():
                data = dialog.createCourse()
                data["server"] = self.context["server"]
                data["token"] = self.context["token"]
                
                self.req = rest.NewRequest()
                self.req.context = data
                
                payload, err = self.req.postRequest("course", tokenre=True)
                if err != None:
                    errorWin = QtWidgets.QErrorMessage(self)
                    errorWin.showMessage(str(err))
                    return
                payloadFileUpload, err = self.req.uploadFile(
                    "course/upload", 
                    fileData={
                        "file": (data["title"], open(data["file"], "rb"))
                        }
                    )
                if err != None:
                    errorWin = QtWidgets.QErrorMessage(self)
                    errorWin.showMessage(str(err))
                    return
                print(payload, payloadFileUpload)
                okWin = QtWidgets.QMessageBox.about(self, const_notice, const_success)

        elif i == const_Actions["delete_course"]:
            pass
        else:
            return



class WidgetMyProfile(QtWidgets.QWidget, profile.Ui_Form):
    def __init__(self, context):
        super().__init__()
        self.setupUi(self)
        self.context = context
        self.req = rest.NewRequest(copy.copy(context))

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
        
        self.req.context = {**self.req.context, **user}
        print("CONTEXT >> ", self.req.context)
        payload, err = self.req.postRequest("user/update", tokenre=True)
        if err != None:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err))
            return
        print(payload)
        okWin = QtWidgets.QMessageBox.about(self, const_notice, const_success)
    
    def DeleteUserProfile(self):
        payload, err = self.req.postRequest("user/delete", tokenre=True)
        if err != None:
            errorWin = QtWidgets.QErrorMessage(self)
            errorWin.showMessage(str(err))
            return
        print(payload)
        okWin = QtWidgets.QMessageBox.about(self, const_notice, const_success)
        self.close()
