Blah blah automation. Blah blah Selenium. Apparently automated testing is what all the cool kids are doing these days. I'm not naive enough to believe that automated testing is some kind of magic spell that when cast the software will test itself and suddenly reveal the location of all the bugs. But, having experienced first hand the pain of long drawn out manual regression testing, if ANYTHING helps ease even a small amount of that pain - I want to know about it!


A few months ago I started wondering what Selenium is and how can I find out more about this automated testing magic. After a bit of searching I realised I wasn't the only one looking for answers. A lot of other people were already asking similar questions. I started my search on [Quora](https://www.quora.com/). I like Quora, its familiar. I frequently read Quora on the train while commuting to pass the time. 


Quora directed me to the [Selenium Webdriver documentation](http://docs.seleniumhq.org/docs/) so I switched from reading Quora on my daily commute to reading Selenium documentation instead. It introduced a lot of new concepts like the difference between asserting, verifying and crazy things called Xpaths. I was still pretty baffled even after reading the documentation. I needed tutorials.


I found some videos on Youtube of an Indian sounding guy [explaining Selenium using Java](https://www.youtube.com/watch?v=-XTp5jAPKNc&list=UUbXIScRCJga-zUtypJHxULw). I tried to copy what he was doing as best I could. Eclipse (the IDE he was using on his videos) was horrid. I spent so much time working out how to "set up" Eclipse to get to the point where it would just open a web browser on its own. Eclipse forced me to learn what environment variables are and how to set them in Windows.


I persevered though the frustration and was eventually able to write this. 



[![image](http://2.bp.blogspot.com/-HtN08G8exTM/VIt5Ubwf-tI/AAAAAAAABRw/cKe0Sg-XByk/s400/javaselenium.jpg)](http://2.bp.blogspot.com/-HtN08G8exTM/VIt5Ubwf-tI/AAAAAAAABRw/cKe0Sg-XByk/s1600/javaselenium.jpg)



My first ever Selenium script. It opened a browser, went to Google and searched for cute cats. I like cats.


By this point I was aware that Selenium Webdriver could be powered by lots of different programming languages. Now the thing is, I'm not a programmer. Selenium needs a language and I wanted to be sure that I chose wisely before filling my head with copious amounts of this kind of nonsense. 


The next language I decided to investigate was PHP. I chose to look at PHP because I test software written in PHP for my day job. The first obstacle I encountered was that PHP is a server side scripting language and I didn't have a server. Luckily I share the bus to the station every day with two PHP developers. They enlightened me and told me all about [XAMPP](https://www.apachefriends.org/index.html). After a bit of faffing around I had a server. After a bit more faffing around I had written this.



[![image](http://2.bp.blogspot.com/-bCJO4j_npNM/VIt--k9AZHI/AAAAAAAABSA/55mOhUhnMC0/s400/phpselenium.jpg)](http://2.bp.blogspot.com/-bCJO4j_npNM/VIt--k9AZHI/AAAAAAAABSA/55mOhUhnMC0/s1600/phpselenium.jpg)



General observations between PHP and Java were that the PHP was faster to create than Java due to typing it into notepad++ and not having to fight against Eclipse. With Java the output from the test was shown in real time, with PHP the output was only shown once the test had finished. PHP needed a server to run, Java didn't. Java compile errors might as well have been written in alien. Getting Java to compile was especially painful. PHP was ever so slightly easier to understand why things were broken. 


Python was next on my list but as I have discovered so far, one does not simply start writing selenium scripts. I had to get my head around Python first. I had managed to download something called Python and install it. It was on my computer somewhere but had no idea how to even start using it. So I started reading this book called  [Invent your own computer games with Python](http://www.amazon.co.uk/Invent-Your-Computer-Games-Python/dp/0982106017). My general interest levels in games and gaming are much higher than non-gaming related topics. I think this really helped hold my attention. I would start copying Python from the book to make hangman or noughts and crosses. Then messing around and adding my own extra bits to see what they did. 


At one point I wrote code that had a bug in it. I couldn't see the bug in the code as I didn't really know what I was looking at. But I could sure as hell see it when I tested what I'd written. Was able to eventually work out that four white spaces were in the wrong place which had made all the logic go wonky. I hadn't really realised until that point just how fragile code can be and how easy it is to accidently write bugs.


It took a bit more Googling to work out how to install the Selenium bindings for Python. Having never really had any need to type things into a command line before that took a bit of working out too. But by trial, error and possibly a small amount of luck. I was able to tell Python "python -m ensurepip --upgrade" and "python -m pip install selenium" About an hour later after staring at various error messages and pleading with Google for help. I managed to write my first ever Selenium script in Python and get it to run.



[![image](http://3.bp.blogspot.com/-573GxgIzaU8/VIuHBJm2QRI/AAAAAAAABSQ/4rNYHfEZ_Q8/s400/pythonselenium.jpg)](http://3.bp.blogspot.com/-573GxgIzaU8/VIuHBJm2QRI/AAAAAAAABSQ/4rNYHfEZ_Q8/s1600/pythonselenium.jpg)



I still have a long way to go on my Selenium adventure but I'm in a better place now than I was a few months ago. 


I think subconsciously, I've already decided that I want to continue learning Python. It held my attention much longer than Java or PHP. I especially like the lack of semi-colons at the end of every line. I think the more Python I learn, the more Selenium will fall into place. I just have to keep chipping away slowly at online tutorials until it starts to [look like an Elephant](http://www.boardofwisdom.com/cachetogo/images/quotes/307913.png). 