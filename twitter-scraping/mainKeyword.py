import sys
from module.processController import ProcessController

class TwitterPyScrapingByKeyword:
    def __init__(self) -> None:
        self.procedure = ProcessController()
    
    def runSingleSearching(self, run=1, keyword=""):
        "Run a single searching by selenium"
        self.procedure.singleSearching(run=run, keyword=keyword)
    
    def runMultipleSearching(self, run=1, keywords=[]):
        "Run a multiple searching by selenium"
        self.procedure.multipleSearching(run, keywords=keywords)
        
if __name__ == "__main__":
    # args[0] is the file path and file name
    # args[1] is the execute mode
    # args[2] is the run times
    # args[3] is the keyword(s)
    def keywordBuilder(args):
        output = []
        for i in range(3, len(args)):
            output.append(args[i])
        return output
    
    args = sys.argv
    if args[1] == "Single":
        TwitterPyScrapingByKeyword().runSingleSearching(int(args[2]), args[3])
    else:
        keywords = keywordBuilder(args)
        TwitterPyScrapingByKeyword().runMultipleSearching(int(args[2]), keywords)
        
        
        
    