from lib2to3.pgen2 import driver
from dotenv import load_dotenv
from selenium import webdriver
import os

class Config:
    def __init__(self) -> None:
        load_dotenv()
    def getDriver(self):
        system = os.getenv("OS")
        prefix = os.getcwd()
        match system:
            case "windows":
                path = prefix + "/chromedrive.exe"
                return webdriver.Chrome(path)
            case "mac_x86":
                path = prefix + "/chromedriver_x86"
                return webdriver.Chrome(path)
            case "mac_ARM":
                path = prefix + "/chromedriver_ARM"
                return webdriver.Chrome(path)