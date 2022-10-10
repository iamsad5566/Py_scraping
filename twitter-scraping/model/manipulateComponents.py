import csv, os
from datetime import datetime
from configuration.config import Config
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys

import time

class ManipulateComponents:
    def __init__(self) -> None:
        self.prefix = "https://twitter.com/"
        self.config = Config()
        self.driver = self.config.getDriver()
    
    # Login twitter    
    def login(self):
        driver = self.driver
        driver.get("https://twitter.com/i/flow/login")
        time.sleep(3)
        # Key in account name
        inputEle = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[5]/label/div/div[2]/div/input")
        inputEle.send_keys(self.config.getAccountMail())
        nextButton = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[6]")
        nextButton.click()
        
        # Key in user name
        time.sleep(2)
        try:
            inputEle = driver.find_element(By.TAG_NAME, "input")
            inputEle.send_keys(self.config.getAccountName())
            nextButton = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div/div/div/div")
            nextButton.click()
        except:
            print("Keep moving on!")
        
        # Key in password
        time.sleep(2)
        inputEle = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div/div[3]/div/label/div/div[2]/div[1]/input")
        inputEle.send_keys(self.config.getPassword())
        loginButton = driver.find_element(By.XPATH, "//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/div/div/span/span")
        loginButton.click()
        
    def firstSearch(self, keyword):
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
            if len(article) < 15 or article.startswith("@"):
                continue
            
            with open(outputFile, "a", newline="", encoding="UTF-8") as csvfile:
                writer = csv.writer(csvfile)
                writer.writerow([article])
    
    def newSearch(self, keyword):
        time.sleep(3)
        driver = self.driver
        searchInput = driver.find_element(By.XPATH, "//*[@id=\"react-root\"]/div/div/div[2]/main/div/div/div/div[1]/div/div[1]/div[1]/div[1]/div/div/div/div/div[2]/div[2]/div/div/div/form/div[1]/div/div/div/label/div[2]/div/input")
        lastKeyword = searchInput.get_attribute("value")
        for i in range(len(lastKeyword)):
            searchInput.send_keys(Keys.BACKSPACE)
        time.sleep(1)
        searchInput.send_keys(keyword)
        searchInput.send_keys(Keys.RETURN)
        
    def scrapeByID(self, ID, run):
        "scraeByID visits the home page of the ID owner and scraping the tweets one by one"
        driver = self.driver
        driver.get(self.prefix + ID)
        time.sleep(3)
        self.__parsingTweets(ID, run)
            
    def __parsingTweets(self, ID, run):
        "__parsingTweets is serving for scrapeByID, once the page was navigated to the home page of the ID owner, __startScrapingTweets will be called for the following scraping works"
        driver = self.driver
        now = datetime.now()
        currDate = now.strftime("%m_%d_%y")
        output = os.getenv("fileOutput")
        suffix = ".csv"
        
        for i in range(run):
            timeLine = driver.find_element(By.XPATH, "//*[@id=\"react-root\"]/div/div/div[2]/main/div/div/div/div[1]/div/div[3]/div/div/section/div")
            tweetsArr = timeLine.text.split("@"+ID)
            self.__exportParsedTweets(tweetsArr, output+currDate+suffix)
            
            driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
            time.sleep(2)
    
    def __exportParsedTweets(self, tweetsArr, outputfile):
        for tweet in tweetsArr:
            with open(outputfile, "a", newline="", encoding="UTF-8") as csvfile:
                writer = csv.writer(csvfile)
                writer.writerow([tweet])
    
    def closeDriver(self):
        self.driver.close()