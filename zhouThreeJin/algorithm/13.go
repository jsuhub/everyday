/*线段树：在一个数组的范围内要求对每个子数组进行更新之类的操作。
就类似于0———————3  有 0——0  0——1  0——2  0——3  1——1  1——2  1——3  2——2  2——3  3——3 的所有子数组都表示出来了  
那么维护所有的子数组大概需要O(n^2)的时间复杂度 、查询（任意一个子空间的值）为O（1）——————>此时维护的时间偏多
同样0——————3  有0——0 1——1 2——2 3——3 
那么假设我们维护的时间设为O(n) 查询(因为可能要查询0——2==》那么就需要查询0——0 + 1——1 +2——2 三段数组)为O(qn)————>此时查询时间变多
所以需要更加平衡的方案：  
同样0——3  有0——0 1——1 2——2 3——3    0——1  1——3    0——3 
类似于一颗二叉树的形式这样的话维护是假如 n=2^k==> n+n/2+n/4......+2+1= 2(1-2^k)/-1 +1 =2n-1
所以维护时间复杂度为O（n） 查询的时间复杂度为O(qlogn) 
 
根据这个原理我们就可以知道假设区间为0————n 我们需要开1+2+4+...n个节点大小==》4*n个节点大小虽然说最底层的树可能没有放满。
这样可能有点浪费空间了所以我们可以使用把n如10 看为二进制 1010   然后给他补全 1111 ==16 *2 =32 个大小就可以了。

当处于叶子节点的时候为l==r  而且节点为o的子节点为2*o 和 2*o+1  
所以查询的时候可以视为拆分为几个区间

类似图 ![image](./img/fig1.jepg)
*/

package algorithm
 
type lazySeg [] struct{
 sum , todo int 	
}

func ( t lazySeg )  queryAndAdd(o,l,r,L,R int ) int{   //l,r表示最大的区间 L，R
   res :=0
	if L<=l && r<=R{     //表示查询的区间完全在当前节点区间————这是给递归用的
		res  += t[o].sum
		t.do(o,l,r,t[o].todo)   //区间内的每个元素+1 是因为新添加的元素已经出现过了 并且子数组里面的值必须相邻
		return res
	}
	mid :=(l+r)/2
	if add :=t[o].todo;add!=0{
		t.do(o*2,l,mid,add)
		t.do(o*2+1,mid+1,r,add)
		t[o].todo =0
	}
	if L<=mid{
		res +=t.queryAndAdd(o*2,l,mid,L,R)
	}
	if R>mid{
		res +=t.queryAndAdd(o*2+1,mid+1,r,L,R)
	}
	return res
}

func  ( t lazySeg ) do( o int,l,r,add  int){   // t 表示数组 o当前节点对应的下表，l,r表示当前节点对应的区间，add表示增量	
	t[o].sum +=(r-l+1)*add   //表示更新当前区间的和
	t[o].todo +=add  //表示当前区间的增量 
}

func sumCount(nums []int)(ans int ){
	last :=make(map[int]int)   //表示当前元素出现的最后位置
	n:=len(nums)
	tree :=make(lazySeg,4*n)  //表示线段树
	s:= 0   //表示当前的和
	for i,v:=range nums{
		i++
		j:=last[v]
		s+= tree.queryAndAdd(1,1,n,j+1,i)*2+i-j  //从1开始到n 
		ans =(ans+s)%1000000000+7
		last[v]=i
	}
	return ans
}
