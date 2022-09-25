import unittest
from manipulateComponents import ManipulateComponents
import time

class TestManipulateComponents(unittest.TestCase):
    def testLogin(self):
        ManipulateComponents().login()
        time.sleep(1)
    
    def testSearch(self):
        manipulateComponents = ManipulateComponents()
        manipulateComponents.login()
        manipulateComponents.firstSearch("Taiwan")
        time.sleep(1)
    
    def testScrpeArticles(self):
        manipulateComponents = ManipulateComponents()
        manipulateComponents.login()
        manipulateComponents.firstSearch("Taiwan")
        manipulateComponents.scrapeArticles(1)
        manipulateComponents.newSearch("Tennis")
        time.sleep(100)
        

if __name__ == "__main__":
    TestManipulateComponents().testScrpeArticles()