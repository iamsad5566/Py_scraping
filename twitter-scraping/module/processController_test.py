from processController import ProcessController
import numpy as np
import unittest, time

class TestProcessController(unittest.TestCase):
    def __init__(self, methodName: str = ...) -> None:
        super().__init__(methodName)
        self.testController = ProcessController()
        
    def testSingleSearching(self):
        self.testController.singleSearching(5, "taiwan")
        time.sleep(10)
        
    def testMultipleSearching(self):
        testSearchingList = np.array(["taiwan", "japan", "us", "uk", "switzerland"])
        self.testController.multipleSearching(5, testSearchingList)
        time.sleep(10)

if __name__ == "__main__":
    def suite():
        suite = unittest.TestSuite()
        suite.addTest(TestProcessController("testMultipleSearching"))
        return suite
    
    runner = unittest.TextTestRunner()
    runner.run(suite())