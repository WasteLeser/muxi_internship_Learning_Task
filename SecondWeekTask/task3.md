    package main

    import (
    	"sort"
    )

    type Student struct {
    	Name   string
    	Scores []int
    }

    // Average 返回平均分；Scores 为空时返回 0
        func (s *Student) Average() float64 {
      	if len(s.Scores) == 0 {
      	return 0
	    }
    	sum := 0
    	for _, v := range s.Scores {
    		sum += v
    	}
    	return float64(sum) / float64(len(s.Scores))
    }
  
    // AddScore 追加一个成绩，调用方能看见改动
    func (s *Student) AddScore(score int) {
    	s.Scores = append(s.Scores, score)
	    fmt.Printf("Added score %d for student %s. New scores: %v\n", score, s.Name, s.Scores)//看见改动
    }

    // Top 返回最高的 n 个成绩，从高到低
    func  (s *Student) Top(n int) []int {
	    if n <= 0 {
		    return nil
	    }
	    if len(s.Scores) == 0 {
		    return nil
	    }
    
    
	    tmp := make([]int, len(s.Scores))
	    copy(tmp, s.Scores)
	    sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
	    if n > len(tmp) {
		    n = len(tmp)
	    }
	    result := make([]int, n)
	    copy(result, tmp[:n])
	    return result
    }

    // Merge 合并两个人的成绩，返回新的 Student（Name 用接收者的）
    func (s *Student) Merge(other Student) Student {
	    merged := make([]int, 0, len(s.Scores)+len(other.Scores))
	    merged = append(merged, s.Scores...)
	    merged = append(merged, other.Scores...)
    
	    return Student{
		    Name:   s.Name,
		    Scores: merged,
	    }
    }
    
    func main() {
	    // 创建两个学生
	    student1 := Student{Name: "Alice", Scores: []int{85, 90, 78}}
	    student2 := Student{Name: "Bob", Scores: []int{88, 92, 80}}
    
	    // 计算平均分
	    fmt.Printf("%s's average score: %.2f\n", student1.Name, student1.Average())
	    fmt.Printf("%s's average score: %.2f\n", student2.Name, student2.Average())
  
	    // 添加成绩
	    student1.AddScore(95)
	    student2.AddScore(89)
  
	    // 获取最高的 n 个成绩
	    topScores1 := student1.Top(3)
	    topScores2 := student2.Top(2)
	    fmt.Printf("%s's top 3 scores: %v\n", student1.Name, topScores1)
	    fmt.Printf("%s's top 2 scores: %v\n", student2.Name, topScores2)
    
	    //  合并成绩
	    mergedStudent := student1.Merge(student2)
	    fmt.Printf("Merged student: %s, Scores: %v\n", mergedStudent.Name, mergedStudent.Scores)
    }	


有部分代码是tab大蛇造就的（
