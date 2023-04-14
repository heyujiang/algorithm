package array

//870. 优势洗牌
import (
	"sort"
)

type valIdx [][2]int

func (v valIdx) Len() int {
	return len(v)
}

func (v valIdx) Swap(i,j int) {
	v[i],v[j] = v[j],v[i]
}

func (v valIdx) Less(i, j int) bool {
	return v[i][0] < v[j][0]
}

func advantageCount(nums1 []int, nums2 []int) []int {
	//构造带有原下标信息的新的nums2的数组
	valIdxs := make(valIdx,0,len(nums2))
	for k, v := range nums2 {
		valIdxs = append(valIdxs,[2]int{v,k})
	}

	sort.Sort(valIdxs)
	sort.Ints(nums1)

	res := make([]int,len(nums1))
	lessIndex := -1
	for i,j := len(nums1) - 1,len(valIdxs) - 1; i > lessIndex; {
		if nums1[i] <= valIdxs[j][0] {
			res[valIdxs[j][1]] = nums1[lessIndex+1]
			lessIndex++
			j--
		}	else {
			res[valIdxs[j][1]] = nums1[i]
			i--
			j--
		}
	}

	return res
}


/**
…or create a new repository on the command line

	echo "# Repository01" >> README.md
	git init:q
	git add README.md
	git commit -m "first commit"
	git branch -M master
	git remote add origin git@github.com:heyujiang/Repository01.git
	git push -u origin master

…or push an existing repository from the command line
	git remote add origin git@github.com:heyujiang/Repository01.git
	git branch -M master
	git push -u origin master

 */