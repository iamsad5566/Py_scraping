import csv, os
from datetime import datetime
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
        now = datetime.now()
        currDate = now.strftime("%m_%d_%y")
        output = os.getenv("fileOutput")
        suffix = ".csv"
        
        for i in range(run):
            articles = driver.find_element(By.XPATH, "//*[@id=\"react-root\"]/div/div/div[2]/main/div/div/div/div[1]/div/div[3]/div/section/div")
            self.__processArticles(articles.text, output+currDate+suffix)
            driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
            time.sleep(2)
    
    def __processArticles(self, articles, outputFile):
        for article in articles.split("\n"):
            if len(article) < 10 or article.startswith("@"):
                continue
            
            with open(outputFile, "a", newline="") as csvfile:
                writer = csv.writer(csvfile)
                writer.writerow([article])
                
            
    def closeDriver(self):
        self.driver.close()