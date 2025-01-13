// import "sort"
// func minMovesToSeat(seats []int, students []int) int {
//    seat1 := make ([]int,len(seats))
//    student1 := make([]int,len(students))
//    copy(seat1,seats)
//    copy(student1,students)
//    sort.Ints(seat1)
//    sort.Ints(student1)
//    sum :=0
//    for i,v :=range seat1{
//     if(v>=student1[i]){
//         sum=sum+v-student1[i]
//     }else{
//         sum=sum+student1[i]-v
//     }

//    }
//    return sum 
// }