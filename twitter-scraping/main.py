from manipulateComponents import ManipulateComponents

class TwitterPyScraping:
    def __init__(self) -> None:
        self.manipulate = ManipulateComponents()
        self.manipulate.login()
        self.manipulate.search("Taiwan")
        self.manipulate.scrapeArticles(5)
        self.manipulate.closeDriver()
        
        
        
if __name__ == "__main__":
    TwitterPyScraping()
