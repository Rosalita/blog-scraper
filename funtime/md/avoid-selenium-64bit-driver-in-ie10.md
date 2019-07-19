When it comes to testing anything in a browser, Internet Explorer tends to have the reputation of being the black sheep of the browser family. Anyone with any experience of testing know that there is a greater chance of something being broken in Internet Explorer than any other browser. Let's face it, IE doesn't have a great track record. As software testers, we all remember the pain of having to support IE8, IE7, IE6. We also remember the moments when support for certain versions of IE was dropped, along with the subsequent wave of euphoria upon realising we no longer had to test in them. But I digress, new versions of Internet Explorer come along, like buses, to replace the older versions, which we eventually drop.



[![image](http://3.bp.blogspot.com/-jHgB20CTF74/VkTx_A0NvXI/AAAAAAAABa8/nTdyi_aTYts/s320/slowpoke.jpg)](http://3.bp.blogspot.com/-jHgB20CTF74/VkTx_A0NvXI/AAAAAAAABa8/nTdyi_aTYts/s1600/slowpoke.jpg)



Internet Explorer 10 is currently a supported browser for some software that I test. I follow the widely accepted practice of writing automated tests to do the repetitive grunt testing work so that I have more time to test the complex bits (that can't be automated) by hand.


I've been running automated tests in IE10 happily using the 32-bit version of Selenium Internet Explorer Driver for quite some time. Until this morning. This morning, everything broke. 


Well I say broke, what actually happened was that all the tests that used to take 5-10 seconds each to run, suddenly started taking 2 -3 minutes each to run! I watched some of these tests running and I saw that the IE driver was mysteriously typing text into all the text boxes very, very slowly. It's speed was comparable to an asthmatic snail.


So what changed? Well, a bit of investigation revealed that someone else had 'upgraded' the test suite to use the Selenium 64-bit Internet Explorer Driver from the usual 32-bit driver.



But why would this cause everything to break so horrifically in IE10?



Well, in IE there is a manager process that looks after the top level window, then there are separate content processes that look after rendering the HTML inside the browser.


Before IE10 came along, the manager process and the content processes both used the same number of bits.


So if you ran a 32-bit version of IE you got a 32-bit manager process to look after the top level window and you got 32-bit content processes to render the HTML.


Likewise, if you ran a 64-bit version of IE you got a 64-bit manager process to look after the top level window and you got 64-bit content processes to render the HTML.


Then IE10 came along and changed everything because it could. In 64-bit IE10 the manager process was 64-bit (as you would expect) but the content processes, well they weren't 64-bit any more. That would be too logical and sensible. The content processes remained 32-bit. I think the reason they didn't change the content process to 64-bit was to try keep IE10 compatible with all the existing browser plug-ins.


Anyway, part of IE10 (the manager process that controls the top level window) is 64-bit and the rest of it (the content processes that render the HTML) are 32-bit. Now this might seem a tiny bit crazy because a Windows a 32-bit executable can't load a 64-bit DLL and vice-versa, a 64-bit executable can't load a 32-bit DLL. This is the very reason why there was a separate 32-bit and 64-bit versions of IE in the first place!



So what was actually happening to my tests when they were using the 64-bit Selenium Internet Explorer driver?



The tests were sending key presses to the browser. The sending of a key press is done using a hook. The IE Driver sends a 'key down' message followed by a the name of the key, followed by a 'key up' message. It does this for each key press. Because the way these messages are sent is [asynchronous](”http://searchnetworking.techtarget.com/definition/asynchronous”), the driver has to wait to make sure that the 'key down' message is processed first so that the key presses don't happen out of order. The driver does this by listening for the 'key down' message to be processed before continuing.  


In 64-bit IE10 the hook can be attached to the top level manager process (because that part is 64-bit) but the hook then fails to attach to the content process (because that part is 32-bit).  


So the 64-bit manager process sends a key press, then listens to hear whether or not the 'key down' message was received by the 32-bit content process. But because the 32-bit content process can't load a 64-bit DLL, it never responds to say "Yeah I've dealt with the 'key down' you sent". Which means the manager process times out waiting for the content process to respond. This time-out takes about 5 seconds and is triggered for **every single key press**. 


The resulting effect is that the IE driver types 1 key every 5 seconds. So if your test data contains fancy long words like "inexplicably" it's going to take a whole minute to type that string in. You know your automated tests are seriously broken when a human can perform the same test in less time than it takes the test script.   


This issue is at the heart of the Selenium 64-bit Internet Explorer Driver and is certainly never, ever going to be fixed.  Especially given that [Microsoft intend to discontinue all support for legacy versions of IE from January 12th 2016](”https://support.microsoft.com/en-us/gp/microsoft-internet-explorer”)


Fortunately I was lucky and the work around in my situation was simply to roll back to using the 32-bit version of IE Driver. 


Beware the Selenium 64-bit Internet Explorer Driver. Apparently it can't handle taking screen-shots either for exactly the same 32-bit trying to use 64-bit reason.

