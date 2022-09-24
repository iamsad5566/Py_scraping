from configuration.config import Config
from selenium.webdriver.common.by import By
import time

class ProcessController:
    def __init__(self) -> None:
        self.config = Config()
        
    def login(self) -> None:
        driver = self.config.getDriver()
        driver.get("https://twitter.com/i/flow/login")
        time.sleep(5)
        # Key in account name
        inputEle = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[5]/label/div/div[2]/div/input")
        inputEle.send_keys(self.config.getAccountMail())
        nextButton = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[6]")
        nextButton.click()
        
        # Key in user name
        time.sleep(3)
        inputEle = driver.find_element(By.TAG_NAME, "input")
        if inputEle != None:
            inputEle.send_keys(self.config.getAccountName())
            nextButton = driver.find_element(By.CSS_SELECTOR, "#layers > div > div > div > div > div > div > div.css-1dbjc4n.r-1awozwy.r-18u37iz.r-1pi2tsx.r-1777fci.r-1xcajam.r-ipm5af.r-g6jmlv > div.css-1dbjc4n.r-1867qdf.r-1wbh5a2.r-kwpbio.r-rsyp9y.r-1pjcn9w.r-1279nm1.r-htvplk.r-1udh08x > div > div > div.css-1dbjc4n.r-14lw9ot.r-6koalj.r-16y2uox.r-1wbh5a2 > div.css-1dbjc4n.r-16y2uox.r-1wbh5a2.r-1jgb5lz.r-1ye8kvj.r-13qz1uu > div.css-1dbjc4n.r-1isdzm1 > div > div > div > div > div")
            nextButton.click()
        
        # Key in password
        time.sleep(3)
        inputEle = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div/div[3]/div/label/div/div[2]/div[1]/input")
        inputEle.send_keys(self.config.getPassword())     
        loginButton = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/div/div")
        loginButton.click()
        time.sleep(10)