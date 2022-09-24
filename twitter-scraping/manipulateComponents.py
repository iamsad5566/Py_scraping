from configuration.config import Config
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys

import time

class ManipulateComponents:
    def __init__(self) -> None:
        self.config = Config()
        self.driver = self.config.getDriver()
        
    def login(self) -> None:
        driver = self.driver
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
        loginButton = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/div/div/span/span")
        loginButton.click()
        
    def search(self):
        time.sleep(6)
        driver = self.driver
        searchInput = driver.find_element(By.XPATH, "//*[@id=\"react-root\"]/div/div/div[2]/main/div/div/div/div[2]/div/div[2]/div/div/div/div[1]/div/div/div/form/div[1]/div/div/div/label/div[2]/div/input")
        searchInput.send_keys("Taiwan")
        searchInput.send_keys(Keys.RETURN)
        time.sleep(10)
        