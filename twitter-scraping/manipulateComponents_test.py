import unittest
from manipulateComponents import ManipulateComponents
import time

class testManipulateComponents(unittest.TestCase):
    def testLogin(self):
        ManipulateComponents().login()
        time.sleep(1)
    
    def testSearch(self):
        manipulateComponents = ManipulateComponents()
        manipulateComponents.login()
        manipulateComponents.search()
        time.sleep(1)
    
    def testScrpeArticles(self):
        manipulateComponents = ManipulateComponents()
        manipulateComponents.login()
        manipulateComponents.search("Taiwan")
        manipulateComponents.scrapeArticles(3)
        time.sleep(1)
        

if __name__ == "__main__":
    unittest.main()