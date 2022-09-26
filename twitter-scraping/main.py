from ast import keyword
from module.processController import ProcessController

class TwitterPyScraping:
    def __init__(self) -> None:
        self.procedure = ProcessController()
    
    def runSingleSearching(self):
        "Run a single searching by selenium"
        keyword = "tw"
        self.procedure.singleSearching(run=5, keyword=keyword)
    
    def runMultipleSearching(self):
        "Run a multiple searching by selenium"
        keyword = ["tw", "uk", "jp", "us"]
        self.procedure.multipleSearching(run=5, keyword=keyword)
        
if __name__ == "__main__":
    TwitterPyScraping().runMultipleSearching()
