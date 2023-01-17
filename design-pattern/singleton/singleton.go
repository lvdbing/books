package main

// 单例模式大致分为两类：饿汉模式和懒汉模式
//     饿汉模式 - 在系统初始化时完成单例对象的实例化
//     懒汉模式 - 调用时才进行延迟实例化

import "sync"

// ******* 饿汉模式 *******
type network struct{}

var instance = &network{}

func Instance() *network {
	return instance
}

// ******* 懒汉模式 *******

var mutex sync.Mutex // 互斥锁，防止并发问题。

// 1. 普通加锁，缺点是每次调用时都需要加锁。（不推荐）
func LazyInstance() *network {
	mutex.Lock()
	if instance == nil {
		instance = &network{}
	}
	mutex.Unlock()
	return instance
}

// 2. 双重检验后加锁，实例化后无需加锁。
func LazyInstance2() *network {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = &network{}
		}
		mutex.Unlock()
	}
	return instance
}

// 3. Go语言更优雅的实现方式，利用sync.Once。（强烈推荐，更优雅）
var once sync.Once

func LazyInstance3() *network {
	once.Do(func() {
		instance = &network{}
	})
	return instance
}
