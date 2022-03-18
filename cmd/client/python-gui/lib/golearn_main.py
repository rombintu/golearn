# -*- coding: utf-8 -*-

# Form implementation generated from reading ui file 'cmd/client/python-gui/ui/main_window.ui'
#
# Created by: PyQt5 UI code generator 5.15.6
#
# WARNING: Any manual changes made to this file will be lost when pyuic5 is
# run again.  Do not edit this file unless you know what you are doing.


from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_MainWindow(object):
    def setupUi(self, MainWindow):
        MainWindow.setObjectName("MainWindow")
        MainWindow.resize(1091, 579)
        self.centralwidget = QtWidgets.QWidget(MainWindow)
        self.centralwidget.setObjectName("centralwidget")
        self.gridLayout = QtWidgets.QGridLayout(self.centralwidget)
        self.gridLayout.setObjectName("gridLayout")
        self.frame_4 = QtWidgets.QFrame(self.centralwidget)
        sizePolicy = QtWidgets.QSizePolicy(QtWidgets.QSizePolicy.Expanding, QtWidgets.QSizePolicy.Minimum)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.frame_4.sizePolicy().hasHeightForWidth())
        self.frame_4.setSizePolicy(sizePolicy)
        self.frame_4.setFrameShape(QtWidgets.QFrame.StyledPanel)
        self.frame_4.setFrameShadow(QtWidgets.QFrame.Raised)
        self.frame_4.setObjectName("frame_4")
        self.horizontalLayout = QtWidgets.QHBoxLayout(self.frame_4)
        self.horizontalLayout.setObjectName("horizontalLayout")
        self.pushRefreshCourses = QtWidgets.QPushButton(self.frame_4)
        sizePolicy = QtWidgets.QSizePolicy(QtWidgets.QSizePolicy.Minimum, QtWidgets.QSizePolicy.Minimum)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.pushRefreshCourses.sizePolicy().hasHeightForWidth())
        self.pushRefreshCourses.setSizePolicy(sizePolicy)
        self.pushRefreshCourses.setMaximumSize(QtCore.QSize(228, 52))
        self.pushRefreshCourses.setObjectName("pushRefreshCourses")
        self.horizontalLayout.addWidget(self.pushRefreshCourses)
        self.pushMyCourses = QtWidgets.QPushButton(self.frame_4)
        sizePolicy = QtWidgets.QSizePolicy(QtWidgets.QSizePolicy.Minimum, QtWidgets.QSizePolicy.Minimum)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.pushMyCourses.sizePolicy().hasHeightForWidth())
        self.pushMyCourses.setSizePolicy(sizePolicy)
        self.pushMyCourses.setMaximumSize(QtCore.QSize(227, 52))
        self.pushMyCourses.setObjectName("pushMyCourses")
        self.horizontalLayout.addWidget(self.pushMyCourses)
        self.pushOpenAdmin = QtWidgets.QPushButton(self.frame_4)
        sizePolicy = QtWidgets.QSizePolicy(QtWidgets.QSizePolicy.Minimum, QtWidgets.QSizePolicy.Minimum)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.pushOpenAdmin.sizePolicy().hasHeightForWidth())
        self.pushOpenAdmin.setSizePolicy(sizePolicy)
        self.pushOpenAdmin.setMaximumSize(QtCore.QSize(228, 52))
        self.pushOpenAdmin.setObjectName("pushOpenAdmin")
        self.horizontalLayout.addWidget(self.pushOpenAdmin)
        self.pushMyProfile = QtWidgets.QPushButton(self.frame_4)
        sizePolicy = QtWidgets.QSizePolicy(QtWidgets.QSizePolicy.Minimum, QtWidgets.QSizePolicy.Minimum)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.pushMyProfile.sizePolicy().hasHeightForWidth())
        self.pushMyProfile.setSizePolicy(sizePolicy)
        self.pushMyProfile.setMaximumSize(QtCore.QSize(227, 52))
        self.pushMyProfile.setObjectName("pushMyProfile")
        self.horizontalLayout.addWidget(self.pushMyProfile)
        self.gridLayout.addWidget(self.frame_4, 1, 0, 1, 4)
        self.frameAudit = QtWidgets.QFrame(self.centralwidget)
        self.frameAudit.setEnabled(True)
        sizePolicy = QtWidgets.QSizePolicy(QtWidgets.QSizePolicy.Expanding, QtWidgets.QSizePolicy.Preferred)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.frameAudit.sizePolicy().hasHeightForWidth())
        self.frameAudit.setSizePolicy(sizePolicy)
        self.frameAudit.setAutoFillBackground(False)
        self.frameAudit.setFrameShape(QtWidgets.QFrame.StyledPanel)
        self.frameAudit.setFrameShadow(QtWidgets.QFrame.Raised)
        self.frameAudit.setObjectName("frameAudit")
        self.verticalLayout = QtWidgets.QVBoxLayout(self.frameAudit)
        self.verticalLayout.setObjectName("verticalLayout")
        self.label_2 = QtWidgets.QLabel(self.frameAudit)
        self.label_2.setObjectName("label_2")
        self.verticalLayout.addWidget(self.label_2, 0, QtCore.Qt.AlignHCenter)
        self.plainTextEdit = QtWidgets.QPlainTextEdit(self.frameAudit)
        self.plainTextEdit.setObjectName("plainTextEdit")
        self.verticalLayout.addWidget(self.plainTextEdit)
        self.exportJSONBtn = QtWidgets.QPushButton(self.frameAudit)
        self.exportJSONBtn.setObjectName("exportJSONBtn")
        self.verticalLayout.addWidget(self.exportJSONBtn)
        self.gridLayout.addWidget(self.frameAudit, 0, 3, 1, 1)
        self.frame = QtWidgets.QFrame(self.centralwidget)
        sizePolicy = QtWidgets.QSizePolicy(QtWidgets.QSizePolicy.MinimumExpanding, QtWidgets.QSizePolicy.Expanding)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(self.frame.sizePolicy().hasHeightForWidth())
        self.frame.setSizePolicy(sizePolicy)
        self.frame.setMinimumSize(QtCore.QSize(659, 427))
        self.frame.setFrameShape(QtWidgets.QFrame.StyledPanel)
        self.frame.setFrameShadow(QtWidgets.QFrame.Raised)
        self.frame.setObjectName("frame")
        self.gridLayout_2 = QtWidgets.QGridLayout(self.frame)
        self.gridLayout_2.setObjectName("gridLayout_2")
        self.columnView = QtWidgets.QColumnView(self.frame)
        self.columnView.setObjectName("columnView")
        self.gridLayout_2.addWidget(self.columnView, 0, 0, 1, 1)
        self.gridLayout.addWidget(self.frame, 0, 0, 1, 3)
        MainWindow.setCentralWidget(self.centralwidget)
        self.menubar = QtWidgets.QMenuBar(MainWindow)
        self.menubar.setGeometry(QtCore.QRect(0, 0, 1091, 24))
        self.menubar.setObjectName("menubar")
        self.menupowered_by_Nickolsky = QtWidgets.QMenu(self.menubar)
        self.menupowered_by_Nickolsky.setObjectName("menupowered_by_Nickolsky")
        MainWindow.setMenuBar(self.menubar)
        self.statusbar = QtWidgets.QStatusBar(MainWindow)
        self.statusbar.setObjectName("statusbar")
        MainWindow.setStatusBar(self.statusbar)
        self.actionby_Nickolsky = QtWidgets.QAction(MainWindow)
        self.actionby_Nickolsky.setObjectName("actionby_Nickolsky")
        self.actionExit = QtWidgets.QAction(MainWindow)
        self.actionExit.setObjectName("actionExit")
        self.actionAudit = QtWidgets.QAction(MainWindow)
        self.actionAudit.setObjectName("actionAudit")
        self.action_2 = QtWidgets.QAction(MainWindow)
        self.action_2.setObjectName("action_2")
        self.menupowered_by_Nickolsky.addAction(self.actionExit)
        self.menupowered_by_Nickolsky.addAction(self.action_2)
        self.menupowered_by_Nickolsky.addAction(self.actionby_Nickolsky)
        self.menubar.addAction(self.menupowered_by_Nickolsky.menuAction())

        self.retranslateUi(MainWindow)
        QtCore.QMetaObject.connectSlotsByName(MainWindow)

    def retranslateUi(self, MainWindow):
        _translate = QtCore.QCoreApplication.translate
        MainWindow.setWindowTitle(_translate("MainWindow", "Golearn-account"))
        self.pushRefreshCourses.setText(_translate("MainWindow", "Обновить курсы"))
        self.pushMyCourses.setText(_translate("MainWindow", "Мои курсы"))
        self.pushOpenAdmin.setText(_translate("MainWindow", "Управление"))
        self.pushMyProfile.setText(_translate("MainWindow", "Профиль"))
        self.label_2.setText(_translate("MainWindow", "Аудит"))
        self.exportJSONBtn.setText(_translate("MainWindow", "Экспорт в json"))
        self.menupowered_by_Nickolsky.setTitle(_translate("MainWindow", "Опции"))
        self.actionby_Nickolsky.setText(_translate("MainWindow", "by Nickolsky"))
        self.actionExit.setText(_translate("MainWindow", "Выйти"))
        self.actionAudit.setText(_translate("MainWindow", "Аудит"))
        self.action_2.setText(_translate("MainWindow", "Настройки"))
