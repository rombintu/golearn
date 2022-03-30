from PyQt5 import QtWidgets
import copy
from datetime import datetime

import rest

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
        self.layout.addWidget(self.plainAbout) 
        self.layout.addWidget(self.buttonBox)

        self.setLayout(self.layout)

    def createCourse(self):
        return {
            "title": self.lineName.text(),
            "about": self.plainAbout.toPlainText(),
            "tags": self.lineTags.text()
        }

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
                print(payload)
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
