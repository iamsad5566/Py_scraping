from model.manipulateComponents import ManipulateComponents

class ProcessController:
    def __init__(self) -> None:
        self.components = ManipulateComponents()
        
    # For common logics
    def __commonProcess(self, keyword, run):
        # Login into twitter account
        self.components.login()
        # Start the first searching
        self.components.firstSearch(keyword)
        # Start scraping and saving the data into a csv file
        self.components.scrapeArticles(run)
    
    # For sigle searching
    def singleSearching(self, run=5, keyword=""):
        self.__commonProcess(keyword, run)
        # Closing the driver
        self.components.closeDriver()
        
    # For multiple searching
    def multipleSearching(self, run=5, keyword=["a", "b", "c"]):
        self.__commonProcess(keyword[0], run)
        for i in range(1, len(keyword), 1):
            self.components.newSearch(keyword[i])
            self.components.scrapeArticles(run)
        # Closing the driver
        self.components.closeDriver()