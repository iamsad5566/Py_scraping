from lib2to3.pgen2 import driver
import unittest
from processController import ProcessController

class testProcessController(unittest.TestCase):
    def testLogin(self):
        ProcessController().login()
        

if __name__ == "__main__":
    unittest.main()