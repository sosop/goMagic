# goMagic golang开发的爬虫工具

**设计思路参考WebMagic，依赖goquery**  


```
go get -u github.com/PuerkitoBio/goquery
go get github.com/sosop/goMagic
```


### 一、使用
#### 1、实现processor.Processor接口

```
type ToutiaoProcessor struct {
}

func (tt *ToutiaoProcessor) Process(p *downloader.Page) {
	q, err := p.Parser()

	if err != nil {
		log.Println(err)
		return
	}
	q.Find(".post").Each(func(index int, s *goquery.Selection) {
		content := s.Find(".content .title a")
		title := content.Text()
		url, _ := content.Attr("href")
		url = strings.TrimSpace(url)
		p.PutField("Title", title)
		p.PutField("URL", url)
		if url != "" {
			p.AddTargetURL(url)
		}
	})
	// var articles []Article
	// p.Objects(&articles)
}
```

#### 2、启动


```
core.NewMagic("test", &ToutiaoProcessor{}).AddURL("http://toutiao.io/"). /*.SetThread(8).SetPipeline(pipeline).SetQueue(q)*.SetOutMode(pipe.MAPS)*/ Run()
```

#### 3、输出
输出html、map(自己实现序列化或json或对象)

#### 4、实现pipe.Pipeline接口，完成自己的pipeline，默认ConsolePipeline，已实现FilePipeline

#### 5、实现scheduler.Queue或scheduler.BlockQueue，完成自己的队列，已实现channel的内存队列

#### 6、输出内容

```
crawler: http://toutiao.io/

Title:
	程序员如何优雅地挣零花钱？
	Java String 对 null 对象的巧妙处理
	iOS 应用架构谈：组件化方案
	理解 Java 中的 ThreadLocal
	博客即简历
	简单排序之冒泡、选择、插入、希尔详细总结
	如何从菜鸟成长为（伪）架构师
	Netty 精粹之玩转 NIO 缓冲区
	系列文章：如果有 10000 台机器，你想怎么玩？
	[译] 为移动应用设计优雅的离线支持策略
	iOS 视图控制器转场详解
	React Native 开发技术周报（第 2 期）
	一起来啃犀牛书：事件处理
	教你写响应式框架（二）
	Linux 的 IO 调度
	非常流畅的滑动 tableView（iOS）
	Web-Fontmin：在线提取你需要的字体
	学习 Alamofire 封装库 Moya
	[译] 集群调度框架的架构演进之路
	Webpack 使用小记
	[译] Android 中的依赖注入：Dagger 函数库的使用（一）
	深入理解 Web 安全：ClickJacking 攻击
	Android 开发技术周报 Issue#72
	巧妙化解 dismissViewController 闪烁问题（iOS）
	[译] Linkedin 对 JVM 垃圾回收暂停的研究和解决方案
	GraphQL and Relay 浅析
	线性回归：预测上海车牌成交价格
	从无法开启 OCSP Stapling 说起
	Golang 版本的 ring buffer（变长，持久化）
	记一次移动端开发环境调试
	一个用于学习 RxJava 操作符的 App
	Gradle 依赖的统一管理
	ItChat：个人微信号的机器人（Python）
	Binder 系列（四）：获取 Service Manager
	Swift 语言指南 - Issue 48
	motion (Flint) 探秘
	Evaluating Container Platforms at Scale
URL:
	http://toutiao.io/r/geq23z
	http://toutiao.io/r/gh4d48
	http://toutiao.io/r/s9jf1i
	http://toutiao.io/r/6nra0m
	http://toutiao.io/r/up543z
	http://toutiao.io/r/44oa4v
	http://toutiao.io/r/a8b8pk
	http://toutiao.io/r/8xxtg5
	http://toutiao.io/r/tbdmv6
	http://toutiao.io/r/e2eg22
	http://toutiao.io/r/mw7fhk
	http://toutiao.io/r/8h9sbc
	http://toutiao.io/r/ungwk9
	http://toutiao.io/r/0ak9h8
	http://toutiao.io/r/dqmucy
	http://toutiao.io/r/nsf339
	http://toutiao.io/r/gw2hie
	http://toutiao.io/r/x2uz5r
	http://toutiao.io/r/zaop35
	http://toutiao.io/r/3717w6
	http://toutiao.io/r/5kn0g8
	http://toutiao.io/r/myc18i
	http://toutiao.io/r/jy091a
	http://toutiao.io/r/rr2vg8
	http://toutiao.io/r/ofo4jp
	http://toutiao.io/r/mfiioy
	http://toutiao.io/r/buu46z
	http://toutiao.io/r/xs1d1d
	http://toutiao.io/r/12jri2
	http://toutiao.io/r/ep822s
	http://toutiao.io/r/w2k6a8
	http://toutiao.io/r/gix9v2
	http://toutiao.io/r/48x2bx
	http://toutiao.io/r/fyvfv7
	http://toutiao.io/r/9f06ls
	http://toutiao.io/r/omqwug
	http://toutiao.io/r/5u7p4m
```

