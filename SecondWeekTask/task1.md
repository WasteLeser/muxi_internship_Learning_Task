
主要区别：
切片扩容是开一个更大的数组，将旧数据复制后，回收原来切片的底层数组

桶扩容是渐进式的，在扩容过程中先完成一部分扩容，记录扩容进度，直到所有桶迁移完成，在进行map操作时迁移数据，每次至多迁移两个桶，一个是正在操作的桶，以保证使用的数据是就绪的，一个是扩容进度标记的桶，以保证每个桶都能顺利迁移

---

# 学习笔记：
我进行了一些搜索，大致学习了切片与桶的底层结构
以下为源代码
扩容的代码部分我没有搞清楚，遂在bilibili上进行了学习


        type slice struct {//切片的底层结构
         array unsafe.Pointer
         len   int
         cap   int
        }


        type hmap struct {//桶的底层结构（？
        	count     int    // 当前 map 中的元素个数，用于 len() 操作。
        	flags     uint8  
        	B         uint8  
        	noverflow uint16 // 用到的溢出桶的大概个数
        	hash0     uint32 

        	buckets    unsafe.Pointer // 指向当前桶数组地址的指针
        	oldbuckets unsafe.Pointer // 扩容过程中，保存扩容旧的桶数组的指针
        	nevacuate  uintptr        // 扩容过程中记录迁移进度

        	extra *mapextra 
        }


