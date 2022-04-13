

## set and dictionary

There is no set structure in golang, will use dictionary instead.

    //cpp
    set<int> nums1_set
    
    //golang
    nums1_set := make(map[int]int); // golang
    for _, val := range nums1 {
    	nums1_set[val] = 1
    }
    	
To traverse a set/dictionary in golang 
  
    for key, _ := range m1 {
        //do something here
    }  	
    
To check if key exists in a dict in golang:   
 
    if _, ok := m2[key]; ok {
    	// do something 
    }    

## slice in golang (vector in cpp)

	var answer [][]int
	var answer0 []int
	
	answer0 = append(answer0, key)
	answer = append(answer, answer0)

## 	
	var intArray [5]int
	var strArray [5]string