from model.manipulateComponents import ManipulateComponents

class ProcessController:
    def __init__(self) -> None:
        self.components = ManipulateComponents()
    
    def __commonProcess(self, keyword, run):
        "For common logics"
        # Login into twitter account
        self.components.login()
        # Start the first searching
        self.components.firstSearch(keyword)
        # Start scraping and saving the data into a csv file
        self.components.scrapeArticles(run)
    
    def singleSearching(self, run=5, keyword=""):
        "For sigle searching"
        self.__commonProcess(keyword, run)
        # Closing the driver
        self.components.closeDriver()
        
    def multipleSearching(self, run=5, keywords=["a", "b", "c"]):
        "For multiple searching"
        self.__commonProcess(keywords[0], run)
        for i in range(1, len(keywords), 1):
            self.components.newSearch(keywords[i])
            self.components.scrapeArticles(run)
        # Closing the driver
        self.components.closeDriver()