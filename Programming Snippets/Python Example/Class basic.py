class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    # __str__ -> this method will help a class to formate a string default
    def __str__(self):
        return f"{self.name} & {self.age}"
    
    # Create a simple method for this object
    def Print(self):
        print("Name : ", self.name)
        print("Age : ", self.age)
        

p1 = Person("test name", 33)

print(p1)

# Calling the method
p1.Print()
