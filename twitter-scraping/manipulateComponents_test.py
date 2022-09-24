import unittest
from manipulateComponents import ManipulateComponents
import time

class testManipulateComponents(unittest.TestCase):
    def testLogin(self):
        ManipulateComponents().login()
        time.sleep(10)
    
    # def testSearch(self):
    #     ManipulateComponents().login()
    #     ManipulateComponents().search()
    #     time.sleep(10)
        

if __name__ == "__main__":
    unittest.main()