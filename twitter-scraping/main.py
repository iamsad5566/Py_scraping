from processController import ProcessController

class TwitterPyScraping:
    def __init__(self) -> None:
        self.process = ProcessController()
    
    def runSingleSearching(self):
        exampleSearch = "Taiwan"
        self.process.singleSearching(run=5, keyword=exampleSearch)
    
    def runMultipleSearching(self):
        exampleSearchList = ["tw", "jp", "us", "uk"]
        self.process.multipleSearching(run=5, keyword=exampleSearchList)
        
if __name__ == "__main__":
    TwitterPyScraping()
