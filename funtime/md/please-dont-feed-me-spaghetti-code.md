Everything is urgent, everything is critical, rush rush, develop develop, test test, now now now!

This is a common theme in both games development and software development. Management pressure always trying to get the product shipped or the next chunk of code released. Everything may appear shiny and happy on the surface but underneath, code becomes a twisted, tangled, distorted mess which starts testing the sanity of anyone that has to interact with it. 


I recently found out about the term ['Technical Debt'](http://en.wikipedia.org/wiki/Technical_debt). I hadn't heard of it before but after reading a description, I realised all software development will have encountered this debt in some form or another.  


So what is Technical Debt and how does it affect software? Well, cutting corners leads to bad decisions, which in turn leads to problems. Fixing problems takes time so when corners are cut a technical debt is created. The debt can be paid at some point in the future when the consequences of the bad decisions are fixed but frequently these debts are ignored.  When technical debt is left in place without repaying it, it grows, and accumulates interest as those bad decisions start requiring even more bad decisions to work around them.  


Technical debt leads to architectural nightmares made out of [spaghetti code](http://en.wikipedia.org/wiki/Spaghetti_code). New features gradually start requiring an ever growing number of hacks and workarounds to implement. Before long, the code base starts looks like a really high Jenga tower held together by wishes and tears in danger of collapsing at any moment. Even making simple changes to the software becomes increasingly challenging as the technical debt grows.


Taking on small amounts of technical debt does seem to be completely unavoidable. But some companies don't know how to manage their technical debt. Even fewer companies know when they should avoid taking on new technical debt or even how to start making repayments. Technical debt is actually really dangerous because it is one of a few things that can kill companies dead. 


Continuous regression testing is possibly the easiest way to find problems and identify potential code changes that create technical debt. When these kinds of problems are found at the testing stage a choice can be made between either fixing the problems (paying back some of the debt) or backing out of making the change (completely avoiding any new technical debt).


Reducing technical debt ideally should be part of a company's culture because once it starts building up, it won't remove itself. There are various testing activities that can help identify technical debt however as this debt is created by bad code and poor architectural decisions, testing won't make this debt go away. Only refactoring code, redesigning and recreating can pay back technical debt.


There isn't very much a test team can do on its own to reduce technical debt other than shouting loudly and hoping developers and architects pay attention and listen. 




[![image](http://3.bp.blogspot.com/-HYthXOcdKEw/VNHdz3u6QuI/AAAAAAAABT4/HfvPddrI_b8/s320/catcode.jpg)](http://3.bp.blogspot.com/-HYthXOcdKEw/VNHdz3u6QuI/AAAAAAAABT4/HfvPddrI_b8/s1600/catcode.jpg)
