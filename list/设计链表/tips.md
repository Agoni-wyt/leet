### 设计链表

## 获取链表: get(index)
    for i<index{
        node = node.Next
        i++
    }
    return node

## 添加链表头：addAtHead(val)
    newNode.Next = headNode


## 添加链表尾：addAtTail(val)
    tailNode.Next=newNode

## 插入节点在指定位置: addAtIndex(index,val)
    newNode.Next:=preNode.Next
    preNode.Next= newNode

## 删除节点在指定位置：deleteAtIndex(index)
    preNode.Next = toDelNode.Next


## 链表的两种操作方式
1. 在原链表操作
2. 设置一个假头，并让假头指向头，然后再操作