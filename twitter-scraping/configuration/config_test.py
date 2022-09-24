import os
from dotenv import load_dotenv
import unittest
from config import Config

class TestConfig(unittest.TestCase):
    def __init__(self, methodName: str = ...) -> None:
        super().__init__(methodName)
        load_dotenv()
        self.config = Config()
        
    def testGetDriver(self):
        driver = self.config.getDriver()
        self.assertEqual(type(driver).__name__, "WebDriver")
        driver.close()
        
    def testGetAccountMail(self):
        testAccountMail = os.getenv("mail")
        self.assertEqual(testAccountMail, self.config.getAccountMail())
    
    def testGetAccountName(self):
        testName = os.getenv("name")
        self.assertEqual(testName, self.config.getAccountName())
        
    def testGetPassword(self):
        testPassword = os.getenv("password")
        self.assertEqual(testPassword, self.config.getPassword())
        
if __name__ == "__main__":
    unittest.main()