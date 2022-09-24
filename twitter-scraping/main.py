from manipulateComponents import ManipulateComponents

class TwitterPyScraping:
    def __init__(self) -> None:
        self.manipulate = ManipulateComponents()
        self.manipulate.login()
        self.manipulate.search()
    
        
        
if __name__ == "__main__":
    TwitterPyScraping()
