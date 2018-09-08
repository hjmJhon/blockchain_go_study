package hashmap

type Entry struct {
	//key
	K string
	//value
	V interface{}
	//下一个Entry的指针
	next *Entry
}

type HashMap struct {
	//HashMap 的大小
	size int
	//存放的 Entry
	bucket []Entry
}

/*
	计算 hash 值
 */
func hashCode(k string, length int) int {
	sum := 0
	bytes := []byte(k)
	for _, value := range bytes {
		sum += int(value)
	}
	return sum % length
}

var nowCapacity = 10

const maxCapacity = 100
const loadFactor = 0.75

/*
	创建 HashMap
 */
func CreateHashMap() *HashMap {
	return &HashMap{
		size:   0,
		bucket: make([]Entry, nowCapacity, maxCapacity),
	}
}

/*
	往 HashMap 添加元素
 */
func (hm *HashMap) Put(k string, v interface{}) {
	entry := Entry{
		K:    k,
		V:    v,
		next: nil,
	}
	hm.insert(entry)

	//扩容

	if float64(hm.size)/float64(len(hm.bucket)) > loadFactor {
		if nowCapacity*2 > maxCapacity {
			nowCapacity = maxCapacity
		} else {
			nowCapacity = nowCapacity * 2
		}
		//创建新的 HashMap
		newHm := &HashMap{
			size:   0,
			bucket: make([]Entry, nowCapacity, maxCapacity),
		}
		//将原来 HashMap 的数据复制到新的 HashMap
		for _, v := range hm.bucket {
			if v.K == "" {
				continue
			}

			for v.next != nil {
				newHm.insert(v)
				v = *(v.next)
			}

			newHm.insert(v)
		}

		*hm = *newHm
	}
}

/*
	根据 key 获取 value
 */
func (hm *HashMap) Get(k string) interface{} {
	return hm.getEntry(k).V
}

func (hm *HashMap) Size() int {
	return hm.size
}

func (hm *HashMap) Cap() int {
	return len(hm.bucket)
}

/*
	插入到 HashMap
 */
func (hm *HashMap) insert(entry Entry) {
	//1.计算出key在数组中的index
	//2.该index的Entry的key为空, 直接插入
	//3.该index的Entry的key不为空,遍历所有的Entry

	index := hashCode(entry.K, nowCapacity)
	e := &hm.bucket[index]
	if e.K == "" {
		*e = entry
		hm.size++
	} else {
		for e.next != nil {
			if e.K == entry.K {
				*e = entry
				break
			}
			e = e.next
		}

		if e.K == entry.K {
			*e = entry
		} else {
			e.next = &entry
			hm.size++
		}
	}
}

/*
	根据 key 获取 entry
 */
func (hm *HashMap) getEntry(k string) Entry {
	index := hashCode(k, nowCapacity)
	entry := hm.bucket[index]
	if entry.K == "" {
		return Entry{}
	} else {
		for entry.next != nil {
			if entry.K == k {
				return entry
			}
			entry = *(entry.next)
		}

		if entry.K == k {
			return entry
		} else {
			return Entry{}
		}
	}
}