```
crawler: http://toutiao.io/

URL: http://toutiao.io/r/geq23z		Title: 程序员如何优雅地挣零花钱？
URL: http://toutiao.io/r/gh4d48		Title: Java String 对 null 对象的巧妙处理
URL: http://toutiao.io/r/s9jf1i		Title: iOS 应用架构谈：组件化方案
URL: http://toutiao.io/r/6nra0m		Title: 理解 Java 中的 ThreadLocal
URL: http://toutiao.io/r/up543z		Title: 博客即简历
Title: 简单排序之冒泡、选择、插入、希尔详细总结		URL: http://toutiao.io/r/44oa4v
URL: http://toutiao.io/r/a8b8pk		Title: 如何从菜鸟成长为（伪）架构师
URL: http://toutiao.io/r/8xxtg5		Title: Netty 精粹之玩转 NIO 缓冲区
URL: http://toutiao.io/r/tbdmv6		Title: 系列文章：如果有 10000 台机器，你想怎么玩？
URL: http://toutiao.io/r/e2eg22		Title: [译] 为移动应用设计优雅的离线支持策略
URL: http://toutiao.io/r/mw7fhk		Title: iOS 视图控制器转场详解
URL: http://toutiao.io/r/8h9sbc		Title: React Native 开发技术周报（第 2 期）
URL: http://toutiao.io/r/ungwk9		Title: 一起来啃犀牛书：事件处理
URL: http://toutiao.io/r/0ak9h8		Title: 教你写响应式框架（二）
URL: http://toutiao.io/r/dqmucy		Title: Linux 的 IO 调度
Title: 非常流畅的滑动 tableView（iOS）		URL: http://toutiao.io/r/nsf339
URL: http://toutiao.io/r/gw2hie		Title: Web-Fontmin：在线提取你需要的字体
URL: http://toutiao.io/r/x2uz5r		Title: 学习 Alamofire 封装库 Moya
URL: http://toutiao.io/r/zaop35		Title: [译] 集群调度框架的架构演进之路
URL: http://toutiao.io/r/3717w6		Title: Webpack 使用小记
URL: http://toutiao.io/r/5kn0g8		Title: [译] Android 中的依赖注入：Dagger 函数库的使用（一）
URL: http://toutiao.io/r/0gse1u		Title: 从零开始的 Android 新项目（一）：架构搭建篇
URL: http://toutiao.io/r/myc18i		Title: 深入理解 Web 安全：ClickJacking 攻击
Title: Android 开发技术周报 Issue#72		URL: http://toutiao.io/r/jy091a
URL: http://toutiao.io/r/rr2vg8		Title: 巧妙化解 dismissViewController 闪烁问题（iOS）
URL: http://toutiao.io/r/ofo4jp		Title: [译] Linkedin 对 JVM 垃圾回收暂停的研究和解决方案
URL: http://toutiao.io/r/mfiioy		Title: GraphQL and Relay 浅析
URL: http://toutiao.io/r/buu46z		Title: 线性回归：预测上海车牌成交价格
Title: 从无法开启 OCSP Stapling 说起		URL: http://toutiao.io/r/xs1d1d
URL: http://toutiao.io/r/12jri2		Title: Golang 版本的 ring buffer（变长，持久化）
URL: http://toutiao.io/r/ep822s		Title: 记一次移动端开发环境调试
URL: http://toutiao.io/r/w2k6a8		Title: 一个用于学习 RxJava 操作符的 App
URL: http://toutiao.io/r/gix9v2		Title: Gradle 依赖的统一管理
URL: http://toutiao.io/r/48x2bx		Title: ItChat：个人微信号的机器人（Python）
URL: http://toutiao.io/r/fyvfv7		Title: Binder 系列（四）：获取 Service Manager
URL: http://toutiao.io/r/9f06ls		Title: Swift 语言指南 - Issue 48
URL: http://toutiao.io/r/omqwug		Title: motion (Flint) 探秘
URL: http://toutiao.io/r/5u7p4m		Title: Evaluating Container Platforms at Scale
```
