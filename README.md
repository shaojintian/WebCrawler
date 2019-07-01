#WebCrawler网络爬虫与框架
网络爬虫是互联网用户的模仿者

实现一个开箱即用的工具类WebCrawler

特点：可扩展，高度定制化的网络爬虫框架
##总体功能和需求
1. 下载器：下载给定网址相应的内容
2. 分析器：
- 分析下载到的内容，分析筛选到可用的内容（以下均称为条目）;<br/>
- 组装新的下载请求发给"下载器";<br/>
- 过滤掉不符合要求的网址;<br/>
特点：WebCrawler提供高度定制化接口，根据用户的需求改变分析策略和内容

3. 条目处理管道：接受所有的条目，对其进行相应定制化的处理（eg：命名，存储...）
##总体设计
/module
- 下载器：
- 分析器：
- 条目处理管道
- 调度器<br/>
负责整个任务的启动和各个模块的整合处理

##详细设计
###数据流图
![dataFlow](/docs/dataFlow.png)<br/>
###模块架构图
![moduleArchi](/docs/moduleArchi.png)
##工具的实现
- /errors 错误处理
- /helper/log   监控日志处理
- /toolkit/buffer 缓冲器，缓冲池
- /toolkit/reader 多重读取器
##组件的实现
- /module/local 主要数据结构的实现<br/>
下载器downloader,分析器analyzer,条目处理管道pipeline<br/>

##调度器的实现
- /scheduler 调度器  
##未来工作
- 并发
- 支持cookie

