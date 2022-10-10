import sys
from module.processController import ProcessController

class TwitterScrapingByID:
    def __init__(self) -> None:
        self.procedure = ProcessController()
        
    def singleID(self, run, ID):
        "singleID scrapes the tweets from single ID"
        self.procedure.singleID(ID, run)
         
    def multipleIDs(self, run, IDs):
        "multipleIDs scrapes the tweets from multiple IDs"
        self.procedure.multipleIDs(IDs, run)

if __name__ == "__main__":
    # args[0] is the path of venv
    # args[1] is the file path and file name
    # args[2] is the execute mode
    # args[3] is the run times
    def IDBuild(args):
        output = []
        for i in range(3, len(args)):
            output.append(args[i]) 
        return output
    
    args = sys.argv
    if len(args) == 4:
        TwitterScrapingByID().singleID(int(args[2]), args[3])
    else:
        IDs = IDBuild(args)
        TwitterScrapingByID().multipleIDs(int(args[2]), IDs)