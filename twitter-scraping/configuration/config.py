from dotenv import load_dotenv
from selenium import webdriver
import os

class Config:
    def __init__(self) -> None:
        load_dotenv()
        
    def getDriver(self) -> webdriver:
        "return a webdriver by the OS variable setted in .env file"
        system = os.getenv("operateSystem")
        prefix = os.getenv("PY_PATH")
        
        match system:
            case "windows":
                path = prefix + "/chromedriver.exe"
                return webdriver.Chrome(path)
            case "mac_x86":
                path = prefix + "/chromedriver_x86"
                return webdriver.Chrome(path)
            case "mac_ARM":
                path = prefix + "/chromedriver_ARM"
                return webdriver.Chrome(path)
            
    def getAccountMail(self) -> str:
        "return a webdriver by the mail variable setted in .env file"
        mail = os.getenv("mail")
        return mail
    
    def getAccountName(self) -> str:
        "return a webdriver by the name variable setted in .env file"
        name = os.getenv("name")
        return name
    
    def getPassword(self) -> str:
        "return a webdriver by the password variable setted in .env file"
        password = os.getenv("password")
        return password