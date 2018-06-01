def minion_game(string):
    # your code goes here
    vowels = 'AEIOU'
    scoreA = 0
    scoreB = 0
    letterCount = len(string)

    for i in range(letterCount):
        if string[i] in vowels:
            scoreB += (letterCount-i)
        else:
            scoreA += (letterCount-i)
            
    if scoreA > scoreB:
        print('Stuart '+ str(scoreA))
    elif scoreB > scoreA:
        print('Kevin '+ str(scoreB))
    elif scoreA == scoreB:
        print('Draw')

minion_game('NANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANNANANN')