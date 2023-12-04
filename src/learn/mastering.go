
// package learn
// //deleting from a slice at index i soloution 1 :

// if i > len(aSlice) -1 {
// 	fmt.Println("Cannot delete element", i)
// return
// }
// aSlice = append(aSlice[:i], aSlice[i+1:]...)

// //deleting from a slice at index i soloution 2 :
// if i > len(aSlice) -1 {
// 	fmt.Println("Cannot delete element", i)
// 	// Replace element at index i with last element
// 	aSlice[i] = aSlice[len(aSlice)-1]
// 	// Remove last element
// 	aSlice = aSlice[:len(aSlice)-1]
// ////////////////////////////////////////////////////////////////////////////