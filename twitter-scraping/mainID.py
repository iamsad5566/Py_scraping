import sys
from module.processController import ProcessController


class TwitterScrapingByID:
    def __init__(self) -> None:
        self.procedure = ProcessController()
        
    def singleID(self, ID, run):
        self.procedure.singleID(ID, run)
         
    def multipleIDs(self, IDs, run):
        self.procedure.multipleIDs(IDs, run)

if __name__ == "__main__":
    TwitterScrapingByID().singleID("matsu_bouzu", 2)