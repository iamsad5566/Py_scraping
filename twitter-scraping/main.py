import sys
from module.processController import ProcessController

class TwitterPyScraping:
    def __init__(self) -> None:
        self.procedure = ProcessController()
    
    def runSingleSearching(self, keyword):
        "Run a single searching by selenium"
        self.procedure.singleSearching(run=5, keyword=keyword)
    
    def runMultipleSearching(self, keywords):
        "Run a multiple searching by selenium"
        self.procedure.multipleSearching(run=5, keywords=keywords)
        
if __name__ == "__main__":
    # args[0] would be file path and file name
    # args[1] is the execute mode
    def keywordBuilder(args):
        output = []
        for i in range(2, len(args)):
            output.append(args[i])
        return output
    
    args = sys.argv
    if len(args) == 2:
        TwitterPyScraping().runSingleSearching(args[2])
    else:
        keywords = keywordBuilder(args)
        TwitterPyScraping().runMultipleSearching(keywords)
        
        
        
    