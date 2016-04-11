# RETURN LARGEST NUMBERS IN ARRAYS
Return an array consisting of the largest number from each provided sub-array.<br>For simplicity, the provided array will contain exactly 4 sub-arrays.  

```go
array1 := []int{13, 27, 18, 26}
array2 := []int{4, 5, 1, 3}
array3 := []int{32, 35, 37, 39}
array4 := []int{1000, 1001, 857, 1}

largestOfFour(array1, array2, array3, array4) // => [27,5,39,1001]

array5 := []int{4, 9, 1, 3}
array6 := []int{13, 35, 18, 26}
array7 := []int{32, 35, 97, 39}
array8 := []int{1000000, 1001, 857, 1}
largestOfFour(array5, array6, array7, array8) // => [9, 35, 97, 1000000]
```
