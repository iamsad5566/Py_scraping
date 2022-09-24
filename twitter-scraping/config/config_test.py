from distutils.command.config import config
from os import unlink
import unittest
from config import Config

class TestConfig(unittest.TestCase):
    def testGetDriver(self):
        driver = Config().getDriver()
        self.assertEqual(type(driver).__name__, "WebDriver")
        driver.close()
        
if __name__ == "__main__":
    unittest.main()