def find_commons(str):
    letterDict = {}
    
    # get a count for every letter
    for i in range(len(str)):
        if str[i] in letterDict:
            letterDict[str[i]] += 1
        else:
            letterDict[str[i]] = 1

    # get all the values sorted in desc order
    sortedVals = sorted(letterDict.values(), reverse=True)
    # get a list of all keys sorted
    sortedKeys = sorted(letterDict.keys())
    topThreeList = {}
    

    # for every key, see if its value is in the top 3 values of the desc order values
    # if it is store it and the value
    for k in sortedKeys:
        for i in range(0,3):
            if letterDict[k] == sortedVals[i]:
                topThreeList[k] = sortedVals[i]
                break      
    
    # reverse the keys and values and sort by the values
    d_view = [ (v,k) for k,v in topThreeList.items() ]
    print(d_view)
    d_view.sort(reverse=True)
    print(d_view)

    # print each pair in the topThreeList
    # TODO: out how to sort sublists of pairs by letter whose count is the same
    for v,k in d_view:
        print("%s  %d" % (k,v))


find_commons("aabbbccde")