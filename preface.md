# Go語言聖經（中文版）

Go語言聖經 [《The Go Programming Language》](http://gopl.io) 中文版本，僅供學習交流之用。

[![](cover_middle.jpg)](https://github.com/golang-china/gopl-zh)

- 在線版本：http://golang-china.github.io/gopl-zh
- 離線版本：http://github.com/golang-china/gopl-zh/archive/gh-pages.zip
- 項目主頁：http://github.com/golang-china/gopl-zh
- 原版官網：http://gopl.io


-------

# 譯者序

在上個世紀70年代，貝爾實驗室的[Ken Thompson][KenThompson]和[Dennis M. Ritchie][DennisRitchie]合作發明了[UNIX](http://doc.cat-v.org/unix/)操作繫統，同時[Dennis M. Ritchie][DennisRitchie]爲了解決[UNIX](http://doc.cat-v.org/unix/)繫統的移植性問題而發明了C語言，貝爾實驗室的[UNIX](http://doc.cat-v.org/unix/)和C語言兩大發明奠定了整個現代IT行業最重要的軟件基礎（目前的三大桌面操作繫統的中[Linux](http://www.linux.org/)和[Mac OS X](http://www.apple.com/cn/osx/)都是源於[UINX]()繫統，兩大移動平台的操作繫統iOS和Android也都是源於[UNIX](http://doc.cat-v.org/unix/)繫統。C繫家族的編程語言占據統治地位達幾十年之久）。在[UINX]()和C語言發明40年之後，目前已經在Google工作的[Ken Thompson](http://genius.cat-v.org/ken-thompson/)和[Rob Pike](http://genius.cat-v.org/rob-pike/)（他們在貝爾實驗室時就是同事）、還有[Robert Griesemer](http://research.google.com/pubs/author96.html)（設計了V8引擎和HotSpot虛擬機）一起合作，爲了解決在21世紀多核和網絡化環境下越來越複雜的編程問題而發明了Go語言。從Go語言庫早期代碼庫日誌可以看出它的演化歷程（Git用`git log --before={2008-03-03} --reverse`命令査看）：

![](./images/go-log04.png)

從早期提交日誌中也可以看出，Go語言是從[Ken Thompson](http://genius.cat-v.org/ken-thompson/)發明的B語言、[Dennis M. Ritchie](http://genius.cat-v.org/dennis-ritchie/)發明的C語言逐步演化過來的，是C語言家族的成員，因此很多人將Go語言稱爲21世紀的C語言。縱觀這幾年來的發展趨勢，Go語言已經成爲雲計算、雲存儲時代最重要的基礎編程語言。

在C語言發明之後約5年的時間之後（1978年），[Brian W. Kernighan](http://www.cs.princeton.edu/~bwk/)和[Dennis M. Ritchie](http://genius.cat-v.org/dennis-ritchie/)合作編寫出版了C語言方面的經典敎材《[The C Programming Language](http://s3-us-west-2.amazonaws.com/belllabs-microsite-dritchie/cbook/index.html)》，該書被譽爲C語言程序員的聖經，作者也被大家親切地稱爲[K&R](https://en.wikipedia.org/wiki/K%26R)。同樣在Go語言正式發布（2009年）約5年之後（2014年開始寫作，2015年出版），由Go語言核心糰隊成員[Alan A. A. Donovan](https://github.com/adonovan)和[K&R](https://en.wikipedia.org/wiki/K%26R)中的[Brian W. Kernighan](http://www.cs.princeton.edu/~bwk/)合作編寫了Go語言方面的經典敎材《[The Go Programming Language](http://gopl.io)》。Go語言被譽爲21世紀的C語言，如果説[K&R](https://en.wikipedia.org/wiki/K%26R)所著的是聖經的舊約，那麽D&K所著的必將成爲聖經的新約。該書介紹了Go語言幾乎全部特性，併且隨着語言的深入層層遞進，對每個細節都解讀得非常細致，每一節內容都精綵不容錯過，是廣大Gopher的必讀書目。大部分Go語言核心糰隊的成員都參與了該書校對工作，因此該書的質量是可以完全放心的。

同時，單憑閲讀和學習其語法結構併不能眞正地掌握一門編程語言，必須進行足夠多的編程實踐——親自編寫一些程序併研究學習别人寫的程序。要從利用Go語言良好的特性使得程序模塊化，充分利用Go的標準函數庫以Go語言自己的風格來編寫程序。書中包含了上百個精心挑選的習題，希望大家能先用自己的方式嚐試完成習題，然後再參考官方給出的解決方案。

該書英文版約從2015年10月開始公開發售，其中日文版本最早參與翻譯和審校（參考致謝部分）。在2015年10月，我們併不知道中文版是否會及時引進、將由哪家出版社引進、引進將由何人來翻譯、何時能出版，這些信息都成了一個祕密。中国的Go語言社區是全球最大的Go語言社區，我們從一開始就始終緊跟着Go語言的發展腳步。我們應該也完全有能力以中国Go語言社區的力量同步完成Go語言聖經中文版的翻譯工作。與此同時，国內有很多Go語言愛好者也在積極關註該書（本人也在第一時間購買了紙質版本，[亞馬遜價格314人民幣](http://www.amazon.cn/The-Go-Programming-Language-Donovan-Alan-A-A/dp/0134190440/)。補充：国內也卽將出版英文版，[價格79元](http://product.china-pub.com/4912464)）。爲了Go語言的學習和交流，大家決定合作免費翻譯該書。

翻譯工作從2015年11月20日前後開始，到2016年1月底初步完成，前後歷時約2個月時間（在其它語言版本中，全球第一個完成翻譯的，基本做到和原版同步）。其中，[chai2010](https://github.com/chai2010)翻譯了前言、第2~4章、第10~13章，[Xargin](https://github.com/cch123)翻譯了第1章、第6章、第8~9章，[CrazySssst](https://github.com/CrazySssst)翻譯了第5章，[foreversmart](https://github.com/foreversmart)翻譯了第7章，大家共同參與了基本的校驗工作，還有其他一些朋友提供了積極的反饋建議。如果大家還有任何問題或建議，可以直接到中文版項目頁面提交[Issue](https://github.com/golang-china/gopl-zh/issues)，如果發現英文版原文在[勘誤](http://www.gopl.io/errata.html)中未提到的任何錯誤，可以直接去[英文版項目](https://github.com/adonovan/gopl.io/)提交。

最後，希望這本書能夠幫助大家用Go語言快樂地編程。

2016年 1月 於 武漢

-------

# 前言

*“Go是一個開源的編程語言，它很容易用於構建簡單、可靠和高效的軟件。”（摘自Go語言官方網站：http://golang.org ）*

Go語言由來自Google公司的[Robert Griesemer](http://research.google.com/pubs/author96.html)，[Rob Pike](http://genius.cat-v.org/rob-pike/)和[Ken Thompson](http://genius.cat-v.org/ken-thompson/)三位大牛於2007年9月開始設計和實現，然後於2009年的11月對外正式發布（譯註：關於Go語言的創世紀過程請參考 http://talks.golang.org/2015/how-go-was-made.slide ）。語言及其配套工具的設計目標是具有表達力，高效的編譯和執行效率，有效地編寫高效和健壯的程序。

Go語言有着和C語言類似的語法外表，和C語言一樣是專業程序員的必備工具，可以用最小的代價獲得最大的戰果。
但是它不僅僅是一個更新的C語言。它還從其他語言借鑒了很多好的想法，同時避免引入過度的複雜性。
Go語言中和併發編程相關的特性是全新的也是有效的，同時對數據抽象和面向對象編程的支持也很靈活。
Go語言同時還集成了自動垃圾收集技術用於更好地管理內存。

Go語言尤其適合編寫網絡服務相關基礎設施，同時也適合開發一些工具軟件和繫統軟件。
但是Go語言確實是一個通用的編程語言，它也可以用在圖形圖像驅動編程、移動應用程序開發
和機器學習等諸多領域。目前Go語言已經成爲受歡迎的作爲無類型的腳本語言的替代者：
因爲Go編寫的程序通常比腳本語言運行的更快也更安全，而且很少會發生意外的類型錯誤。

Go語言還是一個開源的項目，可以免費獲編譯器、庫、配套工具的源代碼。
Go語言的貢獻者來自一個活躍的全球社區。Go語言可以運行在類[UNIX](http://doc.cat-v.org/unix/)繫統——
比如[Linux](http://www.linux.org/)、[FreeBSD](https://www.freebsd.org/)、[OpenBSD](http://www.openbsd.org/)、[Mac OSX](http://www.apple.com/cn/osx/)——和[Plan9](http://plan9.bell-labs.com/plan9/)繫統和[Microsoft Windows](https://www.microsoft.com/zh-cn/windows/)操作繫統之上。
Go語言編寫的程序無需脩改就可以運行在上面這些環境。

本書是爲了幫助你開始以有效的方式使用Go語言，充分利用語言本身的特性和自帶的標準庫去編寫清晰地道的Go程序。

{% include "./links.md" %}
