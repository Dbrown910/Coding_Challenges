def capitalize(string):
    fullName = string.split(" ")
    capName = ""
    for i in range(len(fullName)):
        capName += fullName[i].capitalize()
        if (i < len(fullName) - 1):
            capName += " "
    print(capName)
        
    return capName

capitalize("hello world lol")