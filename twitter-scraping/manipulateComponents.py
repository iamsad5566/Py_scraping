from lib2to3.pgen2 import driver
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
            nextButton = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div/div/div/div")
            nextButton.click()
        
        # Key in password
        time.sleep(3)
        inputEle = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div/div[3]/div/label/div/div[2]/div[1]/input")
        inputEle.send_keys(self.config.getPassword())
        loginButton = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/div/div/span/span")
        loginButton.click()
        
    def search(self, keyword):
        time.sleep(5)
        driver = self.driver
        searchInput = driver.find_element(By.XPATH, "//*[@id=\"react-root\"]/div/div/div[2]/main/div/div/div/div[2]/div/div[2]/div/div/div/div[1]/div/div/div/form/div[1]/div/div/div/label/div[2]/div/input")
        searchInput.send_keys(keyword)
        searchInput.send_keys(Keys.RETURN)
        
    def scrapeArticles(self, run):
        time.sleep(5)
        driver = self.driver
        for i in range(run):
            articles = driver.find_element(By.XPATH, "//*[@id=\"react-root\"]/div/div/div[2]/main/div/div/div/div[1]/div/div[3]/div/section/div")
            print(articles.text)
            driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
            time.sleep(2)