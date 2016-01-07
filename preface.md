# Go語言聖經（中文版）

Go語言聖經 [《The Go Programming Language》](http://gopl.io) 中文版本，僅供學習交流之用。

- 項目主頁：http://github.com/golang-china/gopl-zh
- 項目進度：http://github.com/golang-china/gopl-zh/blob/master/progress.md
- 參與人員：http://github.com/golang-china/gopl-zh/blob/master/CONTRIBUTORS.md
- 離線版本：http://github.com/golang-china/gopl-zh/archive/gh-pages.zip
- 在線預覽：http://golang-china.github.io/gopl-zh
- 原版官網：http://gopl.io

[![](cover_middle.jpg)](https://github.com/golang-china/gopl-zh)

**版權聲明：** <a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>。

<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="./images/by-nc-sa-4.0-88x31.png"></img></a>

嚴禁任何商業行爲使用或引用該文檔的全部或部分內容!

歡迎大家提供建議!

-------

# 譯者序

在上個世紀70年代，貝爾實驗室的Ken Thompson和Dennis M. Ritchie合作發明了UNIX操作繫統，同時Dennis M. Ritchie爲了解決UNIX繫統的移植性問題而發明了C語言，貝爾實驗室的UNIX和C語言兩大發明奠定了整個現代IT行業最重要的軟件基礎（目前的三大桌面操作繫統的中Linux和Mac OS X都是源於UINX繫統，兩大移動平台的操作繫統iOS和Android也都是源於UNIX繫統。C繫家族的編程語言占據統治地位達幾十年之久）。在UINX和C語言發明40年之後，目前已經在Google工作的Ken Thompson和Rob Pike（他們在貝爾實驗室時就是同事）、同樣還有Robert Griesemer一起合作，爲了解決21世紀的多覈和網絡化編程問題而發明了Go語言。從Go語言庫早期代碼庫日誌可以看出它的演化歷程：

![](./images/go-log04.png)

從早期提交日誌中也可以看出，Go語言是從Ken Thompson發明的B語言、Dennis M. Ritchie發明的C語言逐步演化過來的，是C語言家族的成員，因此很多人將Go語言稱爲21世紀的C語言。

在C語言發明之後約5年的時間之後（1978年），Brian W. Kernighan和Dennis M. Ritchie合作編寫出版了C語言方面的經典敎材《The C Programming Language》，該書被譽爲C語言程序員的聖經，作者也被大家親切地稱爲K&R。同樣在Go語言正式發布（2009年）約5年之後（2014年開始寫作，2015年出版），由Go語言覈心糰隊成員Alan A. A. Donovan和K&R中的Brian W. Kernighan合作編寫了Go語言方面的經典敎材《The Go Programming Language》。Go語言被譽爲21世紀的C語言，如果説K&R所著的是聖經的舊約，那麽D&K所著的必將成爲聖經的新約。該書介紹了Go語言幾乎全部特性，併且隨着語言的深入層層遞進，對每個細節都解讀得非常細致，每一節內容都精綵不容錯過，是廣大Gopher的必讀書目。同時，大部分Go語言覈心糰隊的成員都參與了該書校對工作，因此該書的質量是可以完全放心的。

同時，單憑閲讀和學習其語法結構併不能眞正地掌握一門編程語言，必鬚進行足夠多的編程實踐——親自編寫一些程序併研究學習别人寫的程序。要從利用Go語言良好的特性使得程序模塊化，充分利用Go的標準函數庫以Go語言自己的風格來編寫程序。書中包含了上百個精心挑選的習題，希望大家能先用自己的方式嚐試完成習題，然後再參考官方給出的解決方案。

該書英文版約從2015年10月開始公開發售，同時同步發售的還有日文版本。不過比較可惜的是，中文版併沒有在同步發售之列，甚至連中文版是否會引進、卽使引進將由何人翻譯、何時出版都成了一個祕密。中国的Go語言社區是全球最大的Go語言社區，我們一直在緊跟Go語言發展的腳步。我們應該也完全有能力以Go語言社區的力量同步完成Go語言聖經中文版的翻譯工作。與此同時，国內有很多Go語言愛好者也在積極關註該書（本人也在第一時間購買了紙質版本，亞馬遜價格314人民幣）。爲了Go語言的學習和交流，大家決定合作免費翻譯該書。

翻譯工作從2015年11月20日前後開始，到2016年1月底初步完成，前後歷時約2個月時間。其中，chai2010翻譯了前言、第2~4章、第10~13章，Xargin翻譯了第1章、第6章、第8~9章，CrazySssst翻譯了第5章，foreversmart翻譯了第7章，大家一起參與了基本的校驗工作，還有其他一些朋友提供了積極的反饋建議。如果大家有任何問題或建議，可以直接到中文版項目頁面提交Issue，如果發現英文原文有錯誤，可以直接去英文版項目提交。

最後，希望這本書能夠幫助大家成爲Go語言的高手。

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


