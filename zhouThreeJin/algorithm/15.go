// import "fmt"
// func satisfiesConditions(grid [][]int) bool {
//     number := true
//     for key ,value := range grid{
//        if key!=len(grid)-1 {
//      for i,v :=range value{
//             if i!= len(value)-1{
//                 if v!=grid[key+1][i]|| v==grid[key][i+1] {
//                     number=false 
//                     break 
//                 }
//                     }else{
//                         if v!=grid[key+1][i]{
//                             number =false
//                         }
//                            }
//                            if(number==false){
//                             break
//                            }
//        }
//        }else {  
//         for i,v :=range value{
//             if i!= len(value)-1{
//                 if  v==grid[key][i+1] {
//                     number=false 
//                     break 
//                 }
//                     }
//                            }
//                            if(number==false){
//                             break
//                            }
//        }  
//     }
//     return number
// }